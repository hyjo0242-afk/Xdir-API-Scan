package main

import (
	"regexp"
	"strings"
)

func findEndpoints(jsContent string) []string {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`["'](/api/[a-zA-Z0-9_/-]+)["']`),
		regexp.MustCompile(`["'](/v[0-9]/[a-zA-Z0-9_/-]+)["']`),
		regexp.MustCompile(`["'](/graphql[a-zA-Z0-9_/-]*)["']`),
		regexp.MustCompile(`["'](/rpc/[a-zA-Z0-9_/-]+)["']`),
	}

	seen := make(map[string]bool)
	var endpoints []string

	for _, re := range patterns {
		matches := re.FindAllStringSubmatch(jsContent, -1)
		for _, match := range matches {
			if len(match) < 2 { continue }
			ep := match[1]
			if len(ep) < 5 || strings.Contains(ep, "..") { continue }
			if !seen[ep] {
				seen[ep] = true
				endpoints = append(endpoints, ep)
			}
		}
	}
	return endpoints
}

func groupEndpoints(endpoints map[string]bool) map[string][]string {
	groups := make(map[string][]string)
	for ep := range endpoints {
		category := "OTHER"
		switch {
		case strings.Contains(ep, "/auth/"): category = "AUTH"
		case strings.Contains(ep, "/users/") || strings.Contains(ep, "/me"): category = "USERS"
		case strings.Contains(ep, "/chats/") || strings.Contains(ep, "/messages/"): category = "CHATS"
		case strings.Contains(ep, "/servers/") || strings.Contains(ep, "/channels/"): category = "SERVERS"
		case strings.Contains(ep, "/admin/"): category = "ADMIN"
		case strings.Contains(ep, "/upload"): category = "UPLOADS"
		case strings.Contains(ep, "/search"): category = "SEARCH"
		}
		groups[category] = append(groups[category], ep)
	}
	return groups
}
