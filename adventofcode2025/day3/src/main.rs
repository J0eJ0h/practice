use std::str::FromStr;

use common::{read_lines, InputSource};
use num_bigint::BigUint;
use num_traits::Zero;

fn max_joltage(s: &String, len: usize) -> String {
    let mut max = 0;
    let mut max_index = 0;
    let mut total: String = String::new();
    for digit in 1..=len {
        for (i, c) in s.chars().enumerate() {
            if i <= max_index && digit > 1 {
                continue;
            }
            let val = c.to_digit(10).unwrap_or(0);
            if val > max && i < s.chars().count() - (len - digit) { // ignore last char since it can't lead
                max = val;
                max_index = i;  
            }
        }
        total = format!("{}{}", total, max);
        max = 0;
    }
    return total
}

fn part3(lines: &Vec<String>) {
    println!("Processing Day 3 with {} lines", lines.len());
    let mut total: BigUint = BigUint::zero();
    let mut total12: BigUint = BigUint::zero();
    for (i, l) in lines.iter().enumerate() {
        let mj = max_joltage(l, 2);
        let mj12 = max_joltage(l, 12);
        println!("  Line {}: Max Joltage: {}", i + 1, mj);
        // parse mj (string of digits) into BigUint and add to total
        let mj_big = BigUint::from_str(&mj).unwrap_or_else(|_| BigUint::zero());
        total += mj_big;
        // parse mj12 into BigUint and add to total12
        let mj12_big = BigUint::from_str(&mj12).unwrap_or_else(|_| BigUint::zero());
        total12 += mj12_big;
    }

    println!("Total Max Joltage Sum (len=2): {}", total);
    println!("Total Max Joltage Sum (len=12): {}", total12);
}

fn main() {
    println!("Hello, Advent of Code 2025 Day 3!");

    let sample = 
r#"987654321111111
811111111111119
234234234234278
818181911112111"#;
    match read_lines(InputSource::Text(sample)) {
        Ok(lines) => {
            part3(&lines);
        }
        Err(e) => println!("Error reading lines: {}", e),
    }

    let input_file = "input.txt";
    match read_lines(InputSource::File(input_file)) {
        Ok(lines) => {
            part3(&lines);
        }
        Err(e) => println!("Error reading lines from file {}: {}", input_file, e),
    }
}
