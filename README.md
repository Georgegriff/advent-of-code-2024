# Advent of Code 2024 - Go

Go to a folder for each days solution

- To run a solution `go run`
- To run tests `go test`
- To run code in watch mode `gow run`. Changes should re-run the code

> [!TIP]
> To run multiple days, go up a directory and run `go test ./...`

> [!TIP]
> Watch mode
> Install gow: https://github.com/mitranim/gow
> To run tests in watch mode `gow test`. Changes should re-run the tests

> [!TIP]
> To get terminal colors add an alias in zshrc go go test or gow test `go test -v . | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''`

```zsh
colorize() {
  awk '{
    if ($0 ~ /PASS/) {
      print "\033[32m" $0 "\033[0m"
    } else if ($0 ~ /FAIL/) {
      print "\033[31m" $0 "\033[0m"
    } else {
      print $0
    }
  }'
}

# misc
alias got="go test -v ./... | colorize"
alias gowt="gow test -v ./... | colorize"
```
