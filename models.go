package jsonbenchmark

//go:generate easyjson -all -no_std_marshalers models.go

// FlatDoc is a wide, shallow document: one object + flat item array.
type FlatDoc struct {
	ID      string     `json:"id"`
	Title   string     `json:"title"`
	Version int64      `json:"version"`
	Active  bool       `json:"active"`
	Score   float64    `json:"score"`
	Tags    []string   `json:"tags"`
	Items   []FlatItem `json:"items"`
}

type FlatItem struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
	Ok    bool    `json:"ok"`
	Note  string  `json:"note"`
}

// NestedDoc is a deep tree: object → array → object → array → object.
type NestedDoc struct {
	ID    string     `json:"id"`
	Label string     `json:"label"`
	Root  NestedNode `json:"root"`
}

type NestedNode struct {
	Name     string            `json:"name"`
	Kind     string            `json:"kind"`
	Weight   float64           `json:"weight"`
	Meta     map[string]string `json:"meta"`
	Leaves   []NestedLeaf      `json:"leaves"`
	Children []NestedNode      `json:"children"`
}

type NestedLeaf struct {
	Key   string   `json:"key"`
	Value float64  `json:"value"`
	Flag  bool     `json:"flag"`
	Tags  []string `json:"tags"`
}
