# Go Practice Problems

A progressive set of 10 hands-on Go challenges. Each problem has its own folder with a `main.go` containing TODOs.
Solve each in its folder. Create a branch per problem when collaborating.

## Quick Start
```bash
# from the repo root
go run ./1_hello_world
go run ./2_fizzbuzz
# ...
```

## Contribution Workflow
1. Fork and clone the repo.
2. Create a branch like `solve/2-fizzbuzz-mike`.
3. Implement the TODOs.
4. Run locally: `go run ./2_fizzbuzz` (and add tests if you want).
5. Open a Pull Request describing your approach and tradeoffs.

## Problems Overview

1. Hello World + CLI arg
2. FizzBuzz
3. File I/O transform (uppercase lines)
4. HTTP server `/hello`
5. JSON API `/users`
6. CLI flags `--name`
7. Goroutines & channels fan-out/fan-in
8. REST API client
9. Errors and custom error wrapping
10. In-memory URL shortener

Each `main.go` has instructions and example I/O.
