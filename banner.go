package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

const bannerSmall = `
  ██╗  ██╗██████╗ ██╗██████╗ 
  ╚██╗██╔╝██╔══██╗██║██╔══██╗
   ╚███╔╝ ██║  ██║██║██████╔╝
   ██╔██╗ ██║  ██║██║██╔══██╗
  ██╔╝ ██╗██████╔╝██║██║  ██║
  ╚═╝  ╚═╝╚═════╝ ╚═╝╚═╝  ╚═╝
`

var fadeSteps = []color.Attribute{
	color.FgHiBlack, color.FgBlack, color.FgWhite,
	color.FgHiWhite, color.FgHiCyan, color.FgCyan,
}

func printBanner() {
	lines := strings.Split(bannerSmall, "\n")
	for step, attr := range fadeSteps {
		c := color.New(attr)
		if step == 0 {
			fmt.Print("\033[H\033[2J")
		}
		fmt.Print("\033[H")
		for _, line := range lines {
			c.Println(line)
		}
		if step < len(fadeSteps)-1 {
			time.Sleep(150 * time.Millisecond)
		}
	}

	green := color.New(color.FgGreen, color.Bold)
	dim := color.New(color.FgHiBlack)
	cyan := color.New(color.FgCyan)

	fmt.Println()
	time.Sleep(200 * time.Millisecond)
	green.Printf("  v1.0.0 ")
	dim.Printf("— API Endpoint Extractor\n")
	time.Sleep(100 * time.Millisecond)
	dim.Println("  github.com/твой_ник/xdir")
	time.Sleep(100 * time.Millisecond)
	cyan.Println("\n  " + strings.Repeat("─", 40))
	fmt.Println()
}
