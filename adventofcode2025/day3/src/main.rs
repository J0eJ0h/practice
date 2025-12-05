use common::{read_lines, InputSource};

fn main() {
    println!("Hello, Advent of Code 2025 Day 3!");

    let sample = "day3-a\nday3-b\nday3-c";
    match read_lines(InputSource::Text(sample)) {
        Ok(lines) => {
            println!("Day3: {} lines read", lines.len());
            for (i, l) in lines.iter().enumerate() {
                println!("  {}: {}", i + 1, l);
            }
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
}
