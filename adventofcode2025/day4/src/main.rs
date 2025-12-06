use common::{read_lines, InputSource};



struct Floor {
    _floor: Vec<Vec<char>>,
}

impl Floor
{
    fn new() -> Self {
        Floor {
            _floor: Vec::new(),
        }
    }

    fn add_row(&mut self, row: &str) {
        self._floor.push(row.chars().collect());
    }

    fn reachable_count(&mut self) -> usize {
        let mut count = 0;
        for i in 0..self._floor.len() {
            for j in 0..self._floor[i].len() {
                if self._floor[i][j] == '@' && self.reachable(i, j) {
                    count += 1;
                }
            }
        }
        count
    }   

    fn reachable_count_recursive(&mut self) -> usize {
        let mut count = 0;
        loop {
            for i in 0..self._floor.len() {
                for j in 0..self._floor[i].len() {
                    if self._floor[i][j] == 'x' {
                        self._floor[i][j] = ',';
                    }
                }
            }
            let rc = self.reachable_count();
            if rc == 0 {
                break;
            }
            count += rc;
            println!("Intermediate reachable count: {}", count);
        }
        count
    }

    fn reachable(&mut self, x: usize, y: usize) -> bool {
        if self.neighbor_count(x, y) < 4 {
            self._floor[x][y] = 'x';
            return true;
        }
        false
    }

    fn neighbor_count(&self, x: usize, y: usize) -> usize {
        let mut count = 0;
        for _i in -1isize..=1 {
            for _j  in -1isize..=1 {
                if _i == 0 && _j == 0 {
                    continue;
                }
                let i = x as isize + _i;
                let j = y as isize + _j;

                if i < 0 || j < 0 || i >= self._floor.len() as isize || j >= self._floor[i as usize].len() as isize { 
                    continue;
                }
                if self._floor[i as usize][j as usize] == '@' || self._floor[i as usize][j as usize] == 'x' {
                    count += 1;
                }
            }
        }
        count
    }

    fn print(&self) {
        for row in &self._floor {
            let line: String = row.iter().collect();
            println!("{}", line);
        }
    }
}

fn part1(source: InputSource) -> Result<(), std::io::Error> {
    let mut floor = Floor::new();
    match read_lines(source) {
        Ok(lines) => {
            println!("Lines from string ({}):", lines.len());
            for (i, l) in lines.iter().enumerate() {
                floor.add_row(l);
                println!("  Line {}: {}", i + 1, l);
            }
            floor.print();
            let reachable = floor.reachable_count();
            floor.print();
            println!("Total reachable positions: {}", reachable);
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
    Ok(())
}

fn part2(source: InputSource) -> Result<(), std::io::Error> {
    let mut floor = Floor::new();
    match read_lines(source) {
        Ok(lines) => {
            println!("Lines from string ({}):", lines.len());
            for (i, l) in lines.iter().enumerate() {
                floor.add_row(l);
                println!("  Line {}: {}", i + 1, l);
            }
            floor.print();
            let reachable = floor.reachable_count_recursive();
            floor.print();
            println!("Total reachable positions (recursive): {}", reachable);
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
    Ok(())
}
fn main() {
    println!("Hello, Advent of Code 2025 Day 4!");

    let sample = 
r#"..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@."#;
    
    part1(InputSource::Text(sample)).unwrap();
    part2(InputSource::Text(sample)).unwrap();
    part2(InputSource::File("input.txt")).unwrap();
}
