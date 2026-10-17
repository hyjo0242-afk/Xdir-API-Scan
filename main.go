package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	urlFlag := flag.String("u", "", "Target URL (e.g., https://example.com)")
	jsonFlag := flag.Bool("json", false, "Save output as JSON")
	outputFlag := flag.String("o", "xdir-output.json", "JSON output filename")
	noBanner := flag.Bool("n", false, "Disable banner")
	flag.Parse()

	if !*noBanner { printBanner() }

	if *urlFlag == "" {
		fmt.Println("Usage: xdir -u https://example.com")
		os.Exit(1)
	}

	target := *urlFlag
	if len(target) < 4 || target[:4] != "http" {
		target = "https://" + target
	}

	startTime := time.Now()
	groups, err := extractFromJS(target)
	if err != nil {
		fmt.Printf("[!] Error: %v\n", err)
		os.Exit(1)
	}
	elapsed := time.Since(startTime)

	printResults(target, groups)

	if *jsonFlag {
		data := map[string]interface{}{
			"target":    target,
			"timestamp": time.Now().Format(time.RFC3339),
			"endpoints": groups,
			"duration":  elapsed.String(),
		}
		file, _ := json.MarshalIndent(data, "", "  ")
		os.WriteFile(*outputFlag, file, 0644)
		fmt.Printf("\n[+] Saved to %s\n", *outputFlag)
	}
}
