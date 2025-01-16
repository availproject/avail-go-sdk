run:
    go run .
build:
    go build -o /dev/null
check:
    just build
fmt:
    go fmt .
test:
    ./run_tests.sh

