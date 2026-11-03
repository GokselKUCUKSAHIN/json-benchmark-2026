package jsonbenchmark

import (
	stdjson "encoding/json"
	"os"
	"strings"
	"testing"
)

var (
	bytesSink []byte
	valueSink any
)

func selectedCodecs() []codec {
	filter := os.Getenv("BENCH_CODEC")
	codecs := allCodecs()
	if filter == "" {
		return codecs
	}
	selected := codecs[:0]
	for _, c := range codecs {
		if strings.Contains(c.name, filter) {
			selected = append(selected, c)
		}
	}
	return selected
}

func newConcrete(ds dataset) any {
	switch ds.shape {
	case shapeFlat:
		return new(FlatDoc)
	case shapeNested:
		return new(NestedDoc)
	default:
		panic("unknown shape")
	}
}

func BenchmarkMarshalConcrete(b *testing.B) {
	for _, ds := range benchmarkDatasets() {
		b.Run(ds.name, func(b *testing.B) {
			for _, c := range selectedCodecs() {
				b.Run(c.name, func(b *testing.B) {
					warm, err := c.marshal(ds.doc)
					if err != nil {
						b.Fatal(err)
					}
					b.ReportAllocs()
					b.SetBytes(int64(len(warm)))
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						out, err := c.marshal(ds.doc)
						if err != nil {
							b.Fatal(err)
						}
						bytesSink = out
					}
					b.ReportMetric(float64(len(warm)), "output-B/op")
				})
			}
		})
	}
}

func BenchmarkMarshalDynamic(b *testing.B) {
	for _, ds := range benchmarkDatasets() {
		var dynamic map[string]any
		if err := stdjson.Unmarshal(ds.json, &dynamic); err != nil {
			b.Fatal(err)
		}
		b.Run(ds.name, func(b *testing.B) {
			for _, c := range selectedCodecs() {
				if c.concreteOnly {
					continue
				}
				b.Run(c.name, func(b *testing.B) {
					warm, err := c.marshal(dynamic)
					if err != nil {
						b.Fatal(err)
					}
					b.ReportAllocs()
					b.SetBytes(int64(len(warm)))
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						out, err := c.marshal(dynamic)
						if err != nil {
							b.Fatal(err)
						}
						bytesSink = out
					}
					b.ReportMetric(float64(len(warm)), "output-B/op")
				})
			}
		})
	}
}

func BenchmarkUnmarshalConcreteFresh(b *testing.B) {
	for _, ds := range benchmarkDatasets() {
		b.Run(ds.name, func(b *testing.B) {
			for _, c := range selectedCodecs() {
				b.Run(c.name, func(b *testing.B) {
					if err := c.unmarshal(ds.json, newConcrete(ds)); err != nil {
						b.Fatal(err)
					}
					b.ReportAllocs()
					b.SetBytes(int64(len(ds.json)))
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dst := newConcrete(ds)
						if err := c.unmarshal(ds.json, dst); err != nil {
							b.Fatal(err)
						}
						valueSink = dst
					}
				})
			}
		})
	}
}

func BenchmarkUnmarshalDynamicFresh(b *testing.B) {
	for _, ds := range benchmarkDatasets() {
		b.Run(ds.name, func(b *testing.B) {
			for _, c := range selectedCodecs() {
				if c.concreteOnly {
					continue
				}
				b.Run(c.name, func(b *testing.B) {
					warm := make(map[string]any)
					if err := c.unmarshal(ds.json, &warm); err != nil {
						b.Fatal(err)
					}
					b.ReportAllocs()
					b.SetBytes(int64(len(ds.json)))
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dst := make(map[string]any)
						if err := c.unmarshal(ds.json, &dst); err != nil {
							b.Fatal(err)
						}
						valueSink = dst
					}
				})
			}
		})
	}
}
