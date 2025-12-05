use std::fs;
use std::io;

/// Source of input: file path or raw text
#[derive(Debug)]
pub enum InputSource<'a> {
    File(&'a str),
    Text(&'a str),
}

/// Read lines from either a file path or raw text and return Vec<String>
pub fn read_lines(source: InputSource) -> Result<Vec<String>, io::Error> {
    let content = match source {
        InputSource::File(path) => fs::read_to_string(path)?,
        InputSource::Text(s) => s.to_string(),
    };
    Ok(content.lines().map(|l| l.to_string()).collect())
}
