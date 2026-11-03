package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	cpuSuffix = regexp.MustCompile(`-\d+$`)
	codecs    = []string{
		"encoding-json-v2/default-safe",
		"encoding-json-v2/deterministic",
		"encoding-json-v2/fast-relaxed",
		"encoding-json",
		"jsoniter/compat",
		"jsoniter/fast",
		"go-json/compat",
		"go-json/fast",
		"segmentio/compat",
		"segmentio/fast",
		"sonic/compat",
		"sonic/fast",
		"easyjson/compat",
		"easyjson/fast",
	}
)

type sample struct {
	ns     float64
	bytes  float64
	allocs float64
	mbps   float64
	output float64
}

func main() {
	baseline := flag.String("baseline", "encoding-json", "baseline codec name for ratios")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: report [-baseline name] benchmark.txt [...]")
		os.Exit(2)
	}

	all := map[string][]sample{}
	for _, path := range flag.Args() {
		if err := readFile(path, all); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	keys := make([]string, 0, len(all))
	for key := range all {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Println("# JSON benchmark özeti")
	fmt.Println()
	fmt.Printf("Baz: `%s`. Hız >1 daha hızlı; bellek/alloc <1 daha iyi.\n", *baseline)
	fmt.Println("Değerler: ortalama ± standart sapma (count≥2).")
	fmt.Println()
	fmt.Println("| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op | output B | hız x | bellek x | alloc x | n |")
	fmt.Println("|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|")
	for _, key := range keys {
		group, codecName := splitKey(key)
		vals := all[key]
		cur := meanStd(vals)
		baseValues, ok := all[group+"/"+*baseline]
		speed, memory, allocations := "-", "-", "-"
		if ok {
			base := meanStd(baseValues)
			speed = ratio(base.mean.ns, cur.mean.ns)
			memory = ratio(cur.mean.bytes, base.mean.bytes)
			allocations = ratio(cur.mean.allocs, base.mean.allocs)
		}
		fmt.Printf("| `%s` | `%s` | %s | %s | %s | %s | %s | %s | %s | %s | %d |\n",
			group, codecName,
			fmtMetric(cur.mean.ns, cur.std.ns, 0),
			fmtMetric(cur.mean.mbps, cur.std.mbps, 1),
			fmtMetric(cur.mean.bytes, cur.std.bytes, 0),
			fmtMetric(cur.mean.allocs, cur.std.allocs, 0),
			fmtMetric(cur.mean.output, cur.std.output, 0),
			speed, memory, allocations, len(vals))
	}
}

type stats struct {
	mean sample
	std  sample
}

func meanStd(values []sample) stats {
	n := float64(len(values))
	if n == 0 {
		return stats{}
	}
	var sum sample
	for _, v := range values {
		sum.ns += v.ns
		sum.bytes += v.bytes
		sum.allocs += v.allocs
		sum.mbps += v.mbps
		sum.output += v.output
	}
	mean := sample{
		ns: sum.ns / n, bytes: sum.bytes / n, allocs: sum.allocs / n,
		mbps: sum.mbps / n, output: sum.output / n,
	}
	if n < 2 {
		return stats{mean: mean}
	}
	var acc sample
	for _, v := range values {
		acc.ns += sq(v.ns - mean.ns)
		acc.bytes += sq(v.bytes - mean.bytes)
		acc.allocs += sq(v.allocs - mean.allocs)
		acc.mbps += sq(v.mbps - mean.mbps)
		acc.output += sq(v.output - mean.output)
	}
	den := n - 1
	return stats{
		mean: mean,
		std: sample{
			ns: math.Sqrt(acc.ns / den), bytes: math.Sqrt(acc.bytes / den),
			allocs: math.Sqrt(acc.allocs / den), mbps: math.Sqrt(acc.mbps / den),
			output: math.Sqrt(acc.output / den),
		},
	}
}

func sq(v float64) float64 { return v * v }

func fmtMetric(mean, std float64, prec int) string {
	if std == 0 {
		return strconv.FormatFloat(mean, 'f', prec, 64)
	}
	return fmt.Sprintf("%.*f±%.*f", prec, mean, prec, std)
}

func readFile(path string, all map[string][]sample) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 2*1024*1024)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 4 || !strings.HasPrefix(fields[0], "Benchmark") {
			continue
		}
		name := cpuSuffix.ReplaceAllString(fields[0], "")
		group, codecName := splitKey(name)
		if codecName == "" {
			continue
		}
		s := sample{}
		for i := 2; i+1 < len(fields); i += 2 {
			value, parseErr := strconv.ParseFloat(fields[i], 64)
			if parseErr != nil {
				continue
			}
			switch fields[i+1] {
			case "ns/op":
				s.ns = value
			case "MB/s":
				s.mbps = value
			case "B/op":
				s.bytes = value
			case "allocs/op":
				s.allocs = value
			case "output-B/op":
				s.output = value
			}
		}
		if s.ns > 0 {
			all[group+"/"+codecName] = append(all[group+"/"+codecName], s)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}

func splitKey(key string) (string, string) {
	// Longest suffix wins so encoding-json-v2/... beats encoding-json.
	best := ""
	for _, codecName := range codecs {
		suffix := "/" + codecName
		if strings.HasSuffix(key, suffix) && len(codecName) > len(best) {
			best = codecName
		}
	}
	if best == "" {
		return key, ""
	}
	return strings.TrimSuffix(key, "/"+best), best
}

func ratio(numerator, denominator float64) string {
	if denominator == 0 {
		return "-"
	}
	return fmt.Sprintf("%.2f", numerator/denominator)
}
