GO := go
PROGRAM_FLAGS := -d
PATH_FILE := src/cmd/cmd.go
EXIT_BINERY := ./dnslab

run:
	$(GO) run $(PATH_FILE) $(PROGRAM_FLAGS)

build:
	$(GO) build -o $(EXIT_BINERY) $(PATH_FILE)

release:
	sh build.sh