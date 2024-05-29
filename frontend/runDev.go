package main

import (
    "fmt"
    "os"
    "os/exec"
    "sync"
)

func runCommand(wg *sync.WaitGroup, name string, args ...string) {
    defer wg.Done()
    cmd := exec.Command(name, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("%s not found or failed to run: %v\n", name, err)
    }
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Running npm...")
        runCommand(&wg, "npm", "run", "dev")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Running pnpm...")
        runCommand(&wg, "pnpm", "run", "dev")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Running yarn...")
        runCommand(&wg, "yarn", "dev")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Running bun...")
        runCommand(&wg, "bun", "run", "dev")
    }()

    wg.Wait()
}