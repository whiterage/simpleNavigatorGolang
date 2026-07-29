package algorithms

import "github.com/whiterage/simpleNavigatorGolang/graph"

func (ga *GraphAlgs) DepthFirstSearch(g *graph.Graph, start int) ([]int, error) {
	return ga.DFS(g, start)
}

func (ga *GraphAlgs) BreadthFirstSearch(g *graph.Graph, start int) ([]int, error) {
	return ga.BFS(g, start)
}
