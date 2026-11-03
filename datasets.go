package jsonbenchmark

import (
	stdjson "encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Size tiers stay small enough that count=3–5 dual-suite runs finish in minutes,
// not half an hour. Larger payloads stay behind BENCH_SCALE=extended.
type sizeSpec struct {
	name        string
	targetBytes int
}

func sizeSpecs() []sizeSpec {
	base := []sizeSpec{
		{name: "512b", targetBytes: 512},
		{name: "8kb", targetBytes: 8 << 10},
		{name: "64kb", targetBytes: 64 << 10},
		{name: "512kb", targetBytes: 512 << 10},
		{name: "2mb", targetBytes: 2 << 20},
	}
	if strings.EqualFold(os.Getenv("BENCH_SCALE"), "extended") {
		base = append(base,
			sizeSpec{name: "8mb", targetBytes: 8 << 20},
			sizeSpec{name: "32mb", targetBytes: 32 << 20},
			sizeSpec{name: "64mb", targetBytes: 64 << 20},
		)
	}
	return base
}

type shape string

const (
	shapeFlat   shape = "flat"
	shapeNested shape = "nested"
)

type dataset struct {
	name  string // e.g. flat/64kb
	shape shape
	size  string
	doc   any
	json  []byte
}

func benchmarkDatasets() []dataset {
	filter := os.Getenv("BENCH_DATASET") // matches full name or size or shape
	shapes := []shape{shapeFlat, shapeNested}
	out := make([]dataset, 0, len(sizeSpecs())*2)
	for _, sh := range shapes {
		for _, sz := range sizeSpecs() {
			name := string(sh) + "/" + sz.name
			if filter != "" &&
				filter != name &&
				filter != sz.name &&
				filter != string(sh) {
				continue
			}
			doc := buildSizedDoc(sh, sz.name, sz.targetBytes)
			data, err := stdjson.Marshal(doc)
			if err != nil {
				panic(err)
			}
			out = append(out, dataset{name: name, shape: sh, size: sz.name, doc: doc, json: data})
		}
	}
	return out
}

func buildSizedDoc(sh shape, sizeName string, target int) any {
	switch sh {
	case shapeFlat:
		return buildFlat(sizeName, target)
	case shapeNested:
		return buildNested(sizeName, target)
	default:
		panic("unknown shape: " + string(sh))
	}
}

func buildFlat(sizeName string, target int) *FlatDoc {
	note := "ascii note with <html> & quotes /slashes/"
	n := estimateFlatItems(target)
	doc := &FlatDoc{
		ID:      "flat-" + sizeName,
		Title:   "flat benchmark " + sizeName,
		Version: 42,
		Active:  true,
		Score:   98.625,
		Tags:    []string{"benchmark", "flat", sizeName},
		Items:   make([]FlatItem, 0, n),
	}
	for i := 0; i < n; i++ {
		doc.Items = append(doc.Items, flatItem(i, note))
	}
	for {
		data, err := stdjson.Marshal(doc)
		if err != nil {
			panic(err)
		}
		if len(data) >= target || len(data) > target*11/10 {
			break
		}
		need := (target - len(data)) / 120
		if need < 1 {
			need = 1
		}
		start := len(doc.Items)
		for i := 0; i < need; i++ {
			doc.Items = append(doc.Items, flatItem(start+i, note))
		}
	}
	return doc
}

func flatItem(i int, note string) FlatItem {
	return FlatItem{
		ID:    int64(i) + 1,
		Name:  fmt.Sprintf("item-%08d", i),
		Code:  fmt.Sprintf("C%06d", i%100000),
		Price: float64(i)*1.013 + 0.99,
		Qty:   (i * 17) % 10000,
		Ok:    i%3 != 0,
		Note:  note,
	}
}

func estimateFlatItems(target int) int {
	const perItem = 120
	n := (target - 120) / perItem
	if n < 1 {
		n = 1
	}
	if target <= 600 {
		n = 2
	}
	return n
}

func buildNested(sizeName string, target int) *NestedDoc {
	depth, branching, leaves := nestedParams(target)
	doc := &NestedDoc{
		ID:    "nested-" + sizeName,
		Label: "nested benchmark " + sizeName,
		Root:  makeNestedNode("root", 0, depth, branching, leaves),
	}
	// Grow by attaching extra leaf-heavy children until near target.
	var perExtra int
	for {
		data, err := stdjson.Marshal(doc)
		if err != nil {
			panic(err)
		}
		goal := int(float64(target) * 0.85)
		if len(data) >= goal {
			break
		}
		extraLeaves := leaves + 2
		if extraLeaves < 4 {
			extraLeaves = 4
		}
		if perExtra == 0 {
			sample := makeNestedNode("root.extra.0", 1, depth, branching, extraLeaves)
			sb, err := stdjson.Marshal(sample)
			if err != nil {
				panic(err)
			}
			perExtra = len(sb)
			if perExtra < 1 {
				perExtra = 1
			}
		}
		need := (goal - len(data)) / perExtra
		if need < 1 {
			need = 1
		}
		for i := 0; i < need; i++ {
			idx := len(doc.Root.Children)
			doc.Root.Children = append(doc.Root.Children,
				makeNestedNode(fmt.Sprintf("root.extra.%d", idx), 1, depth, branching, extraLeaves))
			if idx > 10_000 {
				break
			}
		}
		if len(doc.Root.Children) > 10_000 {
			break
		}
	}
	return doc
}

func nestedParams(target int) (depth, branching, leaves int) {
	switch {
	case target <= 600:
		return 2, 1, 1
	case target <= 10<<10:
		return 3, 2, 2
	case target <= 80<<10:
		return 4, 2, 3
	case target <= 600<<10:
		return 4, 3, 4
	case target <= 3<<20:
		return 5, 3, 5
	case target <= 10<<20:
		return 5, 4, 6
	default:
		return 6, 4, 8
	}
}

func makeNestedNode(name string, level, depth, branching, leaves int) NestedNode {
	node := NestedNode{
		Name:     name,
		Kind:     "node-L" + strconv.Itoa(level),
		Weight:   float64(level) + 0.5,
		Meta:     map[string]string{"level": strconv.Itoa(level), "path": name, "note": "nested <tag> & value"},
		Leaves:   make([]NestedLeaf, leaves),
		Children: []NestedNode{},
	}
	for i := range node.Leaves {
		node.Leaves[i] = NestedLeaf{
			Key:   fmt.Sprintf("k-%d-%d", level, i),
			Value: float64(level*100+i) + 0.25,
			Flag:  i%2 == 0,
			Tags:  []string{"t1", "t2", name},
		}
	}
	if level >= depth {
		return node
	}
	node.Children = make([]NestedNode, branching)
	for i := range node.Children {
		childName := fmt.Sprintf("%s.%d", name, i)
		node.Children[i] = makeNestedNode(childName, level+1, depth, branching, leaves)
	}
	return node
}
