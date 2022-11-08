package main

import "encoding/json"
import "errors"
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
        "-X",
        "test",
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
    } else if startsWithAnyOf(command, []string { "gradle", "./gradlew", "mvnd" }) {
        args = append(args, gradleAdditions...)
    } else if startsWithAnyOf(command, []string { "npm", "yarn" }) {
        if npmSuppressTest() {
            defer npmRestoreTest()
        }
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

func npmSuppressTest() bool {
    stat, err := os.Stat("package.json")
    if errors.Is(err, os.ErrNotExist) {
        panic("No 'package.json' file found in this directory")
    }

    // read in the file
    contents, _ := os.ReadFile("package.json")

    // parse the json into arbitrary nested map-cast-able interface
    var packageJson interface{}
    // return early if its not valid json
    if nil != json.Unmarshal(contents, &packageJson) {
        panic("Could not parse 'package.json'")
    }

    // return early if missing scripts or scripts.test sections
    scriptsBlock := packageJson.(map[string]interface{})["scripts"]
    if nil == scriptsBlock {
        return false
    }

    testScript := scriptsBlock.(map[string]interface{})["test"]
    if nil == testScript {
        return false
    }

    // override test task with 'skipped' message
    scriptsBlock.(map[string]interface{})["test"] = "echo skipped"

    // since we have edited the package, now we can back up the original
    if nil != os.WriteFile("package.json.tmp", contents, stat.Mode()) {
        panic("Could not back up 'package.json' to 'package.json.tmp'")
    }

    newContents, _ := json.Marshal(packageJson)
    _ = os.WriteFile("package.json", newContents, stat.Mode())
    return true
}

func npmRestoreTest() {
    fmt.Println("restoring")
    e := os.Rename("package.json.tmp", "package.json")
    if nil != e {
        _, _ = fmt.Fprintln(os.Stderr, "Could not restore package.json, test script is still skipped")
    }
}
