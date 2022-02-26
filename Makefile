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
#	@echo "\n\"demo help\" = help and exit with success:"
#	@demo help
	@echo "\n\"demo flags\": list available top-level flags (one)"
	@demo flags
	@echo "\n\"demo help top1\" = describe command top1:"
	@demo help top1
	@echo "\n\"demo flags top2\" = describe flags on command top2:"
	@demo flags top2
#	@echo "\n\"demo commands\": list available commands (3 builtin ones)"
#	@demo commands
	@echo "\nRunning \"demo top1\":"
	@demo top1    -prefix hello
	@echo "\nRunning \"demo -v top2\":"
	@demo -v top2 -prefix meaning life 42
