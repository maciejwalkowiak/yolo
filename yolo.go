package main

import "os"
import "os/exec"
import "strings"

func main() {
    args := os.Args[1:]
    command := args[0]

    if startsWithAnyOf(command, []string { "mvn", "./mvnw", "mvnd" }) {
        args = append(args, "-DskipTests")
    } else if startsWithAnyOf(command, []string { "gradle", "./gradlew", "mvnd" }) {
        args = append(args, "-X", "test")
    }

    cmd := exec.Command(command, args[1:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    _ = cmd.Run()
}

func startsWithAnyOf(command string, prefixes []string) bool {
    for i := range prefixes {
        if strings.HasPrefix(command, prefixes[i]) {
            return true
        }
    }
    return false
}