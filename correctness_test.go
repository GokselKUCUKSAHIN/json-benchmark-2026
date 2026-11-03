package jsonbenchmark

import (
	stdjson "encoding/json"
	"reflect"
	"testing"
)

func TestCodecRoundTrip(t *testing.T) {
	cases := []any{
		buildFlat("correctness", 2<<10),
		buildNested("correctness", 2<<10),
	}
	for _, src := range cases {
		t.Run(reflect.TypeOf(src).String(), func(t *testing.T) {
			for _, c := range allCodecs() {
				t.Run(c.name, func(t *testing.T) {
					data, err := c.marshal(src)
					if err != nil {
						t.Fatal(err)
					}
					if !stdjson.Valid(data) {
						t.Fatal("marshal produced invalid JSON")
					}
					dst := reflect.New(reflect.TypeOf(src).Elem()).Interface()
					if err := c.unmarshal(data, dst); err != nil {
						t.Fatal(err)
					}
					if !c.compat {
						return
					}
					// JSON-value equality avoids nil vs empty-slice noise.
					if !jsonValueEqual(src, dst) {
						t.Fatal("round-trip result differs from source")
					}
				})
			}
		})
	}
}

func TestCompatibleOutputSemantics(t *testing.T) {
	baseline := suiteBaselineCodec()
	if baseline.name == "" {
		t.Fatal("no baseline codec")
	}
	// encoding/json/v2 defaults (e.g. HTML escaping) differ from classic
	// encoding/json wire output; third-party "compat" modes target v1.
	if baseline.name != "encoding-json" {
		t.Skip("compat wire check only against classic encoding/json baseline")
	}
	sources := []any{
		buildFlat("compatibility", 1500),
		buildNested("compatibility", 1500),
	}
	for _, source := range sources {
		t.Run(reflect.TypeOf(source).String(), func(t *testing.T) {
			wantJSON, err := baseline.marshal(source)
			if err != nil {
				t.Fatal(err)
			}
			var want any
			if err := stdjson.Unmarshal(wantJSON, &want); err != nil {
				t.Fatal(err)
			}
			for _, c := range allCodecs() {
				if !c.compat || c.name == baseline.name {
					continue
				}
				t.Run(c.name, func(t *testing.T) {
					data, err := c.marshal(source)
					if err != nil {
						t.Fatal(err)
					}
					var got any
					if err := stdjson.Unmarshal(data, &got); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(want, got) {
						t.Fatalf("JSON semantics differ from baseline %s", baseline.name)
					}
				})
			}
		})
	}
}

func TestDatasetSizes(t *testing.T) {
	for _, ds := range benchmarkDatasets() {
		t.Logf("%s => %d bytes (target band for %s)", ds.name, len(ds.json), ds.size)
		if len(ds.json) == 0 {
			t.Fatalf("%s empty", ds.name)
		}
	}
}

func suiteBaselineCodec() codec {
	for _, c := range allCodecs() {
		if c.compat {
			return c
		}
	}
	return codec{}
}

func jsonValueEqual(a, b any) bool {
	aj, err := stdjson.Marshal(a)
	if err != nil {
		return false
	}
	bj, err := stdjson.Marshal(b)
	if err != nil {
		return false
	}
	var av, bv any
	if err := stdjson.Unmarshal(aj, &av); err != nil {
		return false
	}
	if err := stdjson.Unmarshal(bj, &bv); err != nil {
		return false
	}
	return reflect.DeepEqual(av, bv)
}
