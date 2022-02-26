all: run clean

.PHONY: clean
clean:
	@echo "\nCleaning up"
	@rm -f demo

.PHONY: lint
lint:
	@echo "\nLinting"
	@go vet ./...
	@staticcheck ./...

demo: lint
	@echo "\nBuilding"
	@go build -o demo .

.PHONY: run
run: demo
#	@echo "\n\"demo\" Without args = help and exit with error 2:"
#	@demo; true
#	@echo
	@echo "\n\"demo help\" = help and exit with success:"
	@demo help
#	@echo "\n\"demo commands\": list available commands (3 builtin ones)"
#	@demo commands
#	@echo "\n\"demo flags\": list available top-level flags (none)"
#	@demo flags
#	@echo "\nRunning \"demo top1\":"
#	@demo top1
#	@echo "\nRunning \"demo top2 foo bar\":"
#	@demo top2 foo bar
