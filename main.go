package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"Mellow/parser"
	"Mellow/runtime"
)

func load(rt *runtime.Runtime, file string) {
	src, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}

	pf, err := parser.Parse(string(src))
	if err != nil {
		fmt.Println("Parse err:", err)
		return
	}

	fmt.Println("Reloading...")
	rt.Load(pf)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: mellow <file>")
		return
	}

	file := os.Args[1]
	abs, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	rt := runtime.New()
	go rt.Run()

	load(rt, abs)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	dir := filepath.Dir(abs)
	if err := watcher.Add(dir); err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if filepath.Clean(event.Name) == abs &&
					event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0 {
					load(rt, abs)
				}
			case err := <-watcher.Errors:
				fmt.Println("Watcher err:", err)
			}
		}
	}()

	select {}
}

