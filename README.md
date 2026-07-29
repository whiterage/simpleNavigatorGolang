# Simple Navigator

Graph algorithms implemented from scratch in Go, with an interactive terminal UI.

No graph libraries are used — the graph, the traversal structures (stack and queue), and every
algorithm below are written by hand and covered by tests.

[Русская версия](README.ru.md)

---

## Algorithms

| Task | Algorithm | Complexity |
|---|---|---|
| Traversal | Breadth-first search | O(V²) |
| Traversal | Depth-first search | O(V²) |
| Shortest path between two vertices | Dijkstra | O(V²) |
| Shortest paths between all vertices | Floyd–Warshall | O(V³) |
| Minimum spanning tree | Prim | O(V³) |
| Travelling salesman | Ant colony optimization | O(iters · ants · V²) |
| Travelling salesman | Nearest neighbour (greedy) | O(V²) |
| Travelling salesman | Brute force (exact) | O(V!) |

Complexities reflect the adjacency-matrix representation used throughout: Dijkstra and Prim pick the
next vertex by linear scan rather than with a heap, which is the better trade-off on dense matrices.

BFS and DFS are built on the project's own `queue` and `stack` packages rather than on slices, so the
traversal logic reads the same way the textbook algorithm does.

---

## Quick start

Requires Go 1.22 or newer.

```bash
git clone https://github.com/whiterage/simpleNavigatorGolang.git
cd simpleNavigatorGolang
make run
```

Other targets:

```bash
make build          # build ./bin/simple_navigator
make test           # run all tests
make test-coverage  # run tests with coverage
make fmt vet        # format and vet
make clean          # remove build output
```

---

## Using the app

The app opens a menu; pick an entry with the arrow keys. Start with the first item, *load a graph
from file* — every other entry stays greyed out until a graph is loaded.

The interface is in Russian:

```
╭──────────────────────────────────╮
│ Simple Navigator                 │
│ ● граф загружен  assets/tsp_8.txt │
╰──────────────────────────────────╯
  ↑/↓ - navigate / enter - select
  ▸ Загрузить граф из файла          load graph from file
    Обход в ширину (BFS)             breadth-first search
    Обход в глубину (DFS)            depth-first search
    Алгоритм Дейкстры                Dijkstra
    Алгоритм Флойда-Уоршелла         Floyd–Warshall
    Алгоритм Прима                   Prim
    Задача коммивояжера              travelling salesman
    Сравнение алгоритмов TSP         compare TSP algorithms
    Выход                            exit
```

Three sample graphs ship in [`assets/`](assets):

| File | Description |
|---|---|
| `graph.txt` | 5 vertices, sparse, some one-way edges |
| `tsp_4.txt` | 4 vertices, complete — small enough to check brute force by hand |
| `tsp_8.txt` | 8 vertices, complete — the interesting case for the TSP comparison |

The last menu entry runs all three TSP solvers N times on the loaded graph and reports total and
average time for each, which is the direct way to see the exact solver fall behind the heuristics as
the graph grows.

---

## Graph file format

A plain-text square adjacency matrix. Row `i`, column `j` is the weight of the edge from vertex `i`
to vertex `j`; `0` means no edge. Values are whitespace-separated, blank lines are ignored.

```
0 1 5 10
1 0 2 4
5 2 0 3
10 4 3 0
```

Vertices are numbered from 1 in the API and the UI.

---

## Using it as a library

```go
package main

import (
	"fmt"
	"log"

	"github.com/whiterage/simpleNavigatorGolang/algorithms"
	"github.com/whiterage/simpleNavigatorGolang/graph"
)

func main() {
	g := graph.NewGraph()
	if err := g.LoadFromFile("assets/tsp_8.txt"); err != nil {
		log.Fatal(err)
	}

	algs := algorithms.NewGraphAlgs()

	order, err := algs.BFS(g, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BFS order:", order)

	tour, err := algs.SolveTravelingSalesmanProblem(g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tour %v, length %.2f\n", tour.Vertices, tour.Distance)

	if err := g.ExportToDot("graph.dot"); err != nil {
		log.Fatal(err)
	}
}
```

`ExportToDot` writes a Graphviz file, so a loaded graph can be rendered with
`dot -Tpng graph.dot -o graph.png`.

### API

**`graph`**

| Method | Returns |
|---|---|
| `NewGraph()` | empty graph |
| `LoadFromFile(path)` | error — rejects empty, non-square and unparsable input |
| `ExportToDot(path)` | error — writes Graphviz DOT |
| `Matrix()` | defensive copy of the adjacency matrix |
| `Vertex()` | vertex count |
| `HasNegativeWeight()` | whether any weight is negative |

**`algorithms`** — methods on `GraphAlgs`

| Method | Returns |
|---|---|
| `BFS(g, start)` / `DFS(g, start)` | `[]int` — vertices in visit order |
| `GetShortestPathBetweenVertices(g, from, to)` | `float64` — distance |
| `GetShortestPathsBetweenAllVertices(g)` | `[][]float64` — distance matrix |
| `GetLeastSpanningTree(g)` | `[][]float64` — MST as an adjacency matrix |
| `SolveTravelingSalesmanProblem(g)` | `TsmResult` — ant colony |
| `SolveTravelingSalesmanProblemNearestNeighbor(g)` | `TsmResult` — greedy |
| `SolveTravelingSalesmanProblemBruteForce(g)` | `TsmResult` — exact |

`TsmResult` holds the tour as `Vertices []int` (starting and ending at the same vertex) and its total
`Distance float64`.

**`queue`** and **`stack`** are small `int` containers: `Push`, `Pop`, `Size`, `IsEmpty`, plus
`Front`/`Back` on the queue and `Top` on the stack.

---

## Behaviour worth knowing

- Dijkstra, Floyd–Warshall, Prim and the TSP solvers reject graphs with negative weights.
- Brute-force TSP refuses graphs with more than 10 vertices — 10! is already the practical ceiling.
- Prim returns `ErrDisconnectedGraph` when no spanning tree exists.
- Ant colony optimization is stochastic, so repeated runs may return different tours. It runs 500
  iterations with `min(V, 200)` ants, `alpha = 1.0`, `beta = 2.0`, evaporation `0.1`.

---

## Tests

```bash
make test-coverage
```

| Package | Coverage |
|---|---|
| `graph` | 91.4% |
| `algorithms` | 62.8% |
| `queue` | 100% |
| `stack` | 100% |

---

## Project layout

```
cmd/simple_navigator/   terminal UI — menu, prompts, result rendering
graph/                  adjacency-matrix graph, file loading, DOT export
algorithms/             BFS, DFS, Dijkstra, Floyd–Warshall, Prim, three TSP solvers
queue/                  FIFO queue used by BFS
stack/                  LIFO stack used by DFS
assets/                 sample graphs
```

Built with [promptui](https://github.com/manifoldco/promptui) for the menu and
[lipgloss](https://github.com/charmbracelet/lipgloss) for styling.
