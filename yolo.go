package main

import "os"
import "os/exec"
import "strings"
import "fmt"

import("github.com/common-nighthawk/go-figure")
import("github.com/fatih/color")

func main() {
    mavenAdditions := []string {
        "-DskipTests",
        "-Dspotbugs.skip=true",
        "-Dcheckstyle.skip=true",
        "-Dasciidoctor.skip=true",
        "-Dcheckstyle.skip=true",
        "-Ddocker.skip=true",
        "-Dinvoker.skip=true",
        "-Djacoco.skip=true",
        "-Djapicmp.skip=true",
        "-Djgassistant.skip=true",
        "-Dlicense.skip=true",
        "-Dmaven.javadoc.skip=true",
        "-Dpmd.skip=true",
        "-DskipITs=true",
        "-Dmaven.test.skip=true",
    }

    gradleAdditions := []string {
        "-x",
        "check",
    }

    myFigure := figure.NewColorFigure("YOLO", "", "cyan", true)
    myFigure.Print()
    fmt.Println("")
    color.Blue("================================")
    fmt.Println("")

    args := os.Args[1:]
    command := args[0]

    if startsWithAnyOf(command, []string { "mvn", "./mvnw", "mvnd" }) {
        args = append(args, mavenAdditions...)
    } else if startsWithAnyOf(command, []string { "gradle", "./gradlew" }) {
        args = append(args, gradleAdditions...)
    }

    cyan := color.New(color.FgCyan).SprintFunc()

    println(fmt.Sprintf("Running %s\n", cyan(strings.Join(args," "))))

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
