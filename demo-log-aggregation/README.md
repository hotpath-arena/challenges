# Access Log Aggregation

Read the [problem statement](statement/problem.en.md) for the input and
output format and [build requirements](statement/problem.en.md#running-locally).

## Starter projects

| Language | Submission subdirectory |
| --- | --- |
| Go | [`demo-log-aggregation/starter/go`](starter/go/) |
| Rust | [`demo-log-aggregation/starter/rust`](starter/rust/) |
| C++ | [`demo-log-aggregation/starter/cpp`](starter/cpp/) |

Fork the repository to your personal GitHub account, or copy one starter
directory into a personal repository. If you keep the fork's layout, use the
subdirectory above when submitting. If you copy the starter to the root of
your repository, leave the submission subdirectory empty.

All three starters group requests by endpoint and then summarize each group.
They produce the expected output and give you a baseline to measure. Optimize
one language at a time and check that your changes preserve the output.

## Check a starter locally

From this directory, build the starter for your language. These commands keep
build output outside your checkout.

```sh
build_dir=$(mktemp -d)

# Go
(cd starter/go && go build -o "$build_dir/access-log-go" .)

# Rust
cargo build --release --offline --locked --manifest-path starter/rust/Cargo.toml \
  --target-dir "$build_dir/rust"

# C++
cmake -S starter/cpp -B "$build_dir/cpp" -DCMAKE_BUILD_TYPE=Release
cmake --build "$build_dir/cpp"
```

For example, check the Go starter against the public benchmark:

```sh
"$build_dir/access-log-go" < bench/example.in > "$build_dir/output"
diff -u bench/example.ans "$build_dir/output"
```

The Rust and C++ binaries are at `$build_dir/rust/release/access-log` and
`$build_dir/cpp/access-log`. Small examples are in `data/sample/`. Public
limits are in [limits.json](limits.json).
