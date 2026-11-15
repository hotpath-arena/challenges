# HotPath Arena challenges

Public starters, sample Tests and a small public benchmark for each Challenge.
The statement, limits and the leaderboard are on the platform's Challenge page.

- [Access Log Aggregation](demo-log-aggregation/)

## Layout

Each Challenge directory holds:

| Path | Contents |
| --- | --- |
| `README.md` | Starter table and notes for this Challenge |
| `starter/<lang>/` | Starter projects for Go, Rust and C++ |
| `data/sample/` | Sample Tests: `.in` inputs and `.ans` answers |
| `bench/` | Optional public Benchmark Input with its `.ans`, at most 1 MiB |

Hidden Tests, hidden Benchmark Inputs and the reference solutions are not
published.

## Submit

Fork this repository to your personal GitHub account, or copy one starter
directory into a personal repository. Submit from the Challenge page. If you
keep this layout, enter the starter directory as the subdirectory, for example
`demo-log-aggregation/starter/go`. If the starter is at the root of your
repository, leave the subdirectory empty.

## Check locally

The platform runs your program with no arguments. Standard input is the input
file, a regular seekable file, and the working directory is empty. Use an
x86-64 machine that supports the Haswell CPU target. On another architecture,
omit the CPU target settings for a correctness check.

Run these from the root of your solution project. Set `challenge_dir` to the
absolute path of the Challenge directory in this repository.

**Rust.** Use stable Rust and a Cargo package with one binary target and no
registry or Git dependencies:

```sh
RUSTFLAGS="-C target-cpu=haswell -C force-frame-pointers=yes" \
  cargo +stable build --release --offline --locked
./target/release/YOUR_BINARY < "$challenge_dir/bench/example.in" > /tmp/out
diff -u "$challenge_dir/bench/example.ans" /tmp/out
```

Replace `YOUR_BINARY` with your Cargo binary target name.

**Go.** Use a root `main` package and a `go.mod` without `require` entries:

```sh
CGO_ENABLED=0 GOAMD64=v3 GOTOOLCHAIN=local GOPROXY=off GOFLAGS=-mod=mod \
  go build -trimpath -buildvcs=false -o /tmp/solution .
/tmp/solution < "$challenge_dir/bench/example.in" > /tmp/out
diff -u "$challenge_dir/bench/example.ans" /tmp/out
```

**C++.** Put a `CMakeLists.txt` with exactly one executable target at the root
of your solution project. Use no network dependencies:

```sh
CXXFLAGS="-march=haswell -fno-omit-frame-pointer" \
  cmake -S . -B /tmp/solution-cpp -DCMAKE_BUILD_TYPE=Release
cmake --build /tmp/solution-cpp
/tmp/solution-cpp/YOUR_BINARY < "$challenge_dir/bench/example.in" > /tmp/out
diff -u "$challenge_dir/bench/example.ans" /tmp/out
```

Replace `YOUR_BINARY` with your CMake executable target name. The Worker uses
GCC with static linking.
