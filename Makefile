# Application details
APP_NAME = system
INSTALL_DIR = /usr/local/bin

# Go build flags
GO_FLAGS = -o $(APP_NAME)

# Default target: build the application
all: build

# Build the application
build:
	go build $(GO_FLAGS) main.go

# Install the application to the system
install: build
	@echo "Installing $(APP_NAME) to $(INSTALL_DIR)..."
	@install -Dm755 $(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)
	@echo "Installation complete!"

# Uninstall the application from the system
uninstall:
	@echo "Uninstalling $(APP_NAME) from $(INSTALL_DIR)..."
	@rm -f $(INSTALL_DIR)/$(APP_NAME)
	@echo "Uninstallation complete!"

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	@rm -f $(APP_NAME)
	@echo "Clean complete!"

