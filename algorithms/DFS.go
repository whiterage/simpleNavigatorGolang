package algorithms

import (
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"github.com/whiterage/simpleNavigatorGolang/stack"
)

// DFS выполняет поиск в глубину
// Возвращает срез вершин в порядке обхода
func (ga *GraphAlgs) DFS(g *graph.Graph, start int) ([]int, error) {
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

	stack := stack.NewStack()
	stack.Push(start)

	for !stack.IsEmpty() {
		cur, err := stack.Pop()
		if err != nil {
			return nil, err
		}

		if !been[cur] {
			res = append(res, cur+1)
			been[cur] = true

			for near := range g.Vertex() {
				if matrix[cur][near] != graph.NoEdge && !been[near] {
					stack.Push(near)
				}
			}
		}
	}

	return res, nil
}
