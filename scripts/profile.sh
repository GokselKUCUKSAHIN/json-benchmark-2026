#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
RESULTS="${ROOT}/results/profile"
BENCH_FILTER="${BENCH_FILTER:-BenchmarkMarshalConcrete/flat/64kb}"
BENCHTIME="${BENCHTIME:-1s}"
SUITE="${SUITE:-127}" # 123 | 127
GO123_BIN="${GO123_BIN:-$HOME/go/go1.23.5/bin/go}"
GO127_BIN="${GO127_BIN:-$HOME/go/go1.27.3/bin/go}"

mkdir -p "${RESULTS}"
cd "${ROOT}"

if [[ "${SUITE}" == "123" ]]; then
  GO_BIN="${GO123_BIN}"
else
  GO_BIN="${GO127_BIN}"
fi

export GOTOOLCHAIN=local
export GOPROXY="${GOPROXY:-https://proxy.golang.org,direct}"
export GOSUMDB="${GOSUMDB:-sum.golang.org}"

"${GO_BIN}" test -run='^$' -bench="${BENCH_FILTER}" -benchmem -benchtime="${BENCHTIME}" \
  -cpuprofile="${RESULTS}/cpu.pprof" \
  -memprofile="${RESULTS}/mem.pprof" \
  ./...

echo "cpu: go tool pprof ${RESULTS}/cpu.pprof"
echo "mem: go tool pprof ${RESULTS}/mem.pprof"
