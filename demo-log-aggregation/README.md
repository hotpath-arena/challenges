# Access Log Aggregation

The statement, limits and leaderboard are on the platform's Challenge page.
Build and check commands for all languages are in the
[repository README](../README.md#check-locally).

## Starter projects

| Language | Submission subdirectory |
| --- | --- |
| Go | [`demo-log-aggregation/starter/go`](starter/go/) |
| Rust | [`demo-log-aggregation/starter/rust`](starter/rust/) |
| C++ | [`demo-log-aggregation/starter/cpp`](starter/cpp/) |

All three starters group requests by endpoint and then summarize each group.
They produce the expected output and give you a baseline to measure. Optimize
one language at a time and check that your changes preserve the output.

## Data

| Path | Rows |
| --- | ---: |
| `data/sample/totals.*` | 4 |
| `data/sample/ordering.*` | 6 |
| `bench/example.*` | 20,000 |

`bench/example.in` is a public Benchmark Input with its answer. The scored
inputs are hidden and larger.
