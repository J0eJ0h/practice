use std::io;
use std::env;
use common::{read_lines, InputSource};

use std::str::FromStr;
use std::num::ParseIntError;

#[derive(Debug, PartialEq)]
enum Direction {
    Right(i32),
    Left(i32),
}

impl FromStr for Direction {
    // Define the associated error type if parsing fails
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.is_empty() {
            return Err("Input string is empty".to_string());
        }

        let (direction_char, magnitude_str) = s.split_at(1);
        
        // Parse the magnitude part as an unsigned 32-bit integer (u32)
        // You can change u32 to i32, u64, etc., as needed.
        let magnitude: i32 = magnitude_str.parse()
            .map_err(|e: ParseIntError| format!("Failed to parse integer: {}", e))?;

        // Match the direction character to create the enum variant
        match direction_char {
            "r" | "R" => Ok(Direction::Right(magnitude)),
            "l" | "L" => Ok(Direction::Left(magnitude)),
            _ => Err(format!("Unknown direction character: {}", direction_char)),
        }
    }
}

impl Direction {
   pub fn is_right(&self) -> bool {
        matches!(self, Direction::Right(_))
    }
    /* 
    pub fn is_left(&self) -> bool {
        matches!(self, Direction::Left(_))
    }*/

    pub fn magnitude(&self) -> i32 {
        match self {
            Direction::Right(mag) => *mag,
            Direction::Left(mag) => *mag,
        }
    }
}

fn part1(source: InputSource) -> Result<(), io::Error> {
    let mut current: i32 = 50;
    let mut pw  = 0;
    let mut clicks = 0;
    match read_lines(source) {
        Ok(lines) => {
            println!("Lines from string ({}):", lines.len());
            for (i, l) in lines.iter().enumerate() {
                match l.parse::<Direction>() {
                    Ok(dir) => {
                        //println!("  Line {}: {:?} (magnitude: {})", i + 1, dir, dir.magnitude());
                        if dir.is_right() {
                            current += dir.magnitude();
                            
                            let cl = current / 100;
                            if cl > 0 {
                                println!("  Line {}: Wrapped around {} times", i + 1, cl);
                            }
                            clicks += cl;

                            if current > 0 && current % 100 == 0 {
                                pw += 1;
                                clicks -= 1;
                                println!("  Line {}: Hit password position!", i + 1);
                            }
                                
                            
                            current %= 100;
                        } else {
                            let l = if current == 0 { 0 } else {1};
                            current -= dir.magnitude();
                            if current < 0 {
                                let cl = l - (current / 100);
                                println!("  Line {}: Wrapped around {} times", i + 1, cl);
                                clicks += cl; 
                                if current % 100 == 0 {
                                    clicks -= 1;
                                    println!("  Line {}: Hit password position!", i + 1);
                                }
                            }
                            current %= 100;
                            if current < 0 {
                                current += 100;
                            }
                            
                            if current % 100 == 0 {
                                pw += 1; 
                            }
                        }
                    },
                    Err(e) => println!("  Line {}: Error parsing direction: {}", i + 1, e),
                }
            }
        }
        Err(e) => println!("Error reading from string: {}", e),
    }
    println!("Final position: {}, P1 Password: {}, P2 Password: {}", current, pw, pw+clicks);

    Ok(())
}


fn main() {
    println!("Hello, Advent of Code 2025 Day 1!");
    let current_dir = env::current_dir();
    match current_dir {
        Err(e) => println!("Error getting current working directory: {}", e),
        Ok(dir) => println!("The current working directory is: {}", dir.display())
    }

    // Example: read from a string
    let sample = 
r#"L50
L200
R100
L50
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"#;
    part1(InputSource::Text(sample)).unwrap();
    // Example: read from a file (non-fatal if missing)
    let path = "input.txt";
    part1(  InputSource::File(path)).unwrap();

    let r8_str = "r8";
    let l12_str = "l12";
    let invalid_str = "u5";

    // Parse the strings
    let dir1: Result<Direction, _> = r8_str.parse();
    let dir2: Result<Direction, _> = l12_str.parse();
    let dir3: Result<Direction, _> = invalid_str.parse();

    // Handle the results
    match dir1 {
        Ok(ref d) => println!("{:?}", d), // Output: Right(8)
        Err(ref e) => println!("Error: {}", e),
    }

    match dir2 {
        Ok(d) => println!("{:?}", d), // Output: Left(12)
        Err(e) => println!("Error: {}", e),
    }

    match dir3 {
        Ok(d) => println!("{:?}", d),
        Err(e) => println!("Error: {}", e), // Output: Error: Unknown direction character: u
    }

    println!("{:?}", dir1.unwrap().magnitude());
}
