package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type request struct {
	status  uint64
	bytes   uint64
	latency uint64
}

func decimal(field string) uint64 {
	value, err := strconv.ParseUint(field, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func readRequests(input io.Reader) (map[string][]request, error) {
	// Keep each request so reading and summarizing are separate steps.
	groups := make(map[string][]request)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		endpoint := fields[1]
		groups[endpoint] = append(groups[endpoint], request{
			status:  decimal(fields[2]),
			bytes:   decimal(fields[3]),
			latency: decimal(fields[4]),
		})
	}
	return groups, scanner.Err()
}

func writeSummary(output io.Writer, groups map[string][]request) error {
	writer := bufio.NewWriter(output)
	for _, endpoint := range slices.Sorted(maps.Keys(groups)) {
		requests := groups[endpoint]
		var errors, bytes, maxLatency uint64
		for _, request := range requests {
			if request.status >= 500 {
				errors++
			}
			bytes += request.bytes
			maxLatency = max(maxLatency, request.latency)
		}

		_, err := fmt.Fprintf(
			writer,
			"%s %d %d %d %d\n",
			endpoint,
			len(requests),
			errors,
			bytes,
			maxLatency,
		)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	groups, err := readRequests(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := writeSummary(os.Stdout, groups); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
