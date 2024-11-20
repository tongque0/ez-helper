# Build for all platforms
build: build-windows build-linux build-darwin3
	@echo "All platform builds completed successfully!"

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	@if not exist .\build\windows (mkdir .\build\windows)
	@go build -o .\build\windows\ez-helper.exe .\cmd\main.go
	@echo "Windows build successful!"

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	@if not exist .\build\linux (mkdir .\build\linux)
	@SET CGO_ENABLED=0
	@SET GOOS=linux
	@SET GOARCH=amd64
	@go build -o .\build\linux\ez-helper .\cmd\main.go
	@echo "Linux build successful!"

# Build for Mac (darwin3)
build-darwin3:
	@echo "Building for Mac..."
	@if not exist .\build\darwin3 (mkdir .\build\darwin3)
	@SET CGO_ENABLED=0
	@SET GOOS=darwin
	@SET GOARCH=amd64
	@go build -o .\build\darwin3\ez-helper .\cmd\main.go
	@echo "Mac build successful!"
