package main

import (
	"fmt"
	"github.com/whiterage/simpleNavigatorGolang/algorithms"
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"strings"
	"time"
)

type appState struct {
	graph *graph.Graph
	file  string
}

func (s *appState) isGraphLoaded() bool {
	return s.graph != nil
}

func handleLoadGraph() (*graph.Graph, string) {
	path := readLine("Введите путь к файлу")
	if path == "" {
		printError("путь не может быть пустым")
		return nil, ""
	}

	g := graph.NewGraph()
	if err := g.LoadGraphFromFile(path); err != nil {
		printError(err.Error())
		return nil, ""
	}
	printResult("Граф успешно загружен", fmt.Sprintf("Вершин: %d", g.Vertex()))
	waitForEnter()
	return g, path
}

func handleBFS(g *graph.Graph, algs *algorithms.GraphAlgs) {
	begin := readInt("Начальная вершина")

	result, err := algs.BreadthFirstSearch(g, begin)
	if err != nil {
		printError(err.Error())
		return
	}

	printResult("BFS", formatPath(result))
	waitForEnter()
}

func handleDFS(g *graph.Graph, algs *algorithms.GraphAlgs) {
	begin := readInt("Начальная вершина")

	result, err := algs.DepthFirstSearch(g, begin)
	if err != nil {
		printError(err.Error())
		return
	}
	printResult("DFS", formatPath(result))
	waitForEnter()
}

func handleACO(g *graph.Graph, algs *algorithms.GraphAlgs) {
	result, err := algs.ACO(g)
	if err != nil {
		printError(err.Error())
		return
	}

	printResult("Задача коммивояжера", fmt.Sprintf("%s\nДлина: %v", formatPath(result.Vertices), result.Distance))
	waitForEnter()
}

func handleDijkstra(g *graph.Graph, algs *algorithms.GraphAlgs) {
	v1 := readInt("Первая вершина")
	v2 := readInt("Вторая вершина")

	result, err := algs.GetShortestPathBetweenVertices(g, v1, v2)
	if err != nil {
		printError(err.Error())
		return
	}

	printResult("Алгоритм Дейкстры", fmt.Sprintf("%g", result))
	waitForEnter()
}

func handleFloydWarshall(g *graph.Graph, algs *algorithms.GraphAlgs) {
	result, err := algs.GetShortestPathsBetweenAllVertices(g)
	if err != nil {
		printError(err.Error())
		return
	}

	printResult("Алгоритм Флойда-Уоршелла", formatMatrix(result))
	waitForEnter()
}

func handlePrim(g *graph.Graph, algs *algorithms.GraphAlgs) {
	result, err := algs.GetLeastSpanningTree(g)
	if err != nil {
		printError(err.Error())
		return
	}

	printResult("Алгоритм Прима", formatMatrix(result))
	waitForEnter()
}

func formatPath(vertices []int) string {
	parts := make([]string, len(vertices))
	for i, v := range vertices {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, " -> ")
}

func formatMatrix(matrix [][]float64) string {
	if len(matrix) == 0 {
		return ""
	}
	var rows []string
	for i := range matrix {
		var rowItems []string
		for _, v := range matrix[i] {
			rowItems = append(rowItems, fmt.Sprintf("%g", v))
		}
		rows = append(rows, strings.Join(rowItems, " "))
	}
	return strings.Join(rows, "\n")
}

func measureTSP(name string, runs int, fn func() (algorithms.TsmResult, error)) string {
	start := time.Now()

	for i := 0; i < runs; i++ {
		if _, err := fn(); err != nil {
			return fmt.Sprintf("%s: ошибка: %v", name, err)
		}
	}

	total := time.Since(start)
	avg := total / time.Duration(runs)

	return fmt.Sprintf("%s\nЗапусков: %d\nОбщее время: %v\nСреднее время: %v", name, runs, total, avg)
}

func handleTSPComparison(g *graph.Graph, algs *algorithms.GraphAlgs) {
	runs := readInt("Количество запусков N")
	if runs <= 0 {
		printError("количество запусков должно быть больше 0")
		return
	}
	fmt.Println(muted.Render("Выполняется сравнение алгоритмов, это может занять время..."))
	aco := measureTSP("Муравьиный алгоритм", runs, func() (algorithms.TsmResult, error) {
		return algs.SolveTravelingSalesmanProblem(g)
	})

	nearest := measureTSP("Nearest Neighbor", runs, func() (algorithms.TsmResult, error) {
		return algs.SolveTravelingSalesmanProblemNearestNeighbor(g)
	})

	bruteForce := measureTSP("Brute Force", runs, func() (algorithms.TsmResult, error) {
		return algs.SolveTravelingSalesmanProblemBruteForce(g)
	})

	printResult(
		"Сравнение алгоритмов TSP",
		aco+"\n\n"+nearest+"\n\n"+bruteForce,
	)

	waitForEnter()
}
