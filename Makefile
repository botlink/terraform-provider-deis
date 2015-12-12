default: bin

bin: generate
	@sh -c "'$(CURDIR)/scripts/build.sh'"

# generate runs `go generate` to build the dynamically generated
# source files.
generate:
	go generate ./...
