#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
RESULTS="${ROOT}/results"
COUNT="${COUNT:-3}"
BENCHTIME="${BENCHTIME:-200ms}"
BENCH_FILTER="${BENCH_FILTER:-.}"
GOMAXPROCS="${GOMAXPROCS:-1}"

# Prefer local SDK installs; override with GO123_BIN / GO127_BIN if needed.
GO123_BIN="${GO123_BIN:-$HOME/go/go1.23.5/bin/go}"
GO127_BIN="${GO127_BIN:-$HOME/go/go1.27.3/bin/go}"
if [[ ! -x "${GO123_BIN}" ]]; then GO123_BIN="${GO123_BIN:-$HOME/go/go1.23.5/bin/go}"; fi
if [[ ! -x "${GO127_BIN}" ]]; then
  if [[ -x "$HOME/go/go1.27.0/bin/go" ]]; then
    GO127_BIN="$HOME/go/go1.27.0/bin/go"
  fi
fi

export GOPROXY="${GOPROXY:-https://proxy.golang.org,direct}"
export GOSUMDB="${GOSUMDB:-sum.golang.org}"
export GOMAXPROCS
# Avoid auto-upgrading the pinned local toolchains mid-run.
export GOTOOLCHAIN="${GOTOOLCHAIN:-local}"
# Inherited GOROOT (e.g. from IDE) breaks dual-toolchain runs.
unset GOROOT
# Inherited GOROOT (e.g. from IDE) breaks dual-toolchain runs.
unset GOROOT

mkdir -p "${RESULTS}"

if [[ ! -x "${GO123_BIN}" ]]; then
  echo "missing Go 1.23 binary: ${GO123_BIN}" >&2
  echo "set GO123_BIN to a go1.23.x binary" >&2
  exit 1
fi
if [[ ! -x "${GO127_BIN}" ]]; then
  echo "missing Go 1.27 binary: ${GO127_BIN}" >&2
  echo "set GO127_BIN to a go1.27.x binary" >&2
  exit 1
fi

{
  echo "suite go1.23 bin=${GO123_BIN}"
  "${GO123_BIN}" version
  echo "suite go1.27 bin=${GO127_BIN}"
  "${GO127_BIN}" version
  "${GO127_BIN}" env GOOS GOARCH GOAMD64 GOARM64
  printf 'GOMAXPROCS=%s\nCOUNT=%s\nBENCHTIME=%s\nBENCH_SCALE=%s\n' \
    "${GOMAXPROCS}" "${COUNT}" "${BENCHTIME}" "${BENCH_SCALE:-default}"
  sysctl -n machdep.cpu.brand_string 2>/dev/null || true
} >"${RESULTS}/environment.txt"

cd "${ROOT}"

# Separate caches: shared GOCACHE mixes object files across toolchains and breaks builds.
GOCACHE123="${GOCACHE123:-${ROOT}/.gocache/go123}"
GOCACHE127="${GOCACHE127:-${ROOT}/.gocache/go127}"
mkdir -p "${GOCACHE123}" "${GOCACHE127}"

echo "==> suite go1.23 (encoding/json + third-party)"
GOCACHE="${GOCACHE123}" "${GO123_BIN}" test -run='^$' -bench="${BENCH_FILTER}" -benchmem \
  -benchtime="${BENCHTIME}" -count="${COUNT}" ./... |
  tee "${RESULTS}/go123.txt"

echo "==> suite go1.27 (encoding/json/v2 + third-party)"
GOCACHE="${GOCACHE127}" "${GO127_BIN}" test -run='^$' -bench="${BENCH_FILTER}" -benchmem \
  -benchtime="${BENCHTIME}" -count="${COUNT}" ./... |
  tee "${RESULTS}/go127.txt"

echo "==> reports"
"${GO127_BIN}" run ./cmd/report -baseline=encoding-json \
  "${RESULTS}/go123.txt" | tee "${RESULTS}/summary_go123.md"

"${GO127_BIN}" run ./cmd/report -baseline=encoding-json-v2/default-safe \
  "${RESULTS}/go127.txt" | tee "${RESULTS}/summary_go127.md"

"${GO127_BIN}" run golang.org/x/perf/cmd/benchstat@latest \
  "${RESULTS}/go123.txt" "${RESULTS}/go127.txt" |
  tee "${RESULTS}/benchstat.txt"

echo "done: ${RESULTS}/summary_go123.md ${RESULTS}/summary_go127.md"
