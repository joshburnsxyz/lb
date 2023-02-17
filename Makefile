GOX := $(shell which go)
BIN := lb

lb:
	$(GOX) build \
		-v \
		-x \
		-o $(BIN)

clean:
	@rm -f $(BIN)