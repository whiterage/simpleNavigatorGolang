package algorithms

import (
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"math"
)

func (ga *GraphAlgs) GetShortestPathBetweenVertices(g *graph.Graph, vertex1, vertex2 int) (float64, error) {
	if g == nil || g.Vertex() == 0 {
		return 0, ErrEmptyGraph
	}
	if vertex1 < 1 || vertex1 > g.Vertex() {
		return 0, ErrVertex
	}
	if vertex2 < 1 || vertex2 > g.Vertex() {
		return 0, ErrVertex
	}
	if g.HasNegativeWeight() {
		return 0, graph.ErrNegWeight
	}

	start := vertex1 - 1
	finish := vertex2 - 1

	n := g.Vertex()

	dist := make([]float64, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = math.Inf(1)
	}

	dist[start] = 0
	matrix := g.Matrix()

	for i := 0; i < n; i++ {
		cur := -1
		bestDist := math.Inf(1)

		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < bestDist {
				bestDist = dist[v]
				cur = v
			}
		}

		if cur == -1 {
			break
		}

		visited[cur] = true

		for near := 0; near < n; near++ {
			weight := matrix[cur][near]

			if weight == graph.NoEdge || visited[near] {
				continue
			}

			if dist[cur]+weight < dist[near] {
				dist[near] = dist[cur] + weight
			}
		}
	}

	if math.IsInf(dist[finish], 1) {
		return 0, graph.ErrPathNotFound
	}

	return dist[finish], nil
}
