# YOLO

Life's too short for running tests. Prefix Maven or Gradle commands with `yolo` and live stress-free life!

![](https://media.giphy.com/media/3o7qDQ4kcSD1PLM3BK/giphy-downsized.gif)

```bash
$ yolo mvn package # runs mvn package -DskipTests
```

```bash
$ yolo gradle build # runs gradle build -x test
```

## Installation

`yolo` is published to Homebrew:

```bash
$ brew install maciejwalkowiak/tap/yolo
```

## Supported commands

`yolo` supports following commands: `mvn`, `./mvnw`, `mvnd`, `gradle`, `gradlew`

## How does it work

You can run any command with

```bash
$ yolo <command>
```

If the command matches one of the supported commands, it decorates it with flags that cause skipping tests, style checks, static analysis etc.

## Contribute

You are very welcome to contribute support for more languages/frameworks!

## PS

This project is made for fun. **Don't skip your tests.**

