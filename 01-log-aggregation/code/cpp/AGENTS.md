<!-- perfarena-starter v1 slug=01-log-aggregation language=cpp target-cpu=haswell goamd64=v3 hardware-class=my-vps site=https://perfarena.valkirie-shield.xyz -->
# Access Log Aggregation: cpp Starter

You are working on one Starter of a performance Challenge on Hotpath Arena. The
platform builds your code, runs it on hidden inputs and scores the sum of the median
wall times. Lower is better.

- Statement: https://perfarena.valkirie-shield.xyz/challenges/01-log-aggregation (Markdown: https://perfarena.valkirie-shield.xyz/challenges/01-log-aggregation.md)
- Policy: LLMs and coding agents are allowed. You answer for what you submit. [Terms](https://perfarena.valkirie-shield.xyz/terms#llm)

## Limits

- Memory: 256 MiB. Maximum output: 327,680 bytes.
- The time limit per Hardware Class is on the Challenge page; a Run over it ends with `time_limit`.
- No network, in the build or in the run. No dependencies outside this folder.
- The program takes no arguments, reads a seekable file on stdin and writes to stdout, from an empty working directory.

## Loop

1. Edit the code in this folder.
2. `make check` builds with the judge's flags and runs every sample Test; it prints the failing Test and the first differing line.
3. `make bench` runs `make check`, then prints the median of 5 runs on each public Benchmark Input.
4. Submit only when the local median beats your previous one. A Submission costs quota.

Do not edit the flags in `Makefile`: the judge ignores the file, and import refuses a changed copy.

## What the Score is

The Score is the sum, over the hidden Benchmark Inputs, of the median wall time of your program, measured in a fresh microVM on the Hardware Class `my-vps`. Your local time is not the Score. Do not tune for the public `example` input: the hidden inputs have other distributions.

When the host is not Linux on x86_64, `make` drops the CPU target and static linking and prints a note. Correctness still matches the judge; times are only comparable between your own runs.

## Quota

10 Submissions a day on this Challenge for each Participant. A rejection carries `quota_daily` and `retry_at`. Use `make check` and `make bench` first; do not spend Submissions on compile errors and wrong answers.

## How to submit

Push a branch to your own GitHub repository, install the Hotpath Arena GitHub App on it and submit on the Challenge page: repository, branch, subdirectory `01-log-aggregation/code/cpp` (empty if this folder is the repository root), language `cpp`.
