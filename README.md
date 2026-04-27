<div align="center">
  <h1>Mellow</h1>
  <img src="https://res.cloudinary.com/dp7g5aflo/image/upload/v1777310046/mellow_banner_1024-768_w1vjrq.png">
  Mellow is a programming language that can be used to create music on the fly.
  
  ![Time Tracking](https://img.shields.io/badge/Mellow-2h%2049m-critical?logo=neovim&style=plastic)
  ![GitHub Stars](https://img.shields.io/github/stars/Snxhit/Mellow?style=plastic)
  ![GitHub Forks](https://img.shields.io/github/forks/Snxhit/Mellow?style=plastic)
  ![GitHub Issues](https://img.shields.io/github/issues/Snxhit/Mellow?style=plastic)

</div>

## Status

Currently in a very primitive phase.
Going forward, I plan to rewrite everything and move it into a web page so live updates can work.

## Requirements

- Go (see `go.mod` for version)
- An available audio output device (obv!)

## User Instructions

Two test files are already included:
- `test.me`
- `test2.me`

To try them, use:

```bash
./Mellow test.me
./Mellow test2.me
```

You can also modify the test files, or create your own `.me` files.

## Files

- `parser/`: converts source text into AST nodes
- `runtime/`: hot-reload runtime (broken due to how file editors work :<) and scheduler lifecycle
- `scheduler/`: loop scheduling and playback triggering
- `audio/`: audio engine and notes
- `ast/`: AST types

## Current Syntax

This is the only command for now.

```txt
loop note <NOTE> every <DURATION> ms
```

`<NOTE>` can be:
- C4
- D4
- E4
- F4
- G4
- A4
- B4

`<DURATION>` is time in milliseconds.
Examples: `1000 ms`, `100 ms`, `50000 ms`.
