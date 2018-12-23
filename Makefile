.PHONY: all
all: build

ANTLR_BIN=$(which antlr)

formula/parser:
    $(ANTLR_BIN) -Dlanguage=Go -visitor -o formula/parser formula/Dice.g4

.PHONY: fmt
fmt:
    go fmt ../...

.PHONY: static-analysis
static-analysis:
    go vet ../...

.PHONY: static-analysis
test: formula/parser
    go test

.PHONY: build
build: static-analysis test formula/parser