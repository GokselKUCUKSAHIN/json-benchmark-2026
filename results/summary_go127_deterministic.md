# JSON benchmark özeti

Baz: `encoding-json-v2/default-safe`. Hız >1 daha hızlı; bellek/alloc <1 daha iyi.
Değerler: ortalama ± standart sapma (count≥2).

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op | output B | hız x | bellek x | alloc x | n |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `BenchmarkMarshalConcrete/flat/2mb` | `encoding-json-v2/default-safe` | 5949575±83940 | 434.7±6.1 | 2588688±2 | 1 | 2585720 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/2mb` | `encoding-json-v2/deterministic` | 5916118±45839 | 437.1±3.4 | 2588688±1 | 1 | 2585720 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/2mb` | `encoding-json-v2/fast-relaxed` | 5853626±52528 | 441.8±3.9 | 2588687±2 | 1 | 2585720 | 1.02 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512b` | `encoding-json-v2/default-safe` | 1781±163 | 300.8±26.1 | 576 | 1 | 533 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512b` | `encoding-json-v2/deterministic` | 1684±36 | 316.7±6.9 | 576 | 1 | 533 | 1.06 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512b` | `encoding-json-v2/fast-relaxed` | 1650±29 | 323.1±5.7 | 576 | 1 | 533 | 1.08 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512kb` | `encoding-json-v2/default-safe` | 1634548±8287 | 393.5±2.0 | 647172±1 | 1 | 643247 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512kb` | `encoding-json-v2/deterministic` | 1630558±22569 | 394.5±5.5 | 647172 | 1 | 643247 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/512kb` | `encoding-json-v2/fast-relaxed` | 1580536±40505 | 407.2±10.6 | 647172±1 | 1 | 643247 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/64kb` | `encoding-json-v2/default-safe` | 205336±5605 | 388.3±10.5 | 81920 | 1 | 79701 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/64kb` | `encoding-json-v2/deterministic` | 206097±4780 | 386.9±8.9 | 81920 | 1 | 79701 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/64kb` | `encoding-json-v2/fast-relaxed` | 200198±1107 | 398.1±2.2 | 81920 | 1 | 79701 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/8kb` | `encoding-json-v2/default-safe` | 26193±810 | 371.6±11.3 | 9728 | 1 | 9727 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/8kb` | `encoding-json-v2/deterministic` | 27069±2363 | 361.1±30.7 | 9728 | 1 | 9727 | 0.97 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/flat/8kb` | `encoding-json-v2/fast-relaxed` | 25358±34 | 383.6±0.5 | 9728 | 1 | 9727 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/2mb` | `encoding-json-v2/default-safe` | 8299304±3094025 | 230.6±71.6 | 1846957±2 | 4843 | 1766836 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/2mb` | `encoding-json-v2/deterministic` | 6010976±185812 | 294.1±9.0 | 1963179±1 | 12106 | 1766836 | 1.38 | 1.06 | 2.50 | 3 |
| `BenchmarkMarshalConcrete/nested/2mb` | `encoding-json-v2/fast-relaxed` | 6188983±275076 | 285.9±12.9 | 1846956 | 4843 | 1766836 | 1.34 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/512b` | `encoding-json-v2/default-safe` | 3077±130 | 223.9±9.3 | 800 | 7 | 688 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/512b` | `encoding-json-v2/deterministic` | 3451±36 | 199.4±2.1 | 944 | 16 | 688 | 0.89 | 1.18 | 2.29 | 3 |
| `BenchmarkMarshalConcrete/nested/512b` | `encoding-json-v2/fast-relaxed` | 3915±799 | 181.3±40.9 | 800 | 7 | 688 | 0.79 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/512kb` | `encoding-json-v2/default-safe` | 2113670±806580 | 232.5±72.7 | 481828 | 1443 | 452264 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/512kb` | `encoding-json-v2/deterministic` | 1880896±263045 | 243.4±31.8 | 516441±1 | 3606 | 452264 | 1.12 | 1.07 | 2.50 | 3 |
| `BenchmarkMarshalConcrete/nested/512kb` | `encoding-json-v2/fast-relaxed` | 1659544±63731 | 272.8±10.3 | 481828±1 | 1443 | 452264 | 1.27 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/64kb` | `encoding-json-v2/default-safe` | 211839±4466 | 257.1±5.4 | 60736 | 213 | 54452 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/64kb` | `encoding-json-v2/deterministic` | 224951±1795 | 242.1±1.9 | 65825 | 531 | 54452 | 0.94 | 1.08 | 2.49 | 3 |
| `BenchmarkMarshalConcrete/nested/64kb` | `encoding-json-v2/fast-relaxed` | 211811±2424 | 257.1±2.9 | 60736 | 213 | 54452 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/8kb` | `encoding-json-v2/default-safe` | 32066±300 | 243.1±2.3 | 8896 | 45 | 7794 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalConcrete/nested/8kb` | `encoding-json-v2/deterministic` | 35062±903 | 222.4±5.7 | 9952 | 111 | 7794 | 0.91 | 1.12 | 2.47 | 3 |
| `BenchmarkMarshalConcrete/nested/8kb` | `encoding-json-v2/fast-relaxed` | 33011±1632 | 236.5±11.5 | 8896 | 45 | 7794 | 0.97 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/2mb` | `encoding-json-v2/default-safe` | 8392645±273460 | 308.3±9.9 | 2588828±1 | 11 | 2585720 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/2mb` | `encoding-json-v2/deterministic` | 11077534±61191 | 233.4±1.3 | 2588970 | 18 | 2585720 | 0.76 | 1.00 | 1.64 | 3 |
| `BenchmarkMarshalDynamic/flat/2mb` | `encoding-json-v2/fast-relaxed` | 8118878±61816 | 318.5±2.4 | 2588827±1 | 11 | 2585720 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/512b` | `encoding-json-v2/default-safe` | 2916±82 | 182.9±5.2 | 719 | 11 | 533 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/512b` | `encoding-json-v2/deterministic` | 3676±168 | 145.2±6.7 | 832 | 18 | 533 | 0.79 | 1.16 | 1.64 | 3 |
| `BenchmarkMarshalDynamic/flat/512b` | `encoding-json-v2/fast-relaxed` | 2794±35 | 190.8±2.4 | 719 | 11 | 533 | 1.04 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/512kb` | `encoding-json-v2/default-safe` | 2111970±2901 | 304.6±0.4 | 647317 | 11 | 643247 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/512kb` | `encoding-json-v2/deterministic` | 2754281±43679 | 233.6±3.7 | 647439±1 | 18 | 643247 | 0.77 | 1.00 | 1.64 | 3 |
| `BenchmarkMarshalDynamic/flat/512kb` | `encoding-json-v2/fast-relaxed` | 2164370±173854 | 298.4±23.0 | 647316±1 | 11 | 643247 | 0.98 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/64kb` | `encoding-json-v2/default-safe` | 275575±2852 | 289.2±3.0 | 82064 | 11 | 79701 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/64kb` | `encoding-json-v2/deterministic` | 365894±19663 | 218.3±12.1 | 82178 | 18 | 79701 | 0.75 | 1.00 | 1.64 | 3 |
| `BenchmarkMarshalDynamic/flat/64kb` | `encoding-json-v2/fast-relaxed` | 274091±1894 | 290.8±2.0 | 82064 | 11 | 79701 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/8kb` | `encoding-json-v2/default-safe` | 35083±744 | 277.3±5.8 | 9871 | 11 | 9727 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/flat/8kb` | `encoding-json-v2/deterministic` | 48392±4593 | 202.2±18.6 | 9984 | 18 | 9727 | 0.72 | 1.01 | 1.64 | 3 |
| `BenchmarkMarshalDynamic/flat/8kb` | `encoding-json-v2/fast-relaxed` | 35034±351 | 277.7±2.8 | 9871 | 11 | 9727 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/2mb` | `encoding-json-v2/default-safe` | 6588191±406439 | 268.9±16.3 | 1769560±2 | 7 | 1766836 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/2mb` | `encoding-json-v2/deterministic` | 8377667±107594 | 210.9±2.7 | 1769628±1 | 10 | 1766836 | 0.79 | 1.00 | 1.43 | 3 |
| `BenchmarkMarshalDynamic/nested/2mb` | `encoding-json-v2/fast-relaxed` | 6281363±42574 | 281.3±1.9 | 1769558 | 7 | 1766836 | 1.05 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/512b` | `encoding-json-v2/default-safe` | 3399±124 | 202.6±7.5 | 784 | 7 | 688 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/512b` | `encoding-json-v2/deterministic` | 4435±111 | 155.2±3.8 | 832 | 10 | 688 | 0.77 | 1.06 | 1.43 | 3 |
| `BenchmarkMarshalDynamic/nested/512b` | `encoding-json-v2/fast-relaxed` | 3142±57 | 219.0±4.0 | 784 | 7 | 688 | 1.08 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/512kb` | `encoding-json-v2/default-safe` | 1774515±26085 | 254.9±3.8 | 458834±1 | 7 | 452264 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/512kb` | `encoding-json-v2/deterministic` | 2120007±91116 | 213.6±9.2 | 458891±2 | 10 | 452264 | 0.84 | 1.00 | 1.43 | 3 |
| `BenchmarkMarshalDynamic/nested/512kb` | `encoding-json-v2/fast-relaxed` | 1767811±32072 | 255.9±4.6 | 458835±1 | 7 | 452264 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/64kb` | `encoding-json-v2/default-safe` | 203534±501 | 267.5±0.7 | 57424 | 7 | 54452 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/64kb` | `encoding-json-v2/deterministic` | 285718±2733 | 190.6±1.8 | 57473 | 10 | 54452 | 0.71 | 1.00 | 1.43 | 3 |
| `BenchmarkMarshalDynamic/nested/64kb` | `encoding-json-v2/fast-relaxed` | 207782±2271 | 262.1±2.8 | 57424 | 7 | 54452 | 0.98 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/8kb` | `encoding-json-v2/default-safe` | 32476±1049 | 240.2±7.7 | 8272 | 7 | 7794 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkMarshalDynamic/nested/8kb` | `encoding-json-v2/deterministic` | 43367±408 | 179.7±1.7 | 8320 | 10 | 7794 | 0.75 | 1.01 | 1.43 | 3 |
| `BenchmarkMarshalDynamic/nested/8kb` | `encoding-json-v2/fast-relaxed` | 31585±1029 | 246.9±7.9 | 8272 | 7 | 7794 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/2mb` | `encoding-json-v2/default-safe` | 13022521±127308 | 218.7±2.2 | 7705311±1 | 52605 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/2mb` | `encoding-json-v2/deterministic` | 12913407±86040 | 220.5±1.5 | 7705311±2 | 52605 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/2mb` | `encoding-json-v2/fast-relaxed` | 12678474±43865 | 224.6±0.8 | 7705309±5 | 52605 | 0 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512b` | `encoding-json-v2/default-safe` | 3044±16 | 189.9±1.0 | 1048 | 16 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512b` | `encoding-json-v2/deterministic` | 3041±56 | 190.1±3.5 | 1048 | 16 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512b` | `encoding-json-v2/fast-relaxed` | 3006±10 | 192.3±0.7 | 1048 | 16 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512kb` | `encoding-json-v2/default-safe` | 3212470±68163 | 220.7±4.7 | 1836181±1 | 13168 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512kb` | `encoding-json-v2/deterministic` | 3189015±20227 | 222.3±1.4 | 1836181±1 | 13168 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/512kb` | `encoding-json-v2/fast-relaxed` | 3152414±34753 | 224.8±2.5 | 1836182±2 | 13168 | 0 | 1.02 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/64kb` | `encoding-json-v2/default-safe` | 384121±2238 | 228.8±1.3 | 146097±1 | 1650 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/64kb` | `encoding-json-v2/deterministic` | 378964±2753 | 231.9±1.7 | 146097 | 1650 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/64kb` | `encoding-json-v2/fast-relaxed` | 376123±2422 | 233.6±1.5 | 146097 | 1650 | 0 | 1.02 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/8kb` | `encoding-json-v2/default-safe` | 48051±104 | 223.3±0.5 | 18040 | 214 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/8kb` | `encoding-json-v2/deterministic` | 48367±101 | 221.9±0.5 | 18040 | 214 | 0 | 0.99 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/flat/8kb` | `encoding-json-v2/fast-relaxed` | 48042±168 | 223.4±0.8 | 18040 | 214 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/2mb` | `encoding-json-v2/default-safe` | 13263823±25158 | 135.9±0.3 | 5593052±1 | 93581 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/2mb` | `encoding-json-v2/deterministic` | 13326248±115305 | 135.3±1.2 | 5593051±1 | 93581 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/2mb` | `encoding-json-v2/fast-relaxed` | 13381919±27389 | 134.7±0.3 | 5593054±3 | 93581 | 0 | 0.99 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512b` | `encoding-json-v2/default-safe` | 5425±131 | 135.2±3.3 | 2120 | 33 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512b` | `encoding-json-v2/deterministic` | 5417±29 | 135.3±0.7 | 2120 | 33 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512b` | `encoding-json-v2/fast-relaxed` | 5379±28 | 136.3±0.7 | 2120 | 33 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512kb` | `encoding-json-v2/default-safe` | 3471619±59070 | 133.4±2.3 | 1518522±1 | 24544 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512kb` | `encoding-json-v2/deterministic` | 3495760±40133 | 132.5±1.5 | 1518522±1 | 24544 | 0 | 0.99 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/512kb` | `encoding-json-v2/fast-relaxed` | 3409358±28776 | 135.8±1.1 | 1518522±1 | 24544 | 0 | 1.02 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/64kb` | `encoding-json-v2/default-safe` | 429172±360 | 130.6±0.1 | 193825 | 3088 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/64kb` | `encoding-json-v2/deterministic` | 430202±3024 | 130.3±0.9 | 193825 | 3088 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/64kb` | `encoding-json-v2/fast-relaxed` | 417592±2599 | 134.2±0.8 | 193825 | 3088 | 0 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/8kb` | `encoding-json-v2/default-safe` | 62787±1460 | 129.4±3.0 | 25248 | 452 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/8kb` | `encoding-json-v2/deterministic` | 63570±594 | 127.8±1.2 | 25248 | 452 | 0 | 0.99 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalConcreteFresh/nested/8kb` | `encoding-json-v2/fast-relaxed` | 62480±697 | 130.0±1.5 | 25248 | 452 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/2mb` | `encoding-json-v2/default-safe` | 19732661±202630 | 144.3±1.5 | 10564811±1 | 314747 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/2mb` | `encoding-json-v2/deterministic` | 19230626±254570 | 148.1±2.0 | 10564814±2 | 314747 | 0 | 1.03 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/2mb` | `encoding-json-v2/fast-relaxed` | 34721315±386796 | 82.0±0.9 | 12242607±2 | 420260 | 0 | 0.57 | 1.16 | 1.34 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512b` | `encoding-json-v2/default-safe` | 4749±58 | 121.7±1.5 | 2336 | 73 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512b` | `encoding-json-v2/deterministic` | 4910±142 | 117.8±3.4 | 2336 | 73 | 0 | 0.97 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512b` | `encoding-json-v2/fast-relaxed` | 8133±13 | 71.1±0.1 | 2792 | 103 | 0 | 0.58 | 1.20 | 1.41 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512kb` | `encoding-json-v2/default-safe` | 4674461±50734 | 151.6±1.6 | 2583882±1 | 78706 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512kb` | `encoding-json-v2/deterministic` | 4697659±108248 | 150.9±3.5 | 2583881±1 | 78706 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/512kb` | `encoding-json-v2/fast-relaxed` | 8399852±63998 | 84.4±0.6 | 3003385±1 | 105155 | 0 | 0.56 | 1.16 | 1.34 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/64kb` | `encoding-json-v2/default-safe` | 563800±4910 | 155.9±1.4 | 311554 | 9839 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/64kb` | `encoding-json-v2/deterministic` | 575214±3857 | 152.8±1.0 | 311554 | 9839 | 0 | 0.98 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/64kb` | `encoding-json-v2/fast-relaxed` | 1025654±16233 | 85.7±1.4 | 364050±1 | 13151 | 0 | 0.55 | 1.17 | 1.34 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/8kb` | `encoding-json-v2/default-safe` | 71919±1048 | 149.2±2.2 | 38736 | 1233 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/8kb` | `encoding-json-v2/deterministic` | 70928±982 | 151.3±2.1 | 38736 | 1233 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/flat/8kb` | `encoding-json-v2/fast-relaxed` | 128929±1041 | 83.2±0.7 | 45336 | 1654 | 0 | 0.56 | 1.17 | 1.34 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/2mb` | `encoding-json-v2/default-safe` | 18148867±102614 | 99.4±0.6 | 12166264±6 | 329008 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/2mb` | `encoding-json-v2/deterministic` | 18182520±504013 | 99.2±2.8 | 12166263±5 | 329008 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/2mb` | `encoding-json-v2/fast-relaxed` | 34771475±454325 | 51.9±0.7 | 14528505±3 | 455457 | 0 | 0.52 | 1.19 | 1.38 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512b` | `encoding-json-v2/default-safe` | 6661±19 | 110.0±0.3 | 4928 | 126 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512b` | `encoding-json-v2/deterministic` | 6740±50 | 108.8±0.8 | 4928 | 126 | 0 | 0.99 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512b` | `encoding-json-v2/fast-relaxed` | 12729±222 | 57.6±1.0 | 5840 | 166 | 0 | 0.52 | 1.19 | 1.32 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512kb` | `encoding-json-v2/default-safe` | 4374739±40017 | 105.9±1.0 | 3181070±2 | 85808 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512kb` | `encoding-json-v2/deterministic` | 4376879±15176 | 105.8±0.4 | 3181069±2 | 85808 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/512kb` | `encoding-json-v2/fast-relaxed` | 8614380±76109 | 53.8±0.5 | 3794912±1 | 118248 | 0 | 0.51 | 1.19 | 1.38 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/64kb` | `encoding-json-v2/default-safe` | 520932±8676 | 107.6±1.8 | 389594±1 | 10506 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/64kb` | `encoding-json-v2/deterministic` | 514047±12358 | 109.1±2.6 | 389595±1 | 10506 | 0 | 1.01 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/64kb` | `encoding-json-v2/fast-relaxed` | 1000329±28026 | 56.1±1.6 | 463404±1 | 14324 | 0 | 0.52 | 1.19 | 1.36 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/8kb` | `encoding-json-v2/default-safe` | 75233±959 | 108.0±1.4 | 55912 | 1528 | 0 | 1.00 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/8kb` | `encoding-json-v2/deterministic` | 76692±1160 | 105.9±1.6 | 55912 | 1528 | 0 | 0.98 | 1.00 | 1.00 | 3 |
| `BenchmarkUnmarshalDynamicFresh/nested/8kb` | `encoding-json-v2/fast-relaxed` | 148928±4784 | 54.6±1.8 | 66536 | 2050 | 0 | 0.51 | 1.19 | 1.34 | 3 |

## Deterministic(true) etkisi (vs `default-safe`)

Odak: `encoding-json-v2/deterministic` = `jsonv2.Marshal(..., jsonv2.Deterministic(true))`.
Deterministic yalnızca marshal; unmarshal path aynı.

### MarshalConcrete / flat (map yok)
- Hız ~1.00x, bellek/alloc aynı → maliyet yok.

### MarshalConcrete / nested (`Meta map[string]string`)
- Hız ~0.89–0.94x (küçük/orta); büyük boyutlarda yüksek varyans.
- Bellek ~1.06–1.18x, **alloc ~2.5x** (key sort scratch).

### MarshalDynamic (`map[string]any`) — asıl maliyet
- Flat: hız **~0.72–0.79x**, alloc **1.64x**.
- Nested: hız **~0.71–0.84x**, alloc **1.43x**.
- Bellek çoğu boyutta ~1.00x (sort scratch küçük kalıyor).

### Unmarshal (concrete + dynamic)
- Hız/bellek/alloc **~1.00x** — beklenen (opsiyon marshal-only).

### Ne zaman aç?
- **Aç:** cache key, snapshot/diff, test golden, config hash — wire byte stabilitesi gerekir.
- **Kapat (default):** RPC/throughput path; object member sırası zaten anlamsız.
