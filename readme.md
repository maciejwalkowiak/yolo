# YOLO

![](https://media.giphy.com/media/3o7qDQ4kcSD1PLM3BK/giphy-downsized.gif)

## Git Support

Life's to short for messing around with `git add`, writing commit message. You just want to push, right? `yolo` has got you covered!

Run:

```bash
$ yolo git push ...
```

Yolo will translate it into `git add . && git commit -m "yolo" && git push ...`

## Maven & Gradle

Life's too short for running tests. Prefix Maven or Gradle commands with `yolo` and live stress-free life!

```bash
$ yolo mvn package # runs mvn package -DskipTests
```

```bash
$ yolo gradle build # runs gradle build -x check
```

## Installation

`yolo` is published to Homebrew:

```bash
$ brew install maciejwalkowiak/tap/yolo
```

## Supported commands

`yolo` supports following commands: `mvn`, `./mvnw`, `mvnd`, `gradle`, `gradlew`, `git push`

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

