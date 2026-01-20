
.PHONY: all compile run clean

BINARY_NAME=tickTackToeGo
BUILD_DIR=bin

all: compile

compile:
	@echo "Compiling $(BINARY_NAME)..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Compilation complete. Executable: $(BUILD_DIR)/$(BINARY_NAME)"

run: compile
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build ..."
	rm -rf $(BUILD_DIR)
	@echo "Clean complete."
