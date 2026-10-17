package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 15 * time.Second}

func fetchContent(targetURL string) (string, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func findJSFiles(baseURL string, html string) []string {
	re := regexp.MustCompile(`src=["']([^"']+\.js[^"']*)["']`)
	matches := re.FindAllStringSubmatch(html, -1)

	base, _ := url.Parse(baseURL)
	var jsFiles []string
	seen := make(map[string]bool)

	for _, match := range matches {
		if len(match) < 2 { continue }
		jsPath := match[1]
		if strings.Contains(jsPath, "google") || strings.Contains(jsPath, "facebook") { continue }

		jsURL, err := base.Parse(jsPath)
		if err != nil { continue }

		urlStr := jsURL.String()
		if !seen[urlStr] {
			seen[urlStr] = true
			jsFiles = append(jsFiles, urlStr)
		}
	}
	return jsFiles
}

func extractFromJS(targetURL string) (map[string][]string, error) {
	fmt.Printf("[*] Fetching %s...\n", targetURL)
	html, err := fetchContent(targetURL)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch HTML: %v", err)
	}

	jsFiles := findJSFiles(targetURL, html)
	fmt.Printf("[*] Found %d JS file(s)\n", len(jsFiles))

	allEndpoints := make(map[string]bool)

	for _, jsURL := range jsFiles {
		fmt.Printf("[*] Analyzing %s...\n", jsURL)
		jsContent, err := fetchContent(jsURL)
		if err != nil {
			fmt.Printf("[!] Failed: %v\n", err)
			continue
		}
		found := findEndpoints(jsContent)
		for _, ep := range found {
			allEndpoints[ep] = true
		}
	}

	return groupEndpoints(allEndpoints), nil
}
