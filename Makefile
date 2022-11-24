SHELL := /bin/bash

.PHONY: machine-deps
machine-deps:
	@./scripts/setup.sh

.PHONY: test
test: machine-deps
	@mage -v test

.PHONY: testfast
testfast: machine-deps
	@mage -v testfast

.PHONY: dep
dep: machine-deps
	@mage dep

.PHONY: clean
clean: machine-deps
	@mage clean
