package algorithms

import (
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"github.com/whiterage/simpleNavigatorGolang/queue"
)

// BFS выполняет поиск в ширину
// Возвращает срез вершин в порядке обхода
func (ga *GraphAlgs) BFS(g *graph.Graph, start int) ([]int, error) {
	if g == nil || g.Vertex() == 0 {
		return nil, ErrEmptyGraph
	}

	if start < 1 || start > g.Vertex() {
		return nil, ErrVertex
	}

	start--
	been := make([]bool, g.Vertex())
	res := make([]int, 0, g.Vertex())
	matrix := g.Matrix()

	queue := queue.NewQueue()
	queue.Push(start)
	been[start] = true

	for !queue.IsEmpty() {
		cur, err := queue.Pop()
		if err != nil {
			return nil, err
		}

		res = append(res, cur+1)

		for near := range g.Vertex() {
			if matrix[cur][near] != graph.NoEdge && !been[near] {
				been[near] = true
				queue.Push(near)
			}
		}
	}

	return res, nil
}
