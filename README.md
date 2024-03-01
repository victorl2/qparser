# `qpaser`
`qpaser` is a parser for log files produced by **Quake 3 Arena server**.

## Installation
You must have [Go version 1.21 or later](https://go.dev/) installed on your machine. Then, you can clone the repository and install the parser:
```bash
git clone https://github.com/victorl2/qparser.git
go build -o qparser
```

## Usage 
```sh
# help message with all available commands
./qpaser --help

# summarize all matches in the file
./qpaser file.log

# summarize a single match
./qpaser file.log --game 3
```

output:
```json
```

## How `qpaser` works ?
Matches of **Quake 3 Arena** produce a log file containing their kill feed. A single log file can contain zero or more game matches. 

`qparser` reads the log file and parse/interpret the matches. It then summarizes the matches and prints the result to the standard output containing the _kill feed_.