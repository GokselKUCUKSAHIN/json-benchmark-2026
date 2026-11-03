# JSON library benchmark

Go `encoding/json` (1.23) ve `encoding/json/v2` (1.27) ile yaygın third-party
JSON kütüphanelerini hız, tahsis ve bellek açısından karşılaştırır.

## Suite'ler

| Suite | Binary (varsayılan) | Stdlib | Third-party |
|---|---|---|---|
| `go123` | `$HOME/go/go1.23.5/bin/go` | `encoding/json` | sonic, goccy, jsoniter, easyjson, segmentio |
| `go127` | `$HOME/go/go1.27.3/bin/go` | `encoding/json/v2` (+ deterministic, fast-relaxed) | aynı |

`GOEXPERIMENT=nojsonv2` yok. İki gerçek SDK binary ayrı koşar (`GO123_BIN` / `GO127_BIN` ile override).

## Codec'ler

Her third-party için compat + fast varyantı vardır (easyjson concrete-only).

Fast modlar aynı semantiği garanti etmez (map sırası, HTML escape, UTF-8,
zero-copy sahiplik).

## Dataset matrisi

Şekil × boyut:

| Boyut | Hedef |
|---|---|
| `512b` | ~0.5 KB |
| `8kb` | ~8 KB |
| `64kb` | ~64 KB |
| `512kb` | ~512 KB |
| `2mb` | ~2 MB |

Şekil:

- `flat` — sığ object + item dizisi
- `nested` — object → array → object → array iç içe ağaç

Typed (`*FlatDoc` / `*NestedDoc`) ve untyped (`map[string]any`) ayrı ölçülür.

Daha büyük payload (sunum dışı / derinlemesine):

```bash
BENCH_SCALE=extended ./scripts/run.sh   # +8mb, +32mb
```

## Operasyonlar

- `BenchmarkMarshalConcrete`
- `BenchmarkMarshalDynamic`
- `BenchmarkUnmarshalConcreteFresh`
- `BenchmarkUnmarshalDynamicFresh`

## Çalıştırma

Varsayılan hızlı suite (`count=3`, `benchtime=200ms`):

```bash
./scripts/run.sh
```

Smoke:

```bash
COUNT=1 BENCHTIME=50ms BENCH_DATASET=8kb \
  BENCH_FILTER=BenchmarkMarshalConcrete ./scripts/run.sh
```

Tek codec:

```bash
BENCH_CODEC=sonic $HOME/go/go1.27.3/bin/go test -run='^$' -bench=. -benchmem
```

easyjson generate:

```bash
go generate ./...
```

## Çıktılar

- `results/go123.txt` / `results/go127.txt` — ham bench
- `results/summary_go123.md` / `results/summary_go127.md` — ortalama ± stddev + baz oranları
- `results/benchstat.txt` — suite karşılaştırma
- `results/environment.txt` — toolchain / makine

## Metodoloji

- Setup ve warm-up timed loop dışında
- `ReportAllocs` + `SetBytes` açık
- Varsayılan `COUNT=3`, rapor ortalama ± örneklem stddev
- `GOMAXPROCS=1` scheduler gürültüsünü azaltır
- Compat codec'ler stdlib semantik eşitlik testinden geçer
