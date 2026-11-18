# HotPath Arena challenges

Public starters, sample Tests and a small public benchmark for each Challenge.
The statement, limits and the leaderboard are on the platform's Challenge page.

- [Access Log Aggregation](01-log-aggregation/)

## Layout

Each Challenge directory holds:

| Path | Contents |
| --- | --- |
| `README.md` | Starter table and notes for this Challenge |
| `code/<lang>/` | Starter code for Go, Rust and C++ |
| `data/sample/` | Sample Tests: `.in` inputs and `.ans` answers |
| `bench/` | Optional public Benchmark Input with its `.ans`, at most 1 MiB |

Hidden Tests, hidden Benchmark Inputs and the reference solutions are not
published.

## Submit

1. Create a **private** repository on your personal GitHub account.
   Organization repositories are not accepted. Do not fork this repository:
   a fork of a public repository is public, so anyone could read your
   solution.
2. Copy the starter into the same path it has here,
   `<challenge>/code/<language>/`, for example `01-log-aggregation/code/go/`.
   The platform always builds that directory for the language you choose;
   a solution at the repository root or in another directory fails.
3. Sign in on the platform and install the HotPath Arena GitHub App on your
   personal account. Give it access to that repository.
4. Push your changes. On the Challenge page, choose the repository, the
   branch and the language, then submit. The platform captures the latest
   commit of that branch.

The platform downloads the whole repository at that commit, and the
compressed archive must be at most 10 MiB. Do not commit build output such
as `target/` or `build/`, inputs you generated, or other large files.

Track progress on the Submission page. A Score that would improve your best
is measured again at least one hour later. It appears on the Leaderboard
only after that confirmation pass.

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
