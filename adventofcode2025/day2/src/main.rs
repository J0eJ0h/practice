use common::{read_lines, InputSource};

fn is_valid(s: &str, window: usize) -> bool {
    let len = s.chars().count();
    if len % window != 0 {
        return true;
    }

    for i in 0..len-window {
        if s.chars().nth(i) != s.chars().nth(i + window) {
            return true;
        }
    }
    false

}

fn part1_valid(s: &str) -> bool {
    let len = s.chars().count();
    if len % 2 != 0 {
        return true;
    }
    is_valid(s, len / 2)
}

fn part2_valid(s: &str) -> bool {
    let len = s.chars().count();
    for window in 1..=len/2 {
        if !is_valid(s, window) {
            return false;
        }
    }
    true
}

 fn part1(l: &str) {
    let mut total_p1 = 0;
    let mut total_p2 = 0;

     println!("Processing line: {}", l);
    for s in l.split(",") {
        if let Some((start, end)) = s.split_once("-") {
            let start_num: u64 = start.parse().unwrap_or(0);
            let end_num: u64 = end.parse().unwrap_or(0);
            for n in start_num..=end_num {
                let sn = n.to_string();
                let _valid = part1_valid(&sn);
                if !_valid {
                    total_p1 += n;
                    println!("    Number: {} Valid: {}", sn, _valid);
                }
                let _valid2 = part2_valid(&sn);
                if !_valid2 {
                    total_p2 += n;
                    println!("    Number: {} Valid (part2): {}", sn, _valid2);
                }   
            }
            println!("    Range: {} - {} ({} numbers) First Valid: {}.  Second Valid {}", start_num, end_num, end_num - start_num + 1, is_valid(start, 2), is_valid(end, 2));
        }
    }

    println!("Total sum of invalid numbers: {}", total_p1);
    println!("Total sum of invalid numbers (part2): {}", total_p2);
}

fn main() {
    println!("Hello, Advent of Code 2025 Day 2 (initial)!");

    let sample = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
    match read_lines(InputSource::Text(sample)) {
        Ok(lines) => {
            println!("Day2: {} lines read", lines.len());
            for (_, l) in lines.iter().enumerate() {
                part1(l);
            }
        }
        Err(e) => println!("Error: {}", e),
    }

    let input_file = "input.txt";
    match read_lines(InputSource::File(input_file)) {
        Ok(lines) => {
            println!("Day2: {} lines read", lines.len());
            for (_, l) in lines.iter().enumerate() {
                part1(l);
            }
        }
        Err(e) => println!("Error: {}", e),
    }
}
