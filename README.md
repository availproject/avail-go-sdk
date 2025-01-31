# avail-go-sdk

# Note
This repo uses `mdbook` to generate documentation and `just` to run all the shell commands.
You can install both tools by running and executing `./install_dependencies.sh`. It should work as long as you have cargo installed.

# Release Strategy
This project uses [GitHub Flow](https://www.alexhyett.com/git-flow-github-flow/) to manage release and branches.

# Documentation
[Link](https://availproject.github.io/avail-go-sdk/) to documentation (web preview of examples)

# Logging
To enable logging add this to the `main` function.

```go
// Set log level based on the environment variable
level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
if err != nil {
    level = logrus.InfoLevel // Default to INFO if parsing fails
}
logrus.SetLevel(level)
logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
```

And run go command with `LOG_LEVEL` set to debug

```bash
LOG_LEVEL=debug go run .
```

# Commands
```bash
just # Runs `go run .`

just test # Run tests

just book-serve # Builds and serve the documentation

just book-deploy # Deploys the documentation

just fmt # formats files
```
