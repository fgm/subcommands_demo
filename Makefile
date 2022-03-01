all: run clean

.PHONY: clean
clean:
	@echo "\nCleaning up"
	@rm -f demo

.PHONY: lint
lint:
	@echo "\nLinting"
	@gofmt -s -w . cmd
	@go vet ./...
	@staticcheck ./...

.PHONY: test
test:
	@echo "\nTesting"
	@go test -race ./cmd

demo: lint test
	@echo "\nBuilding"
	@go build -o demo .

.PHONY: run
run: demo
#	@echo "\n\"demo\" Without args = help and exit with error 2:"
#	@demo; true
#	@echo
#	@echo "\n\"demo help\" = help including important flags (one):"
#	@demo help
#	@echo "\n\"demo flags\": complete list available top-level flags (two)"
#	@demo flags
#	@echo "\n\"demo help top1\" = describe command top1:"
#	@demo help top1
#	@echo "\n\"demo flags top2\" = describe flags on command top2:"
#	@demo flags top2
#	@echo "\n\"demo commands\": list available commands (3 builtin ones)"
#	@demo commands
#	@echo "\nRunning \"demo 1\":"
#	@demo 1       -prefix hello
#	@echo "\nRunning \"demo -v top2\":"
#	@demo -v top2 -prefix meaning life 42
#	@echo "\nRunning \"demo top3\":"
#	@demo top3
#	@echo "\nRunning \"demo top3 commands\":"
#	@demo top3 commands
#	@echo "\nRunning \"demo top3 sub31\":"
#	@demo top3 sub31
#	@echo "\nRunning \"demo top3 sub32\":"
#	@demo top3 sub32
	@echo "\nRunning \"demo visit\":"
	@demo -v visit
