package algorithms

import (
	"errors"
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"math"
)

func calculateTourDistance(matrix [][]float64, path []int) (float64, bool) {
	distance := 0.0

	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]
		weight := matrix[from][to]

		if weight == graph.NoEdge {
			return 0, false
		}

		distance += weight
	}

	return distance, true
}

func (ga *GraphAlgs) SolveTravelingSalesmanProblemBruteForce(g *graph.Graph) (TsmResult, error) {
	if g == nil || g.Vertex() == 0 {
		return TsmResult{}, ErrEmptyGraph
	}

	if g.HasNegativeWeight() {
		return TsmResult{}, graph.ErrNegWeight
	}

	n := g.Vertex()
	if n > 10 {
		return TsmResult{}, errors.New("brute force is too slow for graph with more than 10 vertices")
	}

	matrix := g.Matrix()

	bestDistance := math.Inf(1)
	bestPath := make([]int, 0, n+1)

	vertices := make([]int, 0, n-1)
	for i := 1; i < n; i++ {
		vertices = append(vertices, i)
	}

	var permute func(pos int)

	permute = func(pos int) {
		if pos == len(vertices) {
			path := make([]int, 0, n+1)
			path = append(path, 0)
			path = append(path, vertices...)
			path = append(path, 0)

			distance, ok := calculateTourDistance(matrix, path)
			if !ok {
				return
			}

			if distance < bestDistance {
				bestDistance = distance
				bestPath = make([]int, len(path))
				copy(bestPath, path)
			}

			return
		}

		for i := pos; i < len(vertices); i++ {
			vertices[pos], vertices[i] = vertices[i], vertices[pos]
			permute(pos + 1)
			vertices[pos], vertices[i] = vertices[i], vertices[pos]
		}
	}

	permute(0)
	if math.IsInf(bestDistance, 1) {
		return TsmResult{}, errors.New("can't find path")
	}
	convertedPath := make([]int, len(bestPath))
	for i, v := range bestPath {
		convertedPath[i] = v + 1
	}

	return TsmResult{
		Vertices: convertedPath,
		Distance: bestDistance,
	}, nil
}
