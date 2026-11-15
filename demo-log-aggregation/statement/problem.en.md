# Access Log Aggregation

A service records one line for each HTTP request. Summarize its traffic by
endpoint. For each endpoint, report the request count, error count, total
response bytes, and highest latency. A request is an error if its status is
500 or above.

## Input

Read `N` rows from standard input, with `1 ≤ N ≤ 1,750,000`. There is no header
or separate count. Each row is `timestamp_ms,endpoint,status,bytes,latency_us`
and ends with a single LF byte (`0x0A`), including the last row. Commas are
literal separators. There are no quotes, spaces, carriage returns, or blank
lines.

| Field | Allowed values |
| --- | --- |
| `timestamp_ms` | 0–9,999,999,999,999, in milliseconds since the Unix epoch |
| `endpoint` | 1–24 ASCII bytes from `[A-Za-z0-9_/-]`, case-sensitive |
| `status` | 100–599 |
| `bytes` | 0–4,294,967,295 |
| `latency_us` | 0–9,999,999, in microseconds |

Numbers use ASCII decimal digits with no sign or leading zeros, except for
the single digit `0`. There are at most 4,096 distinct endpoints. Rows can
arrive in any order, and identical rows count as separate requests.

## Output

Write one line per distinct endpoint in ascending ASCII byte order. If one
endpoint is a prefix of another, the shorter one comes first. Each line is
`endpoint count error_count total_bytes max_latency_us`, with exactly one
ASCII space between fields and an LF at the end. Write numbers in decimal
without signs or leading zeros. All numeric results fit in an unsigned 64-bit
integer.

## Samples

### Totals

Input

```text
100,/z,200,12,25
102,/a,499,5,50
104,/z,503,3,90
110,/a,500,7,40
```

Output

```text
/a 2 1 12 50
/z 2 1 15 90
```

`/a` has one error because 500 counts and 499 does not. The maximum latency
uses both requests, so it is 50 even though the error took 40.

### Ordering

Input

```text
0,a/b,200,0,9
1,a-b,500,2,8
2,B,599,1,3
3,a,304,5,4
4,a,200,5,2
5,B,400,2,7
```

Output

```text
B 2 1 3 7
a 2 0 10 4
a-b 1 1 2 8
a/b 1 0 0 9
```

`B` sorts before `a` because its ASCII byte is smaller. The endpoint `a`
comes before its extensions `a-b` and `a/b`.

## Case sizes

| Role | Case | Rows |
| --- | --- | ---: |
| Public Test | Totals | 4 |
| Public Test | Ordering | 6 |
| Public Benchmark Input | Example | 100,000 |

## Running locally

Clone the public challenges repository and set `package_dir` to the absolute
path to its `demo-log-aggregation` directory. It contains `bench/example.in` and `bench/example.ans`. Use the build
commands below from the root of your solution project on an x86-64 machine
that supports the platform's Haswell CPU target.

For Rust, use stable Rust and a Cargo package with one binary target and no
registry or Git dependencies:

```sh
RUSTFLAGS="-C target-cpu=haswell -C force-frame-pointers=yes" \
  cargo +stable build --release --offline
./target/release/YOUR_BINARY < "$package_dir/bench/example.in" > /tmp/access-log.out
diff -u "$package_dir/bench/example.ans" /tmp/access-log.out
```

Replace `YOUR_BINARY` with your Cargo binary target name. For Go, use a
root `main` package and a `go.mod` without `require` entries:

```sh
CGO_ENABLED=0 GOAMD64=v3 GOTOOLCHAIN=local GOPROXY=off GOFLAGS=-mod=mod \
  go build -trimpath -buildvcs=false -o /tmp/access-log .
/tmp/access-log < "$package_dir/bench/example.in" > /tmp/access-log.out
diff -u "$package_dir/bench/example.ans" /tmp/access-log.out
```

For C++, put a `CMakeLists.txt` with exactly one executable target at the
root of your solution project. Use no network dependencies:

```sh
CXXFLAGS="-march=haswell -fno-omit-frame-pointer" \
  cmake -S . -B /tmp/access-log-cpp -DCMAKE_BUILD_TYPE=Release
cmake --build /tmp/access-log-cpp
/tmp/access-log-cpp/YOUR_BINARY < "$package_dir/bench/example.in" > /tmp/access-log.out
diff -u "$package_dir/bench/example.ans" /tmp/access-log.out
```

Replace `YOUR_BINARY` with your CMake executable target name. The Worker
uses GCC with static linking. On a different CPU architecture, omit the CPU
target setting for a local correctness check.
