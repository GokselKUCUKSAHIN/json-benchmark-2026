# JSON: ne encode/decode, hangi kütüphane

Kaynak: dual suite, Apple M1 Pro, `GOMAXPROCS=1`, count=3.  
Baz (1.27): `encoding/json/v2` default-safe.  
Baz (1.23): `encoding/json`.  
Ölçümler: typed struct / `map[string]any`, flat / nested, 512b→2mb.

**Tek kural:** hız kazanıp alloc patlatmak çoğu serviste net kayıp.  
Önce alloc/B/op’a bak, sonra ns/op.

---

## Hızlı karar

| Ne yapıyorsun | Kullan | Kullanma |
|---|---|---|
| Typed marshal | **go-json** (compat veya fast). Alloc ≈ stdlib (1a flat). | sonic (yavaş + büyük payload’da mem ×3–5). segmentio/fast (mem ×4–6). |
| Typed unmarshal | **go-json**. Hız ×2.5–3 vs v2, bellek ≤ v2. | jsoniter (alloc ×1.8–2 vs v2). |
| `map[string]any` marshal | **json/v2** veya **go-json**. v2 burada third-party’ye en yakın. | **jsoniter** (alloc ×10³–10⁴). **sonic** dynamic marshal (alloc ×10³ + mem ×4–6, hız da kötü). |
| `map[string]any` unmarshal | **sonic/fast** (hız ×1.6–2.2, alloc ×0.14–0.17 vs v2) **veya json/v2** (bağımlılık istemiyorsan). | jsoniter (daha yavaş + daha çok alloc). segmentio (hız kötü). |
| Sade bağımlılık / genel API | **encoding/json/v2** | “stdlib her şeyi çözer” iddiası — hot path typed’de go-json hâlâ ×2.5–3.5 önde. |

Go 1.23 `encoding/json` → Go 1.27 `json/v2`:  
unmarshal typed ~×2, dynamic marshal ~×3, dynamic unmarshal ~×1.4.  
Typed marshal’da v2 biraz **geride** (~0.75×).

---

## Alloc patlaması — kırmızı liste

Bunlar sunumda ayrıca göster. “Hızlı” görünüp heap’i yiyenler:

### 1) jsoniter + `map[string]any` marshal — en kötü

| Senaryo (1.27, 2mb) | vs v2 alloc | vs v2 mem | hız |
|---|---:|---:|---:|
| flat dynamic marshal `jsoniter/compat` | **×25400** | ×7.1 | 0.31× (yavaş) |
| nested dynamic marshal `jsoniter/compat` | **×36800** | ×8.4 | 0.30× |
| nested dynamic marshal `jsoniter/fast` | **×12000** | ×3.7 | 0.62× |

Aynı lib typed marshal’da sakin (1–2 alloc).  
**Teşhis:** dynamic encode yolu reflection/alloc cehennemi.  
**Kural:** `any`/`map[string]any` marshal’da jsoniter yasak.

### 2) sonic + `map[string]any` marshal — hızlı değil, alloc canavar

| Senaryo (1.27, 2mb) | vs v2 alloc | vs v2 mem | hız |
|---|---:|---:|---:|
| flat dynamic marshal sonic | **×1590** | ×4–6 | 0.30–0.37× |
| nested dynamic marshal sonic | **×3010** | ×5–6 | 0.28–0.32× |

sonic **unmarshal dynamic**’te kral; **marshal dynamic**’te felaket.  
Aynı isim, zıt profil — karıştırma.

### 3) segmentio/fast — mem şişmesi (alloc sayısı az olsa da)

| Senaryo | vs v2 mem | not |
|---|---:|---|
| flat typed marshal 2mb | **×5.2** | 37 alloc ama ~13 MB/op |
| nested typed marshal 2mb | **×5.7** | buffer/zero-copy yan etkisi |
| nested typed unmarshal | **×3.0–3.8** | |

compat genelde temiz (1 alloc, mem ≈ baz).  
**Kural:** segmentio kullanacaksan **compat**; fast’i profillemeden alma.

### 4) sonic typed marshal büyük payload — yavaş + mem

flat/nested typed marshal 2mb: mem **×3.4–5.1**, hız **0.6–0.8×** (v2’den yavaş).  
arm64/M1’de SIMD avantajı yok gibi. amd64 ayrı ölç.

### 5) easyjson typed marshal — alloc sayısı yüksek, byte düşük

flat typed marshal 2mb: **95–103 alloc** (stdlib/v2 = 1) ama B/op ≈ aynı.  
Küçük/orta payload’da da 6–14 alloc.  
Hız iyi (~1.5–1.9× v2); GC pressure alloc *sayısından* gelir — hot tight loop’ta dikkat.  
Concrete-only (codegen şart).

### 6) jsoniter typed unmarshal — orta seviye alloc şişmesi

vs v2: alloc ~×1.8, mem ~×1.1–1.3.  
Hız hâlâ > v2 ama go-json/sonic’den kötü. “compat = ucuz” değil.

---

## Temiz profiller (alloc dostu)

| Codec | Typed marshal | Dynamic marshal | Typed unmarshal | Dynamic unmarshal |
|---|---|---|---|---|
| **json/v2** | 1a, mem baz; hız orta/yavaş | **çok iyi** (7a nested!) | iyi (1.23’e göre ×2) | iyi baz |
| **go-json** | **en iyi hız + 1a** | **en iyi hız + 1a** | **en iyi hız**, alloc ≈ v2 | hızlı, alloc ≈ v2 |
| **sonic/fast** | kaçın (yavaş/mem) | **kaçın (alloc patlar)** | iyi, az alloc | **en iyi hız, az alloc** |
| **segmentio/compat** | iyi, temiz | temiz ama v2/go-json gerisinde | OK | zayıf hız |
| **segmentio/fast** | mem patlar | mem patlar | mem yüksek | zayıf |
| **jsoniter** | OK | **yasak** | alloc yüksek | alloc yüksek + yavaş |
| **easyjson** | hızlı, çok alloc adedi | yok (concrete-only) | OK | yok |

---

## Senaryo kartları (1.27, tipik 64kb–2mb)

### A) Typed struct in/out (API DTO)

1. **go-json/compat** — hız + alloc dengesi  
2. json/v2 — bağımlılık istemiyorsan (unmarshal’da 1.23’ten belirgin iyi; marshal’da go-json geride)  
3. easyjson — codegen OK ve alloc adedi sorun değilse  

Kaçın: sonic marshal, segmentio/fast, jsoniter (unmarshal alloc).

### B) `map[string]any` / esnek JSON

1. **Marshal:** json/v2 veya go-json (ikisi de düşük alloc)  
2. **Unmarshal:** sonic/fast (max hız) veya json/v2 (temiz)  

Kaçın: jsoniter marshal, sonic marshal, segmentio.

### C) Nested ağaç / derin JSON

Sıralama aynı; absolute süre ve alloc artar.  
v2 nested dynamic marshal’da zaten çok temiz (7 alloc) — third-party farkı burda **küçülür** (~1.3× go-json).  
segmentio/sonic mem şişmesi nested’de daha görünür.

### D) “Stdlib yeter mi?”

- Evet: genel trafik, typed+dynamic karışık, bağımlılık maliyeti yüksek, hot path yok.  
- Hayır: typed hot path (go-json ×2.5–3.5), veya dynamic decode latency kritik (sonic).  
- v2, 1.23 stdlib’in “dynamic marshal cehennemi”ni büyük ölçüde bitirdi; typed marshal tahtını vermedi.

---

## Compat vs fast (alloc açısından)

| Lib | fast ne yapar | alloc etkisi |
|---|---|---|
| go-json | HTML escape off, unordered map | genelde aynı 1a; güvenli seçim |
| sonic | ConfigFastest | dynamic marshal’da zaten patlak; fast kurtarmıyor |
| segmentio | Append/ZeroCopy | **mem ×3–6** — asıl risk |
| jsoniter | ConfigFastest | dynamic marshal hâlâ facia; typed’de ufak kazanç |
| json/v2 fast-relaxed | duplicate/utf8/HTML gevşek | typed ≈ default; dynamic unmarshal’da bazen **daha yavaş + daha alloc** — “fast” sanma |

---

## Tek bakışlık özet

```
typed marshal          → go-json          (alloc temiz)
typed unmarshal        → go-json          (alloc ≈ v2, hız ×2.5+)
dynamic marshal        → json/v2 | go-json
dynamic unmarshal      → sonic/fast | json/v2

ASLA:
  jsoniter  + map marshal
  sonic     + map marshal
  segmentio/fast  (mem)
  sonic     + typed marshal (arm64, büyük payload)
```

Detay tablolar: `results/summary_go123.md`, `results/summary_go127.md`.
