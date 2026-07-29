package algorithms

import (
	"errors"
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"os"
	"reflect"
	"strconv"
	"testing"
)

// ==================== КОРРЕКТНЫЕ ТЕСТЫ ДЛЯ DFS ====================

func TestDFS_Correct_SimpleGraph(t *testing.T) {
	g := graph.NewGraph()
	matrix := [][]float64{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	result, err := ga.DFS(g, 1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Проверяем, что все вершины посещены
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices, got %d", len(result))
	}

	// Проверяем, что результат не пустой
	if len(result) == 0 {
		t.Error("Expected non-empty result")
	}

	t.Logf("DFS traversal order: %v", result)
}

func TestDFS_Correct_CompleteGraph(t *testing.T) {
	g := graph.NewGraph()
	matrix := [][]float64{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	result, err := ga.DFS(g, 2)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// В полном графе должны посетить все вершины
	if len(result) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(result))
	}

	// Первая вершина должна быть стартовой
	if result[0] != 2 {
		t.Errorf("First vertex should be %d, got %d", 2, result[0])
	}

	t.Logf("DFS traversal from vertex 2: %v", result)
}

// ==================== НЕКОРРЕКТНЫЕ ТЕСТЫ ДЛЯ DFS ====================

func TestDFS_EmptyGraph(t *testing.T) {
	ga := &GraphAlgs{}

	// Тест с nil графом
	result, err := ga.DFS(nil, 1)

	if err == nil {
		t.Error("Expected error for nil graph, got nil")
	}

	if !errors.Is(err, ErrEmptyGraph) {
		t.Errorf("Expected ErrEmptyGraph, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	// Тест с пустым графом (0 вершин)
	g := graph.NewGraph()
	emptyMatrix := [][]float64{}
	tmpFile := createTempMatrixFile(t, emptyMatrix)
	defer removeTempFile(t, tmpFile)

	err = g.LoadFromFile(tmpFile)
	if err == nil {
		// Если загрузка пустого файла не вызывает ошибку, проверяем DFS
		result, err = ga.DFS(g, 1)
		if err == nil {
			t.Error("Expected error for empty graph, got nil")
		}
		if !errors.Is(err, ErrEmptyGraph) {
			t.Errorf("Expected ErrEmptyGraph, got %v", err)
		}
	}
}

func TestDFS_InvalidStartVertex(t *testing.T) {
	g := graph.NewGraph()
	matrix := [][]float64{
		{0, 1},
		{1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	// Тест с start < 1
	result, err := ga.DFS(g, 0)

	if err == nil {
		t.Error("Expected error for start vertex 0 (less than 1), got nil")
	}

	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	// Тест с start > Vertex()
	result, err = ga.DFS(g, 3)

	if err == nil {
		t.Error("Expected error for start vertex 3 (greater than vertex count), got nil")
	}

	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestDFS_DisconnectedGraph(t *testing.T) {
	g := graph.NewGraph()
	// Несвязный граф: две компоненты
	matrix := [][]float64{
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	// Начинаем из первой компоненты
	result, err := ga.DFS(g, 1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Должны посетить только вершины первой компоненты (1 и 2)
	if len(result) != 2 {
		t.Errorf("Expected 2 vertices visited (only first component), got %d", len(result))
	}

	// Проверяем, что посетили только вершины 1 и 2
	visited := make(map[int]bool)
	for _, v := range result {
		visited[v] = true
	}

	if !visited[1] || !visited[2] {
		t.Errorf("Expected vertices 1 and 2 to be visited, got %v", result)
	}

	if visited[3] || visited[4] {
		t.Errorf("Expected vertices 3 and 4 not to be visited, but got %v", result)
	}
}

// ==================== КОРРЕКТНЫЕ ТЕСТЫ ДЛЯ BFS ====================

func TestBFS_Correct_SimpleGraph(t *testing.T) {
	g := graph.NewGraph()
	matrix := [][]float64{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	result, err := ga.BFS(g, 1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Проверяем, что все вершины посещены
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices, got %d", len(result))
	}

	// BFS должен начинаться с вершины 1
	if result[0] != 1 {
		t.Errorf("First vertex should be 1, got %d", result[0])
	}

	t.Logf("BFS traversal order: %v", result)
}

func TestBFS_Correct_LinearGraph(t *testing.T) {
	g := graph.NewGraph()
	// Линейный граф: 1-2-3-4
	matrix := [][]float64{
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{0, 0, 1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	// Начинаем с вершины 3
	result, err := ga.BFS(g, 3)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Должны посетить все вершины
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices, got %d", len(result))
	}

	// Первая вершина должна быть 3
	if result[0] != 3 {
		t.Errorf("First vertex should be 3, got %d", result[0])
	}

	t.Logf("BFS traversal from vertex 3: %v", result)
}

// ==================== НЕКОРРЕКТНЫЕ ТЕСТЫ ДЛЯ BFS ====================

func TestBFS_EmptyGraph(t *testing.T) {
	ga := &GraphAlgs{}

	// Тест с nil графом
	result, err := ga.BFS(nil, 1)

	if err == nil {
		t.Error("Expected error for nil graph, got nil")
	}

	if !errors.Is(err, ErrEmptyGraph) {
		t.Errorf("Expected ErrEmptyGraph, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	// Тест с пустым графом
	g := graph.NewGraph()
	emptyMatrix := [][]float64{}
	tmpFile := createTempMatrixFile(t, emptyMatrix)
	defer removeTempFile(t, tmpFile)

	err = g.LoadFromFile(tmpFile)
	if err == nil {
		result, err = ga.BFS(g, 1)
		if err == nil {
			t.Error("Expected error for empty graph, got nil")
		}
		if !errors.Is(err, ErrEmptyGraph) {
			t.Errorf("Expected ErrEmptyGraph, got %v", err)
		}
	}
}

func TestBFS_InvalidStartVertex(t *testing.T) {
	g := graph.NewGraph()
	matrix := [][]float64{
		{0, 1},
		{1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	// Тест с start < 1
	result, err := ga.BFS(g, 0)

	if err == nil {
		t.Error("Expected error for start vertex 0 (less than 1), got nil")
	}

	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	// Тест с start > Vertex()
	result, err = ga.BFS(g, 5)

	if err == nil {
		t.Error("Expected error for start vertex 5 (greater than vertex count), got nil")
	}

	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestBFS_DisconnectedGraph(t *testing.T) {
	g := graph.NewGraph()
	// Несвязный граф
	matrix := [][]float64{
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0}, // изолированная вершина
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}
	// Начинаем из первой компоненты
	result, err := ga.BFS(g, 1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// BFS должен посетить только вершины первой компоненты (1 и 2)
	if len(result) != 2 {
		t.Errorf("Expected 2 vertices visited (only first component), got %d", len(result))
	}

	// Проверяем, что посетили только вершины 1 и 2
	visited := make(map[int]bool)
	for _, v := range result {
		visited[v] = true
	}

	if !visited[1] || !visited[2] {
		t.Errorf("Expected vertices 1 and 2 to be visited, got %v", result)
	}

	if visited[3] || visited[4] || visited[5] {
		t.Errorf("Expected vertices 3, 4, 5 not to be visited, but got %v", result)
	}
}

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================

func createTempMatrixFile(t *testing.T, matrix [][]float64) string {
	tmpFile, err := os.CreateTemp("", "graph_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Записываем матрицу в файл
	for _, row := range matrix {
		line := ""
		for i, val := range row {
			if i > 0 {
				line += " "
			}
			line += floatToString(val)
		}
		line += "\n"
		_, err := tmpFile.WriteString(line)
		if err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
	}

	tmpFile.Close()
	return tmpFile.Name()
}

func removeTempFile(t *testing.T, path string) {
	err := os.Remove(path)
	if err != nil {
		t.Logf("Warning: failed to remove temp file %s: %v", path, err)
	}
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func TestDijkstra_ShortestPath(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1, 10},
		{1, 0, 2},
		{10, 2, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	result, err := ga.GetShortestPathBetweenVertices(g, 1, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := 3.0
	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}
func TestDijkstra_InvalidVertex(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1},
		{1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	_, err = ga.GetShortestPathBetweenVertices(g, 0, 2)
	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}

	_, err = ga.GetShortestPathBetweenVertices(g, 1, 3)
	if !errors.Is(err, ErrVertex) {
		t.Errorf("Expected ErrVertex, got %v", err)
	}
}

func TestDijkstra_NegativeWeight(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, -1},
		{-1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	_, err = ga.GetShortestPathBetweenVertices(g, 1, 2)
	if !errors.Is(err, graph.ErrNegWeight) {
		t.Errorf("Expected ErrNegWeight, got %v", err)
	}
}

func TestFloydWarshall_ShortestPaths(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1, 10},
		{1, 0, 2},
		{10, 2, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	result, err := ga.GetShortestPathsBetweenAllVertices(g)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := 3.0
	if result[0][2] != expected {
		t.Errorf("Expected shortest path 1 -> 3 = %.2f, got %.2f", expected, result[0][2])
	}

	if result[2][0] != expected {
		t.Errorf("Expected shortest path 3 -> 1 = %.2f, got %.2f", expected, result[2][0])
	}
}

func TestFloydWarshall_NegativeWeight(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, -1},
		{-1, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	_, err = ga.GetShortestPathsBetweenAllVertices(g)
	if !errors.Is(err, graph.ErrNegWeight) {
		t.Errorf("Expected ErrNegWeight, got %v", err)
	}
}

func TestPrim_LeastSpanningTree(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1, 10},
		{1, 0, 2},
		{10, 2, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	result, err := ga.GetLeastSpanningTree(g)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := [][]float64{
		{0, 1, 0},
		{1, 0, 2},
		{0, 2, 0},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected MST %v, got %v", expected, result)
	}
}

func TestTSPNearestNeighbor_SimpleGraph(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1, 5, 10},
		{1, 0, 2, 4},
		{5, 2, 0, 3},
		{10, 4, 3, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	result, err := ga.SolveTravelingSalesmanProblemNearestNeighbor(g)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedPath := []int{1, 2, 3, 4, 1}
	expectedDistance := 16.0

	if !reflect.DeepEqual(result.Vertices, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, result.Vertices)
	}

	if result.Distance != expectedDistance {
		t.Errorf("Expected distance %.2f, got %.2f", expectedDistance, result.Distance)
	}
}

func TestTSPBruteForce_SimpleGraph(t *testing.T) {
	g := graph.NewGraph()

	matrix := [][]float64{
		{0, 1, 5, 10},
		{1, 0, 2, 4},
		{5, 2, 0, 3},
		{10, 4, 3, 0},
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	result, err := ga.SolveTravelingSalesmanProblemBruteForce(g)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedDistance := 13.0

	if result.Distance != expectedDistance {
		t.Errorf("Expected distance %.2f, got %.2f", expectedDistance, result.Distance)
	}

	if len(result.Vertices) != g.Vertex()+1 {
		t.Errorf("Expected path length %d, got %d", g.Vertex()+1, len(result.Vertices))
	}

	if result.Vertices[0] != result.Vertices[len(result.Vertices)-1] {
		t.Errorf("Expected path to return to start, got %v", result.Vertices)
	}
}

func TestTSPBruteForce_TooLargeGraph(t *testing.T) {
	g := graph.NewGraph()

	size := 11
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			if i != j {
				matrix[i][j] = 1
			}
		}
	}

	tmpFile := createTempMatrixFile(t, matrix)
	defer removeTempFile(t, tmpFile)

	err := g.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load graph: %v", err)
	}

	ga := &GraphAlgs{}

	_, err = ga.SolveTravelingSalesmanProblemBruteForce(g)
	if err == nil {
		t.Error("Expected error for too large graph, got nil")
	}
}
