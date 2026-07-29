package algorithms

import (
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"math"
)

func (ga *GraphAlgs) GetLeastSpanningTree(g *graph.Graph) ([][]float64, error) {
	if g == nil || g.Vertex() == 0 {
		return nil, ErrEmptyGraph
	}

	if g.HasNegativeWeight() {
		return nil, graph.ErrNegWeight
	}

	n := g.Vertex()
	matrix := g.Matrix()

	mst := make([][]float64, n)
	for i := range mst {
		mst[i] = make([]float64, n)
	}

	selected := make([]bool, n)
	selected[0] = true

	for edgeCount := 0; edgeCount < n-1; edgeCount++ {
		minWeight := math.Inf(1)
		from := -1
		to := -1

		for i := 0; i < n; i++ {
			if !selected[i] {
				continue
			}

			for j := 0; j < n; j++ {
				weight := matrix[i][j]

				if selected[j] || weight == graph.NoEdge {
					continue
				}

				if weight < minWeight {
					minWeight = weight
					from = i
					to = j
				}
			}
		}

		if from == -1 || to == -1 {
			return nil, ErrDisconnectedGraph
		}

		mst[from][to] = minWeight
		mst[to][from] = minWeight
		selected[to] = true
	}
	return mst, nil
}
