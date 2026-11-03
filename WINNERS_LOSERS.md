# JSON benchmark: senaryo kazananları ve kaybedenleri

Her satır, aynı suite + veri tipi + işlem + şekil + boyut içindeki codec karşılaştırmasıdır.

## Seçim yöntemi

- Kazanan/kaybeden, CPU (`ns/op`), bellek (`B/op`) ve tahsis (`allocs/op`) maliyetlerinin eşit ağırlıklı geometrik ortalamasıyla seçildi.
- Her metrik senaryodaki en iyi değere normalize edildi: `skor = ∛((ns/min_ns) × (B/min_B) × (alloc/min_alloc))`.
- Düşük skor daha iyi. `MB/s`, `ns/op` ile aynı performansı ters yönde gösterdiği için skora tekrar eklenmedi; tabloda bilgi amaçlı verildi.
- `fast` modlar compat modlarla aynı JSON semantiğini garanti etmez. Sonuç “ölçülen maliyet kazananı”; davranış uyumluluğu kararı ayrıca verilmeli.
- easyjson yalnız typed senaryolara katılır.

Metrik hücresi sırası: `ns/op · B/op · allocs/op · MB/s`.

## Tablo 1 — Go 1.23

| Veri tipi | İşlem | Şekil | Boyut | Kazanan | Kazanan metrikleri | Kaybeden | Kaybeden metrikleri |
|---|---|---|---:|---|---|---|---|
| typed | marshal | flat | 512b | `go-json/fast` | 708 ns · 576 B · 1 alloc · 753.2 MB/s | `segmentio/fast` | 1_275 ns · 1_912 B · 8 alloc · 418.0 MB/s |
| typed | marshal | flat | 8kb | `go-json/fast` | 10_642 ns · 9_728 B · 1 alloc · 914.3 MB/s | `segmentio/fast` | 25_087 ns · 46_584 B · 16 alloc · 387.8 MB/s |
| typed | marshal | flat | 64kb | `go-json/fast` | 81_011 ns · 81_920 B · 1 alloc · 983.9 MB/s | `segmentio/fast` | 179_774 ns · 383_736 B · 23 alloc · 443.4 MB/s |
| typed | marshal | flat | 512kb | `go-json/fast` | 622_854 ns · 647_172 B · 1 alloc · 1_032.8 MB/s | `segmentio/fast` | 1_343_830 ns · 3_241_392 B · 30 alloc · 478.7 MB/s |
| typed | marshal | flat | 2mb | `go-json/fast` | 2_449_067 ns · 2_588_686 B · 1 alloc · 1_055.8 MB/s | `sonic/compat` | 11_245_096 ns · 13_249_624 B · 26 alloc · 253.3 MB/s |
| typed | marshal | nested | 512b | `go-json/compat` | 1_692 ns · 768 B · 1 alloc · 433.2 MB/s | `jsoniter/compat` | 3_844 ns · 2_120 B · 29 alloc · 190.7 MB/s |
| typed | marshal | nested | 8kb | `go-json/compat` | 17_314 ns · 8_192 B · 1 alloc · 469.9 MB/s | `jsoniter/compat` | 35_110 ns · 18_056 B · 200 alloc · 231.5 MB/s |
| typed | marshal | nested | 64kb | `go-json/compat` | 107_427 ns · 57_344 B · 1 alloc · 521.7 MB/s | `jsoniter/compat` | 215_209 ns · 104_843 B · 956 alloc · 260.4 MB/s |
| typed | marshal | nested | 512kb | `go-json/compat` | 837_110 ns · 466_948 B · 1 alloc · 553.2 MB/s | `jsoniter/compat` | 1_610_263 ns · 789_980 B · 6_491 alloc · 287.6 MB/s |
| typed | marshal | nested | 2mb | `go-json/compat` | 3_101_579 ns · 1_810_442 B · 1 alloc · 581.4 MB/s | `jsoniter/compat` | 5_884_836 ns · 2_895_106 B · 21_791 alloc · 306.4 MB/s |
| typed | unmarshal | flat | 512b | `go-json/compat` | 1_275 ns · 776 B · 17 alloc · 453.5 MB/s | `encoding-json` | 7_630 ns · 1_552 B · 31 alloc · 75.8 MB/s |
| typed | unmarshal | flat | 8kb | `sonic/fast` | 28_552 ns · 19_776 B · 75 alloc · 375.9 MB/s | `jsoniter/compat` | 52_463 ns · 36_728 B · 568 alloc · 204.6 MB/s |
| typed | unmarshal | flat | 64kb | `sonic/fast` | 221_411 ns · 165_729 B · 553 alloc · 396.9 MB/s | `jsoniter/compat` | 411_927 ns · 288_322 B · 4_549 alloc · 213.3 MB/s |
| typed | unmarshal | flat | 512kb | `segmentio/fast` | 2_987_680 ns · 1_111_568 B · 4_380 alloc · 237.2 MB/s | `jsoniter/fast` | 3_406_559 ns · 2_714_505 B · 36_287 alloc · 208.1 MB/s |
| typed | unmarshal | flat | 2mb | `segmentio/fast` | 11_949_223 ns · 4_408_016 B · 17_489 alloc · 238.3 MB/s | `jsoniter/compat` | 13_811_803 ns · 11_160_687 B · 145_051 alloc · 206.2 MB/s |
| typed | unmarshal | nested | 512b | `sonic/fast` | 3_610 ns · 2_792 B · 29 alloc · 203.1 MB/s | `segmentio/compat` | 9_440 ns · 19_536 B · 53 alloc · 77.7 MB/s |
| typed | unmarshal | nested | 8kb | `sonic/fast` | 35_796 ns · 26_544 B · 250 alloc · 227.0 MB/s | `segmentio/compat` | 78_768 ns · 144_456 B · 545 alloc · 103.2 MB/s |
| typed | unmarshal | nested | 64kb | `sonic/fast` | 227_560 ns · 170_337 B · 1_570 alloc · 246.3 MB/s | `segmentio/compat` | 468_037 ns · 735_928 B · 3_559 alloc · 119.7 MB/s |
| typed | unmarshal | nested | 512kb | `sonic/fast` | 1_856_056 ns · 1_351_215 B · 12_247 alloc · 249.5 MB/s | `encoding-json` | 7_431_114 ns · 1_705_104 B · 42_652 alloc · 62.3 MB/s |
| typed | unmarshal | nested | 2mb | `go-json/fast` | 6_293_338 ns · 4_550_860 B · 108_533 alloc · 286.6 MB/s | `encoding-json` | 28_367_354 ns · 6_296_978 B · 161_151 alloc · 63.6 MB/s |
| untyped | marshal | flat | 512b | `go-json/fast` | 1_603 ns · 576 B · 1 alloc · 332.6 MB/s | `jsoniter/compat` | 7_586 ns · 3_809 B · 61 alloc · 76.2 MB/s |
| untyped | marshal | flat | 8kb | `go-json/fast` | 24_872 ns · 9_728 B · 1 alloc · 391.1 MB/s | `jsoniter/compat` | 128_264 ns · 64_525 B · 1_021 alloc · 83.7 MB/s |
| untyped | marshal | flat | 64kb | `go-json/fast` | 195_816 ns · 81_920 B · 1 alloc · 407.0 MB/s | `jsoniter/compat` | 994_700 ns · 520_755 B · 8_191 alloc · 88.3 MB/s |
| untyped | marshal | flat | 512kb | `go-json/fast` | 1_557_942 ns · 647_172 B · 1 alloc · 412.9 MB/s | `jsoniter/compat` | 8_542_465 ns · 4_158_457 B · 65_537 alloc · 83.0 MB/s |
| untyped | marshal | flat | 2mb | `go-json/fast` | 6_256_151 ns · 2_588_683 B · 1 alloc · 413.3 MB/s | `jsoniter/compat` | 33_961_728 ns · 16_752_377 B · 262_145 alloc · 83.9 MB/s |
| untyped | marshal | nested | 512b | `go-json/fast` | 2_745 ns · 704 B · 1 alloc · 250.6 MB/s | `jsoniter/compat` | 12_133 ns · 6_357 B · 116 alloc · 60.4 MB/s |
| untyped | marshal | nested | 8kb | `go-json/fast` | 30_896 ns · 8_192 B · 1 alloc · 252.3 MB/s | `jsoniter/compat` | 125_737 ns · 62_774 B · 1_177 alloc · 64.6 MB/s |
| untyped | marshal | nested | 64kb | `go-json/fast` | 206_366 ns · 57_344 B · 1 alloc · 263.9 MB/s | `jsoniter/compat` | 811_057 ns · 404_788 B · 7_703 alloc · 69.2 MB/s |
| untyped | marshal | nested | 512kb | `go-json/fast` | 1_729_896 ns · 458_754 B · 1 alloc · 261.4 MB/s | `jsoniter/compat` | 6_798_090 ns · 3_242_391 B · 62_240 alloc · 68.1 MB/s |
| untyped | marshal | nested | 2mb | `go-json/fast` | 6_429_502 ns · 1_769_478 B · 1 alloc · 274.9 MB/s | `jsoniter/compat` | 25_280_656 ns · 12_611_661 B · 236_526 alloc · 71.4 MB/s |
| untyped | unmarshal | flat | 512b | `sonic/fast` | 2_874 ns · 2_904 B · 21 alloc · 201.2 MB/s | `jsoniter/compat` | 6_591 ns · 3_376 B · 136 alloc · 87.7 MB/s |
| untyped | unmarshal | flat | 8kb | `sonic/fast` | 39_670 ns · 43_704 B · 213 alloc · 270.6 MB/s | `jsoniter/compat` | 110_809 ns · 56_544 B · 2_142 alloc · 96.9 MB/s |
| untyped | unmarshal | flat | 64kb | `sonic/fast` | 313_242 ns · 350_393 B · 1_647 alloc · 280.5 MB/s | `jsoniter/compat` | 884_111 ns · 455_459 B · 17_117 alloc · 99.4 MB/s |
| untyped | unmarshal | flat | 512kb | `sonic/fast` | 2_660_833 ns · 3_840_713 B · 13_117 alloc · 266.4 MB/s | `jsoniter/fast` | 7_348_854 ns · 3_732_006 B · 136_782 alloc · 96.5 MB/s |
| untyped | unmarshal | flat | 2mb | `sonic/fast` | 12_042_346 ns · 29_288_314 B · 52_439 alloc · 236.5 MB/s | `jsoniter/fast` | 30_626_439 ns · 15_156_695 B · 547_005 alloc · 93.0 MB/s |
| untyped | unmarshal | nested | 512b | `sonic/fast` | 4_800 ns · 5_584 B · 33 alloc · 152.8 MB/s | `jsoniter/fast` | 10_870 ns · 6_128 B · 215 alloc · 67.4 MB/s |
| untyped | unmarshal | nested | 8kb | `sonic/fast` | 46_837 ns · 57_336 B · 238 alloc · 173.5 MB/s | `jsoniter/compat` | 127_940 ns · 69_184 B · 2_489 alloc · 63.5 MB/s |
| untyped | unmarshal | nested | 64kb | `sonic/fast` | 296_241 ns · 384_154 B · 1_478 alloc · 189.2 MB/s | `segmentio/compat` | 1_533_541 ns · 440_880 B · 11_402 alloc · 36.5 MB/s |
| untyped | unmarshal | nested | 512kb | `sonic/fast` | 2_392_472 ns · 3_120_045 B · 11_785 alloc · 193.6 MB/s | `segmentio/compat` | 13_148_378 ns · 3_598_225 B · 93_711 alloc · 35.2 MB/s |
| untyped | unmarshal | nested | 2mb | `sonic/fast` | 12_092_019 ns · 23_568_667 B · 44_556 alloc · 149.4 MB/s | `segmentio/compat` | 58_535_378 ns · 13_745_693 B · 359_741 alloc · 30.8 MB/s |


## Tablo 2 — Go 1.27

| Veri tipi | İşlem | Şekil | Boyut | Kazanan | Kazanan metrikleri | Kaybeden | Kaybeden metrikleri |
|---|---|---|---:|---|---|---|---|
| typed | marshal | flat | 512b | `go-json/fast` | 586 ns · 576 B · 1 alloc · 909.8 MB/s | `segmentio/fast` | 1_181 ns · 1_912 B · 8 alloc · 451.7 MB/s |
| typed | marshal | flat | 8kb | `go-json/fast` | 8_824 ns · 9_728 B · 1 alloc · 1_102.5 MB/s | `segmentio/fast` | 21_136 ns · 46_584 B · 16 alloc · 460.6 MB/s |
| typed | marshal | flat | 64kb | `go-json/fast` | 66_399 ns · 81_920 B · 1 alloc · 1_200.4 MB/s | `segmentio/fast` | 158_612 ns · 383_736 B · 23 alloc · 502.5 MB/s |
| typed | marshal | flat | 512kb | `go-json/fast` | 513_758 ns · 647_172 B · 1 alloc · 1_252.0 MB/s | `segmentio/fast` | 1_215_803 ns · 3_241_392 B · 30 alloc · 529.1 MB/s |
| typed | marshal | flat | 2mb | `go-json/fast` | 2_034_675 ns · 2_588_686 B · 1 alloc · 1_270.8 MB/s | `sonic/compat` | 11_314_391 ns · 13_249_624 B · 26 alloc · 251.7 MB/s |
| typed | marshal | nested | 512b | `go-json/compat` | 1_495 ns · 768 B · 1 alloc · 490.3 MB/s | `jsoniter/compat` | 4_137 ns · 2_408 B · 32 alloc · 177.2 MB/s |
| typed | marshal | nested | 8kb | `go-json/compat` | 15_407 ns · 8_192 B · 1 alloc · 527.3 MB/s | `jsoniter/compat` | 39_277 ns · 20_168 B · 222 alloc · 206.9 MB/s |
| typed | marshal | nested | 64kb | `go-json/compat` | 94_293 ns · 57_344 B · 1 alloc · 594.4 MB/s | `jsoniter/compat` | 234_530 ns · 115_019 B · 1_062 alloc · 239.0 MB/s |
| typed | marshal | nested | 512kb | `go-json/compat` | 737_175 ns · 466_947 B · 1 alloc · 628.3 MB/s | `jsoniter/compat` | 1_866_751 ns · 859_199 B · 7_212 alloc · 248.1 MB/s |
| typed | marshal | nested | 2mb | `go-json/compat` | 2_693_453 ns · 1_810_442 B · 1 alloc · 669.5 MB/s | `jsoniter/compat` | 6_579_044 ns · 3_225_580 B · 24_212 alloc · 274.1 MB/s |
| typed | unmarshal | flat | 512b | `go-json/compat` | 1_179 ns · 776 B · 17 alloc · 490.4 MB/s | `sonic/compat` | 2_358 ns · 1_464 B · 22 alloc · 245.1 MB/s |
| typed | unmarshal | flat | 8kb | `sonic/fast` | 27_293 ns · 19_776 B · 75 alloc · 393.3 MB/s | `jsoniter/fast` | 48_292 ns · 32_968 B · 367 alloc · 222.6 MB/s |
| typed | unmarshal | flat | 64kb | `sonic/fast` | 214_354 ns · 165_729 B · 553 alloc · 410.0 MB/s | `jsoniter/compat` | 364_534 ns · 257_801 B · 2_914 alloc · 241.1 MB/s |
| typed | unmarshal | flat | 512kb | `segmentio/fast` | 2_964_016 ns · 1_111_568 B · 4_380 alloc · 239.1 MB/s | `jsoniter/compat` | 3_078_305 ns · 2_469_882 B · 23_183 alloc · 230.3 MB/s |
| typed | unmarshal | flat | 2mb | `segmentio/fast` | 11_854_221 ns · 4_408_016 B · 17_489 alloc · 240.3 MB/s | `jsoniter/fast` | 12_671_812 ns · 10_182_071 B · 92_625 alloc · 224.8 MB/s |
| typed | unmarshal | nested | 512b | `sonic/fast` | 3_439 ns · 2_792 B · 29 alloc · 213.2 MB/s | `segmentio/compat` | 8_715 ns · 19_584 B · 56 alloc · 84.1 MB/s |
| typed | unmarshal | nested | 8kb | `sonic/fast` | 34_310 ns · 26_544 B · 250 alloc · 236.8 MB/s | `segmentio/compat` | 77_557 ns · 144_808 B · 567 alloc · 104.8 MB/s |
| typed | unmarshal | nested | 64kb | `sonic/fast` | 216_098 ns · 170_337 B · 1_570 alloc · 259.4 MB/s | `segmentio/compat` | 451_716 ns · 737_624 B · 3_665 alloc · 124.1 MB/s |
| typed | unmarshal | nested | 512kb | `sonic/fast` | 1_747_287 ns · 1_351_217 B · 12_247 alloc · 265.1 MB/s | `segmentio/compat` | 3_459_292 ns · 5_093_616 B · 29_311 alloc · 133.9 MB/s |
| typed | unmarshal | nested | 2mb | `go-json/compat` | 5_849_258 ns · 4_550_861 B · 108_533 alloc · 308.3 MB/s | `segmentio/compat` | 12_868_265 ns · 17_616_771 B · 110_952 alloc · 140.1 MB/s |
| untyped | marshal | flat | 512b | `go-json/fast` | 1_525 ns · 576 B · 1 alloc · 349.7 MB/s | `jsoniter/compat` | 7_889 ns · 4_193 B · 65 alloc · 73.3 MB/s |
| untyped | marshal | flat | 8kb | `go-json/fast` | 23_829 ns · 9_728 B · 1 alloc · 408.3 MB/s | `jsoniter/compat` | 130_194 ns · 71_054 B · 1_089 alloc · 82.4 MB/s |
| untyped | marshal | flat | 64kb | `go-json/fast` | 181_275 ns · 81_920 B · 1 alloc · 440.0 MB/s | `jsoniter/compat` | 1_019_705 ns · 573_171 B · 8_737 alloc · 86.2 MB/s |
| untyped | marshal | flat | 512kb | `go-json/fast` | 1_472_606 ns · 647_172 B · 1 alloc · 436.8 MB/s | `jsoniter/compat` | 9_014_204 ns · 4_577_894 B · 69_906 alloc · 78.6 MB/s |
| untyped | marshal | flat | 2mb | `go-json/fast` | 5_738_741 ns · 2_588_682 B · 1 alloc · 450.9 MB/s | `jsoniter/compat` | 31_423_490 ns · 18_311_343 B · 279_622 alloc · 90.7 MB/s |
| untyped | marshal | nested | 512b | `go-json/fast` | 2_633 ns · 704 B · 1 alloc · 261.3 MB/s | `jsoniter/compat` | 12_789 ns · 7_317 B · 126 alloc · 57.3 MB/s |
| untyped | marshal | nested | 8kb | `go-json/fast` | 29_810 ns · 8_192 B · 1 alloc · 261.5 MB/s | `jsoniter/compat` | 128_079 ns · 72_662 B · 1_280 alloc · 63.4 MB/s |
| untyped | marshal | nested | 64kb | `go-json/fast` | 200_126 ns · 57_344 B · 1 alloc · 272.1 MB/s | `jsoniter/compat` | 823_674 ns · 470_170 B · 8_384 alloc · 68.0 MB/s |
| untyped | marshal | nested | 512kb | `go-json/fast` | 1_646_867 ns · 458_754 B · 1 alloc · 274.7 MB/s | `jsoniter/compat` | 6_886_508 ns · 3_773_183 B · 67_767 alloc · 67.3 MB/s |
| untyped | marshal | nested | 2mb | `go-json/fast` | 5_941_321 ns · 1_769_479 B · 1 alloc · 297.4 MB/s | `jsoniter/compat` | 26_425_360 ns · 14_806_480 B · 257_590 alloc · 68.3 MB/s |
| untyped | unmarshal | flat | 512b | `sonic/fast` | 2_825 ns · 2_904 B · 21 alloc · 204.6 MB/s | `encoding-json-v2/fast-relaxed` | 9_386 ns · 2_792 B · 103 alloc · 61.6 MB/s |
| untyped | unmarshal | flat | 8kb | `sonic/fast` | 39_738 ns · 43_704 B · 213 alloc · 270.1 MB/s | `encoding-json-v2/fast-relaxed` | 147_414 ns · 45_336 B · 1_654 alloc · 72.8 MB/s |
| untyped | unmarshal | flat | 64kb | `sonic/fast` | 310_320 ns · 350_393 B · 1_647 alloc · 283.2 MB/s | `encoding-json-v2/fast-relaxed` | 1_190_878 ns · 364_051 B · 13_151 alloc · 73.8 MB/s |
| untyped | unmarshal | flat | 512kb | `sonic/fast` | 2_663_720 ns · 3_840_713 B · 13_117 alloc · 266.1 MB/s | `encoding-json-v2/fast-relaxed` | 9_963_606 ns · 3_003_381 B · 105_155 alloc · 71.1 MB/s |
| untyped | unmarshal | flat | 2mb | `sonic/fast` | 12_452_084 ns · 29_288_318 B · 52_439 alloc · 228.8 MB/s | `encoding-json-v2/fast-relaxed` | 41_208_206 ns · 12_242_609 B · 420_260 alloc · 69.1 MB/s |
| untyped | unmarshal | nested | 512b | `sonic/fast` | 4_780 ns · 5_584 B · 33 alloc · 153.4 MB/s | `encoding-json-v2/fast-relaxed` | 15_487 ns · 5_840 B · 166 alloc · 47.3 MB/s |
| untyped | unmarshal | nested | 8kb | `sonic/fast` | 50_786 ns · 57_336 B · 238 alloc · 160.1 MB/s | `encoding-json-v2/fast-relaxed` | 178_719 ns · 66_536 B · 2_050 alloc · 45.5 MB/s |
| untyped | unmarshal | nested | 64kb | `sonic/fast` | 304_808 ns · 384_154 B · 1_478 alloc · 183.9 MB/s | `encoding-json-v2/fast-relaxed` | 1_267_130 ns · 463_403 B · 14_324 alloc · 44.2 MB/s |
| untyped | unmarshal | nested | 512kb | `sonic/fast` | 2_591_910 ns · 3_120_046 B · 11_785 alloc · 178.7 MB/s | `encoding-json-v2/fast-relaxed` | 10_288_176 ns · 3_794_915 B · 118_248 alloc · 45.0 MB/s |
| untyped | unmarshal | nested | 2mb | `sonic/fast` | 13_450_768 ns · 23_568_670 B · 44_556 alloc · 134.1 MB/s | `encoding-json-v2/fast-relaxed` | 42_079_183 ns · 14_528_521 B · 455_457 alloc · 42.9 MB/s |

