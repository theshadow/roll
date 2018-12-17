.PHONY: all
all: build

ANTLR_BIN=$(which antlr)

formula/parser:
    cd formula && $(ANTLR_BIN) -Dlanguage=Go -o parser Dice.g4

test: formula/parser
    go test
