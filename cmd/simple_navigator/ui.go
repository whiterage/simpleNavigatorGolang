package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/manifoldco/promptui"
)

var (
	muted     = lipgloss.NewStyle().Foreground(lipgloss.Color("#928374"))
	success   = lipgloss.NewStyle().Foreground(lipgloss.Color("#b8bb26"))
	warn      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fabd2f"))
	val       = lipgloss.NewStyle().Foreground(lipgloss.Color("#d3869b"))
	header    = lipgloss.NewStyle().Foreground(lipgloss.Color("#83a598")).Bold(true)
	resultBox = box.BorderForeground(lipgloss.Color("#bb9af7"))
	err       = lipgloss.NewStyle().Foreground(lipgloss.Color("#fb4934"))
	errBox    = box.BorderForeground(lipgloss.Color("#fb4934"))
	box       = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#504945")).
			Padding(0, 2)
)

func readLine(prompt string) string {
	p := promptui.Prompt{
		Label: prompt,
	}

	result, err := p.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(result)
}

func readInt(prompt string) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)

		if err != nil {
			return fmt.Errorf("введите целое число")
		}

		return nil
	}

	p := promptui.Prompt{
		Label:    prompt,
		Validate: validate,
	}

	result, err := p.Run()
	if err != nil {
		return 0
	}

	val, _ := strconv.Atoi(result)

	return val
}

func printHeader(graphLoaded bool, graphFile string) {
	status := warn.Render("○ граф не загружен")

	if graphLoaded {
		status = success.Render("● граф загружен") +
			"  " +
			muted.Render(graphFile)
	}

	fmt.Println(
		box.Render(
			header.Render("Simple Navigator") +
				"\n" +
				status,
		),
	)
}

func printResult(label, value string) {
	fmt.Println(resultBox.Render(muted.Render(label) + "\n" + val.Render(value)))
}

func printError(str string) {
	fmt.Println(errBox.Render(err.Render("✗ " + str)))
	waitForEnter()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func waitForEnter() {
	prompt := promptui.Prompt{
		Label: "Нажмите Enter для продолжения",
	}
	_, _ = prompt.Run()
}

func runMenu(graphLoaded bool) int {
	type item struct {
		id         int
		label      string
		needsGraph bool
	}

	items := []item{
		{1, "Загрузить граф из файла", false},
		{2, "Обход в ширину (BFS)", true},
		{3, "Обход в глубину (DFS)", true},
		{4, "Алгоритм Дейкстры", true},
		{5, "Алгоритм Флойда-Уоршелла", true},
		{6, "Алгоритм Прима", true},
		{7, "Задача коммивояжера", true},
		{8, "Сравнение алгоритмов TSP", true},
		{0, "Выход", false},
	}

	labels := make([]string, len(items))
	for i, it := range items {
		if it.needsGraph && !graphLoaded {
			labels[i] = muted.Render(it.label)
		} else {
			labels[i] = it.label
		}
	}

	sel := promptui.Select{
		Label: muted.Render("↑/↓ - navigate / enter - select"),
		Items: labels,
		Size:  len(items),
	}

	idx, _, err := sel.Run()
	if err != nil {
		return 0
	}

	if items[idx].needsGraph && !graphLoaded {
		printError("сначала загрузите граф")
		return -1
	}

	return items[idx].id
}
