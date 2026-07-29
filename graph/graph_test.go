package graph

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestGraph_LoadFromFile_NormalCase(t *testing.T) {
	// Создаем временный файл с корректной матрицей (целые числа)
	tmpFile, err := os.CreateTemp("", "graph_normal_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `0 5 3
5 0 2
3 2 0`
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	g := NewGraph()
	err = g.LoadFromFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expected := [][]float64{
		{0, 5, 3},
		{5, 0, 2},
		{3, 2, 0},
	}

	if !reflect.DeepEqual(g.Matrix(), expected) {
		t.Errorf("Expected matrix %v, got %v", expected, g.Matrix())
	}

	if g.Vertex() != 3 {
		t.Errorf("Expected vertex count 3, got %d", g.Vertex())
	}
}

func TestGraph_LoadFromFile_FloatNumbers(t *testing.T) {
	// Тест с числами с плавающей точкой
	tmpFile, err := os.CreateTemp("", "graph_float_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `0 2.5 1.75
2.5 0 3.14
1.75 3.14 0`
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	g := NewGraph()
	err = g.LoadFromFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expected := [][]float64{
		{0, 2.5, 1.75},
		{2.5, 0, 3.14},
		{1.75, 3.14, 0},
	}

	if !reflect.DeepEqual(g.Matrix(), expected) {
		t.Errorf("Expected matrix %v, got %v", expected, g.Matrix())
	}
}

func TestGraph_LoadFromFile_EmptyFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "graph_empty_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	g := NewGraph()
	err = g.LoadFromFile(tmpFile.Name())

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if !errors.Is(err, ErrEmptyFile) {
		t.Errorf("Expected ErrEmptyFile, got %v", err)
	}
}

func TestGraph_LoadFromFile_NonSquareMatrix(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "graph_nonsquare_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `0 5 3
5 0`
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	g := NewGraph()
	err = g.LoadFromFile(tmpFile.Name())

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if !errors.Is(err, ErrNonSquareMatrix) {
		t.Errorf("Expected ErrNonSquareMatrix, got %v", err)
	}
}

func TestGraph_LoadFromFile_InvalidParse(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "graph_invalid_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `0 5 a
5 0 2
3 2 0`
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	g := NewGraph()
	err = g.LoadFromFile(tmpFile.Name())

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if !errors.Is(err, ErrParse) {
		t.Errorf("Expected ErrParse, got %v", err)
	}
}

func TestGraph_LoadFromFile_FileNotExist(t *testing.T) {
	g := NewGraph()
	err := g.LoadFromFile("/nonexistent/file/path.txt")

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if !errors.Is(err, ErrOpenFile) {
		t.Errorf("Expected ErrOpenFile, got %v", err)
	}
}

func TestGraph_ExportToDot(t *testing.T) {
	g := NewGraph()
	g.matrix = [][]float64{
		{0, 5, 3},
		{5, 0, 2},
		{3, 2, 0},
	}
	g.vertex = 3

	tmpFile, err := os.CreateTemp("", "graph_output_*.dot")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	err = g.ExportToDot(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Проверяем содержимое файла
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := `graph G {
    1 -- 2 [label=5.00];
    1 -- 3 [label=3.00];
    2 -- 3 [label=2.00];
}
`

	if string(content) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(content))
	}
}

func TestGraph_HasNegativeWeight(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		expected bool
	}{
		{
			name: "No negative weights",
			matrix: [][]float64{
				{0, 5, 3},
				{5, 0, 2},
				{3, 2, 0},
			},
			expected: false,
		},
		{
			name: "Has negative weight",
			matrix: [][]float64{
				{0, -5, 3},
				{-5, 0, 2},
				{3, 2, 0},
			},
			expected: true,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]float64{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{matrix: tt.matrix}
			result := g.HasNegativeWeight()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
