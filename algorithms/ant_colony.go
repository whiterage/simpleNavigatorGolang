package algorithms

import (
	"errors"
	"github.com/whiterage/simpleNavigatorGolang/graph"
	"math"
	"math/rand"
)

type TsmResult struct {
	Vertices []int
	Distance float64
}

var (
	NoEdgeAnt = math.Inf(1)
)

// Метод, который решает задачу коммивояжера
// В данной имплементации возвращает результат муравьиного алгоритма
func (ga *GraphAlgs) SolveTravelingSalesmanProblem(g *graph.Graph) (TsmResult, error) {
	return ga.ACO(g)
}

// ACO implementation ____________________________________________________

type Hyperparams struct {
	alpha       float64
	beta        float64
	evaporation float64
	tracks      [][]float64
	towns       int
}

type AntColony struct {
	ants      int
	pheromone [][]float64
	hp        Hyperparams
}

func NewAntColony(g *graph.Graph) (*AntColony, error) {
	if g == nil {
		return nil, ErrEmptyGraph
	}

	vertex := g.Vertex()
	if vertex == 0 {
		return nil, ErrVertex
	}

	if g.HasNegativeWeight() {
		return nil, errors.New("negative weights not allowed for ACO")
	}

	matrix := g.Matrix()
	for i := range vertex {
		for j := range vertex {
			if i == j || matrix[i][j] == graph.NoEdge {
				matrix[i][j] = NoEdgeAnt
			}
		}
	}

	ants := min(vertex, 200)

	pheromone := make([][]float64, vertex)
	for i := range pheromone {
		pheromone[i] = make([]float64, vertex)
		for j := range pheromone[i] {
			pheromone[i][j] = 1.0
		}
	}

	hp := Hyperparams{
		alpha:       1.0,
		beta:        2.0,
		evaporation: 0.1,
		tracks:      matrix,
		towns:       vertex,
	}

	return &AntColony{
		ants:      ants,
		pheromone: pheromone,
		hp:        hp,
	}, nil
}

func (ac *AntColony) Run(iters int) ([]int, float64) {
	bestPath := make([]int, 0)
	bestDist := NoEdgeAnt

	for range iters {
		iterTours := make([][]int, ac.ants)
		iterTrails := make([]float64, ac.ants)

		for ant := range ac.ants {
			tour, trail := ac.antTour()
			iterTours[ant], iterTrails[ant] = tour, trail

			if tour != nil && trail < bestDist {
				bestPath = make([]int, len(tour))
				copy(bestPath, tour)
				bestDist = trail
			}
		}

		ac.freshPheromone(iterTours, iterTrails)
	}

	return bestPath, bestDist
}

func (ac *AntColony) antTour() ([]int, float64) {
	tour := make([]int, 0)
	been := make([]bool, ac.hp.towns)

	start := rand.Intn(ac.hp.towns)
	tour = append(tour, start)
	been[start] = true

	trail := 0.0
	cur := start

	for len(tour) < ac.hp.towns {
		next := ac.pickTown(cur, been)
		if next == -1 {
			return nil, NoEdgeAnt
		}

		tour = append(tour, next)
		trail += ac.hp.tracks[cur][next]

		been[next] = true
		cur = next
	}

	if ac.hp.tracks[cur][start] == NoEdgeAnt {
		return nil, NoEdgeAnt
	}

	tour = append(tour, start)
	trail += ac.hp.tracks[cur][start]

	return tour, trail
}

func (ac *AntColony) freshPheromone(tours [][]int, trails []float64) {
	for i := range ac.hp.towns {
		for j := range ac.hp.towns {
			ac.pheromone[i][j] *= (1 - ac.hp.evaporation)
		}
	}

	for iter := range len(tours) {
		if math.IsInf(trails[iter], 1) {
			continue
		}

		delta := 1.0 / trails[iter]
		tour := tours[iter]

		for i := range len(tour) - 1 {
			cur, next := tour[i], tour[i+1]
			ac.pheromone[cur][next] += delta
			// граф неориентированный поэтому обновляем для обоих направлений
			ac.pheromone[next][cur] += delta
		}
	}
}

func (ac *AntColony) pickTown(cur int, been []bool) int {
	pick := make([]float64, ac.hp.towns)
	var choose float64

	for next := range ac.hp.towns {
		if !been[next] && ac.hp.tracks[cur][next] != NoEdgeAnt {
			nicePheromone := math.Pow(ac.pheromone[cur][next], ac.hp.alpha)
			niceTrack := math.Pow(1.0/ac.hp.tracks[cur][next], ac.hp.beta)
			pick[next] = nicePheromone * niceTrack
			choose += pick[next]
		}
	}

	if choose == 0 {
		return -1
	}

	chance := rand.Float64() * choose
	sectorGate := 0.0
	for town := range ac.hp.towns {
		if pick[town] > 0 {
			sectorGate += pick[town]
			if chance <= sectorGate {
				return town
			}
		}
	}

	return -1
}

// Муравьиный алгоритм
// Возвращает структуру ответа для решения проблемы коммивояжера
func (ga *GraphAlgs) ACO(g *graph.Graph) (TsmResult, error) {

	colony, err := NewAntColony(g)
	if err != nil {
		return TsmResult{}, err
	}

	iters := 500
	bestPath, bestDist := colony.Run(iters)
	if bestDist == NoEdgeAnt {
		return TsmResult{}, errors.New("can't find path")
	}

	// 1-базовая индексацию
	convPath := make([]int, len(bestPath))
	for i, v := range bestPath {
		convPath[i] = v + 1
	}

	return TsmResult{Vertices: convPath, Distance: bestDist}, nil
}
