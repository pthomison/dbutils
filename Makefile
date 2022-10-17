test:
	go test ./... -v -count=1

NEXT_TAG=$(shell exoskeleton rev -i $(shell git tag --sort version:refname | tail -n 1))
release:
	git tag $(NEXT_TAG)
	git push origin $(NEXT_TAG)