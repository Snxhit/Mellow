# Mellow

Mellow is an experimental music scripting language written in Go.

## Status

Currently in a very primitive phase!<br>
Going forward I plan to rewrite everything and contain it into a web page so that live updates can work.

## Requirements

- Go (see `go.mod` for version)
- An available audio output device (obv!)

## User Instructions

Two test files are already included along with the executable (In release)
- `test.me`,
- `test2.me`

To try them, use:
```bash
./Mellow test.me
./Mellow test2.me
```

You may also modify the contents of the test files, or create your own `.me` files!

## Files

- `parser/` - converts source text into AST nodes
- `runtime/` - hot-reload runtime (broken due to how file editors work :<) and scheduler lifecycle
- `scheduler/` - loop scheduling and playback triggering
- `audio/` - audio engine and notes
- `ast/` - AST types

## Current Syntax

This is the only possible command for now.
```txt
loop note <NOTE> every <DURATION> ms
```

Where `<NOTE>` is replaced by:
- C4
- D4
- E4
- F4
- G4
- A4
- B4

(Check `audio/notes.go` for more details.)

And, `<DURATION>` is replaced by:
- Time in milliseconds (E.g.: 1000 ms (1 second), 100 ms (0.1 second), 50000 ms ())
