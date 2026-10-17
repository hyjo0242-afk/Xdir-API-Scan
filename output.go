package main

import (
	"fmt"
	"sort"

	"github.com/pterm/pterm"
)

func printResults(targetURL string, groups map[string][]string) {
	fmt.Println()
	pterm.DefaultHeader.WithFullWidth().
		WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
		Println("Xdir Results: " + targetURL)
	fmt.Println()

	var categories []string
	for cat := range groups { categories = append(categories, cat) }
	sort.Slice(categories, func(i, j int) bool {
		if categories[i] == "OTHER" { return false }
		if categories[j] == "OTHER" { return true }
		return categories[i] < categories[j]
	})

	total := 0
	for _, cat := range categories {
		endpoints := groups[cat]
		sort.Strings(endpoints)
		total += len(endpoints)

		pterm.DefaultSection.Println(fmt.Sprintf("%s (%d)", cat, len(endpoints)))
		var tableData [][]string
		for _, ep := range endpoints {
			tableData = append(tableData, []string{ep})
		}
		pterm.DefaultTable.WithData(tableData).Render()
		fmt.Println()
	}

	pterm.Success.Printf("Total endpoints found: %d\n", total)
}
