run:
    go run .
build:
    go build .
check:
    just build
fmt:
    go fmt .
test:
    ./run_tests.sh

