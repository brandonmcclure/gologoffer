ifeq ($(OS),Windows_NT)
	SHELL := pwsh.exe
else
	SHELL := pwsh
endif

.SHELLFLAGS := -NoProfile -Command


.PHONY: build

all: build

build:
	go build -o logoffer
