use common::{read_lines, InputSource};



fn main() {
    println!("Hello, Advent of Code 2025 Day 4!");

    let sample = "a,b,c\n1,2,3";
    match read_lines(InputSource::Text(sample)) {
        Ok(lines) => {
            println!("Day4: {} lines read", lines.len());
            for (i, l) in lines.iter().enumerate() {
                println!("  {}: {}", i + 1, l);
            }
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
}
