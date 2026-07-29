package graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NoEdge = 0.0
)

var (
	ErrNonSquareMatrix = errors.New("matrix is non-square")
	ErrEmptyFile       = errors.New("file is empty")
	ErrOpenFile        = errors.New("couldn't open the file")
	ErrCreateFile      = errors.New("couldn't create the file")
	ErrParse           = errors.New("parsing error")
	ErrNegWeight       = errors.New("matrix has negative weights")
	ErrPathNotFound    = errors.New("path not found")
)

type Graph struct {
	matrix [][]float64
	vertex int
}

func NewGraph() *Graph {
	return &Graph{
		matrix: make([][]float64, 0),
		vertex: 0,
	}
}

func (g *Graph) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOpenFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]float64
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)

		row, err := parseRow(parts)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrParse, err)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%w: %w", ErrParse, err)
	}

	if isEmpty(matrix) {
		return ErrEmptyFile
	}

	if isSquare(matrix) == false {
		return ErrNonSquareMatrix
	}

	g.matrix = matrix
	g.vertex = len(matrix)
	return nil
}

func (g *Graph) ExportToDot(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCreateFile, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "graph G {")

	for i := range g.vertex {
		for j := i; j < g.vertex; j++ {
			if isEdge(g.matrix[i][j]) {
				fmt.Fprintf(writer, "    %d -- %d [label=%.2f];\n", i+1, j+1, g.matrix[i][j])
			}
		}
	}

	fmt.Fprintln(writer, "}")
	return nil
}

func (g *Graph) Matrix() [][]float64 {
	if g.matrix == nil {
		return nil
	}

	copyMatrix := make([][]float64, len(g.matrix))

	for i := range g.matrix {
		if g.matrix[i] != nil {
			copyMatrix[i] = make([]float64, len(g.matrix[i]))
			copy(copyMatrix[i], g.matrix[i])
		}
	}

	return copyMatrix
}

func (g *Graph) Vertex() int {
	copyVertex := g.vertex
	return copyVertex
}

func (g *Graph) HasNegativeWeight() bool {
	for _, row := range g.matrix {
		for _, weight := range row {
			if weight < 0 {
				return true
			}
		}
	}
	return false
}

func parseRow(parts []string) ([]float64, error) {
	row := make([]float64, len(parts))

	for i, part := range parts {
		val, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		row[i] = val
	}

	return row, nil
}

func rows(matrix [][]float64) int {
	return len(matrix)
}

func cols(matrix [][]float64, i int) int {
	return len(matrix[i])
}

func isEmpty(matrix [][]float64) bool {
	return rows(matrix) == 0
}

func isEdge(i float64) bool {
	return i != NoEdge
}

func isSquare(matrix [][]float64) bool {
	for i := range matrix {
		if cols(matrix, i) != rows(matrix) {
			return false
		}
	}
	return true
}

func (g *Graph) LoadGraphFromFile(filename string) error {
	return g.LoadFromFile(filename)
}

func (g *Graph) ExportGraphToDot(filename string) error {
	return g.ExportToDot(filename)
}
