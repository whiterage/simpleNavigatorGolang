package main

import (
	"github.com/whiterage/simpleNavigatorGolang/algorithms"
)

func main() {
	var state appState
	algs := &algorithms.GraphAlgs{}

	for {
		clearScreen()
		printHeader(state.isGraphLoaded(), state.file)

		switch runMenu(state.isGraphLoaded()) {
		case 1:
			state.graph, state.file = handleLoadGraph()
		case 2:
			handleBFS(state.graph, algs)
		case 3:
			handleDFS(state.graph, algs)
		case 4:
			handleDijkstra(state.graph, algs)
		case 5:
			handleFloydWarshall(state.graph, algs)
		case 6:
			handlePrim(state.graph, algs)
		case 7:
			handleACO(state.graph, algs)
		case 8:
			handleTSPComparison(state.graph, algs)
		case 0:
			return
		}
	}
}
