use std::collections::BTreeMap;
use std::io::{self, BufRead, BufWriter, Write};

struct Request {
    status: u64,
    bytes: u64,
    latency: u64,
}

fn main() -> io::Result<()> {
    // Keep each request so reading and summarizing are separate steps.
    let mut groups: BTreeMap<String, Vec<Request>> = BTreeMap::new();
    for line in io::stdin().lock().lines() {
        let line = line?;
        let fields: Vec<_> = line.split(',').collect();
        groups
            .entry(fields[1].to_owned())
            .or_default()
            .push(Request {
                status: fields[2].parse().expect("valid status"),
                bytes: fields[3].parse().expect("valid byte count"),
                latency: fields[4].parse().expect("valid latency"),
            });
    }

    let mut output = BufWriter::new(io::stdout().lock());
    for (endpoint, requests) in groups {
        let mut errors = 0;
        let mut bytes = 0;
        let mut max_latency = 0;
        for request in &requests {
            if request.status >= 500 {
                errors += 1;
            }
            bytes += request.bytes;
            max_latency = max_latency.max(request.latency);
        }
        writeln!(
            output,
            "{endpoint} {} {errors} {bytes} {max_latency}",
            requests.len(),
        )?;
    }
    output.flush()
}
