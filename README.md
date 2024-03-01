# `qpaser`
`qpaser` is a parser for log files produced by **Quake 3 Arena server**.

## Installation
You must have [Go version 1.21 or later](https://go.dev/) installed on your machine. Then, you can clone the repository and install the parser:
```bash
git clone https://github.com/victorl2/qparser.git
cd qparser
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
{
  "game_3": {
    "total_kills": 4,
    "players": [
      "Dono da Bola",
      "Mocinha",
      "Isgalamido",
      "Zeh"
    ],
    "kills": {
      "Dono da Bola": -1,
      "Isgalamido": 1,
      "Zeh": -2
    },
    "kills_by_means": {
      "MOD_FALLING": 1,
      "MOD_ROCKET": 1,
      "MOD_TRIGGER_HURT": 2
    }
  }
}
```

## How `qpaser` works ?
Matches of **Quake 3 Arena** produce a log file containing their kill feed. A log file can contain zero or more game matches. `qparser` reads the log file and parse/interpret the matches. It then summarizes the matches and prints a json to the standard output containing the _kill feed_ as shown above.