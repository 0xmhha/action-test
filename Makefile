.PHONY: all clean gstable

all: gstable

gstable:
	@echo "Building gstable..."
	@mkdir -p build/bin
	go build -o build/bin/gstable ./cmd/gstable
	@chmod +x build/bin/gstable
	@echo "Build complete: build/bin/gstable"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf build
	@echo "Clean complete"
