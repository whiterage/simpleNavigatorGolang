package algorithms

import (
	"errors"
	"github.com/whiterage/simpleNavigatorGolang/graph"
)

func (ga *GraphAlgs) SolveTravelingSalesmanProblemNearestNeighbor(g *graph.Graph) (TsmResult, error) {
	if g == nil || g.Vertex() == 0 {
		return TsmResult{}, graph.ErrEmptyFile
	}

	if g.HasNegativeWeight() {
		return TsmResult{}, graph.ErrNegWeight
	}

	n := g.Vertex()
	matrix := g.Matrix()

	visited := make([]bool, n)
	path := make([]int, 0, n+1)

	start := 0
	current := start
	visited[current] = true
	path = append(path, current)

	distance := 0.0

	for step := 1; step < n; step++ {
		next := -1
		bestWeight := 0.0

		for candidate := 0; candidate < n; candidate++ {
			weight := matrix[current][candidate]

			if visited[candidate] || weight == graph.NoEdge {
				continue
			}

			if next == -1 || weight < bestWeight {
				next = candidate
				bestWeight = weight
			}
		}

		if next == -1 {
			return TsmResult{}, errors.New("can't find path")
		}

		visited[next] = true
		path = append(path, next)
		distance += bestWeight
		current = next
	}

	backWeight := matrix[current][start]
	if backWeight == graph.NoEdge {
		return TsmResult{}, errors.New("can't return to start")
	}

	path = append(path, start)
	distance += backWeight

	convertedPath := make([]int, len(path))
	for i, v := range path {
		convertedPath[i] = v + 1
	}

	return TsmResult{
		Vertices: convertedPath,
		Distance: distance,
	}, nil
}
