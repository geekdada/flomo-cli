PKGS := github.com/geekdada/flomo-cli
SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))
GO := go

test:
	$(GO) test $(PKGS)

build:
	$(GO) build -o ./artifacts $(PKGS)
