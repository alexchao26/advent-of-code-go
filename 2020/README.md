## 2020 Event

To run days for 2020, I'm planning on using `go test -run RegExpToMatchFunctionNames .` to make it cleaner to run any given examples.
`go run main.go -part <1 or 2>` will be usable to run the actual inputs for that day

Scripts have been added to the outer folder to setup files for a particular day more easily
```zsh
for ((i=1; i<26; i++)); do
go run ../scripts/template/template.go -day $i -year 2020
done
```

and to fetch input prompts
```go
// optional -day and -year flags with sensible defaults (today)
go run ../scripts/fetchers/cmd-inputs.go
```