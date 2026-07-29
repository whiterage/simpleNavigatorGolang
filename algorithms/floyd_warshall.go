package algorithms

import (
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"math"
)

func (ga *GraphAlgs) GetShortestPathsBetweenAllVertices(g *graph.Graph) ([][]float64, error) {
	if g == nil || g.Vertex() == 0 {
		return nil, ErrEmptyGraph
	}
	if g.HasNegativeWeight() {
		return nil, graph.ErrNegWeight
	}

	n := g.Vertex()
	matrix := g.Matrix()

	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else if matrix[i][j] != graph.NoEdge {
				dist[i][j] = matrix[i][j]
			} else {
				dist[i][j] = math.Inf(1)
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist, nil
}
