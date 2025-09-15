package http

import (
	"log"
	"os"
	"sort"
)

func updateDomainMetrics(domain string) {
	metricsMu.Lock()
	defer metricsMu.Unlock()

	if count, ok := metricsStore[domain]; ok {
		metricsStore[domain] = count + 1
	} else {
		metricsStore[domain] = 1
	}
}

func getTopDomains(n int) []map[string]any {
	metricsMu.Lock()
	defer metricsMu.Unlock()

	// Convert map to slice for sorting
	type kv struct {
		Domain string
		Count  int
	}
	var sorted []kv
	for d, c := range metricsStore {
		sorted = append(sorted, kv{Domain: d, Count: c})
	}

	// Sort by count (descending)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count > sorted[j].Count
	})

	// Pick top n
	var result []map[string]any
	limit := n
	if len(sorted) < n {
		limit = len(sorted)
	}
	for i := 0; i < limit; i++ {
		result = append(result, map[string]interface{}{
			"domain": sorted[i].Domain,
			"count":  sorted[i].Count,
		})
	}
	return result
}

func getBaseURL() string {
	if val := os.Getenv("BASE_URL"); val != "" {
		log.Fatal("BASE_URL is Missing")
		return val
	}
	return "http://localhost:8080"
}
