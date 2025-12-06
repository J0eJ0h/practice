use common::{read_lines, InputSource};

fn main() {
    println!("Hello, Advent of Code 2025 Day 5!");

    let sample = "abcde\nabcd\n11111";
    match read_lines(InputSource::Text(sample)) {
        Ok(lines) => {
            println!("Day5: {} lines read", lines.len());
            for (i, l) in lines.iter().enumerate() {
                println!("  {}: {}", i + 1, l);
            }
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
}
