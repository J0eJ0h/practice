use std::collections::VecDeque;

use common::{read_lines, InputSource};

#[derive(Clone)]
struct Interval {
    start: usize,
    end: usize,
}

impl Interval {
    fn new(start: usize, end: usize) -> Self {
        Interval { start, end }
    }
   fn from_str(s: &str) -> Self {
        if let Some((start_str, end_str)) = s.split_once("-") {
            let start = start_str.parse().unwrap_or(0);
            let end = end_str.parse().unwrap_or(0);
            return Interval { start, end }
        } 
        Interval::new(0, 0)
        
    }

    fn constains(&self, point: usize) -> bool {
        point >= self.start && point <= self.end
    }
}
// Just a collection of intervals in a vector.  Nothing fancy
struct IntervalSet {
    intervals: Vec<Interval>,
}

impl IntervalSet {
    fn add_interval(&mut self, interval: Interval) {
        self.intervals.push(interval);
    }

    fn contains(&self, point: usize) -> bool {
        for interval in &self.intervals {
            if interval.constains(point) {
                return true;
            }
        }
        false
    }
}

// Set of intervals with overlapping intervals flattened\merged into non-overlapping intervals
// To keep life easy, it assumes it's always in a flat state.
struct FlattenedIntervalSet {
    intervals: Vec<Interval>,
}

impl FlattenedIntervalSet {
    fn new() -> Self {
        FlattenedIntervalSet { intervals: Vec::new() }
    }

    fn get_intervals(&self) -> &Vec<Interval> {
        &self.intervals
    }

    fn get_contains_index(&self, point: usize) -> Option<usize> {
        for (i, interval) in self.intervals.iter().enumerate() {
            if interval.constains(point) {
                return Some(i);
            }
        }
        None
    }

    fn merge_intervals(base_interval: Interval, interval: &Interval) -> Interval {
        let new_start = std::cmp::min(base_interval.start, interval.start);
        let new_end = std::cmp::max(base_interval.end, interval.end);
        return Interval::new(new_start, new_end);
    }

    fn add_interval(&mut self, interval: Interval) {
        let mut to_merge: VecDeque<usize> = VecDeque::new();
        // do push front so we can avoid sorting the list to remove from the back by index first
        for (i, existing_interval) in self.intervals.iter_mut().enumerate() {
            if !(interval.end < existing_interval.start || interval.start > existing_interval.end) {
                // Overlap detected.  Replace with merged interval
                to_merge.push_front(i);  // We do this to make sure the list is revese sorted for the next operation
            }
        }
        // remove intervals that and merge into uber interval
        let mut merged_interval = interval;
        // We assume the to_merge list is reverse sorted by index so we don't shuffle indices as we remove
        for i in to_merge {
            let existing_interval = self.intervals.remove(i);
            merged_interval = FlattenedIntervalSet::merge_intervals(merged_interval, &existing_interval);
        }
        self.intervals.push(merged_interval);
    }

    fn contains(&self, point: usize) -> bool {
        for interval in &self.intervals {
            if interval.constains(point) {
                return true;
            }
        }
        false 
    }
}

fn part1(source: InputSource) {
    println!("Processing Day 5");
    match read_lines(source) {
        Ok(lines) => {
            let mut interval_set = IntervalSet { intervals: Vec::new() };
            let mut flattened_set = FlattenedIntervalSet::new();
            let mut intervals = true;
            let mut fresh = 0;
            for l in lines {
                if l == "" {
                    intervals = false;
                    continue;
                }
                if intervals {
                    let interval = Interval::from_str(&l);
                    interval_set.add_interval(interval.clone());
                    flattened_set.add_interval(interval);
                } else {
                    let point = l.parse::<usize>().unwrap_or(0);
                    let contains = interval_set.contains(point);
                    println!("  Point {} contained in intervals: {}", point, contains);
                    if contains {
                        fresh += 1;
                    }
                }
            }
            println!("Total fresh points: {}", fresh);
            let mut total_covered = 0;
            for interval in flattened_set.get_intervals() {
                total_covered += interval.end - interval.start + 1;
            }
            println!("Total covered points by intervals: {}", total_covered);
            /* 
            // Just print the intervals for now
            for (i, interval) in interval_set.intervals.iter().enumerate() {
                println!("  Interval {}: {} - {}", i + 1, interval.start, interval.end);
            } */
        }
        Err(e) => println!("Error reading lines: {}", e),
    }
}

fn main() {
    println!("Hello, Advent of Code 2025 Day 5!");

    let sample = r#"3-5
10-14
16-20
12-18

1
5
8
11
17
32"#;

    part1(InputSource::Text(sample));
    part1(InputSource::File("input.txt"));
}
