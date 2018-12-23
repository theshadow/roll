.PHONY: all
all: build

ANTLR_BIN=$(which antlr)

formula/parser:
    $(ANTLR_BIN) -Dlanguage=Go -visitor -o formula/parser formula/Dice.g4

test: formula/parser
    go test
