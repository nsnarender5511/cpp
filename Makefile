.PHONY: build test release-test clean release

# Build the binary
build:
	go build -o crules cmd/main.go

# Run tests
test:
	go test ./...

# Test GoReleaser configuration
release-test:
	goreleaser check
	goreleaser release --snapshot --clean --skip=publish

# Clean build artifacts
clean:
	rm -f crules
	rm -rf dist/

# Prepare and publish a release
release:
	@echo "🔍 Starting release process..."
	
	@# Check if tag is provided
	@if [ -z "$(TAG)" ]; then \
		echo "❌ Error: TAG variable is required. Use 'make release TAG=v0.0.1'"; \
		exit 1; \
	fi
	
	@# Check for clean git state
	@echo "📋 Checking git status..."
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "❌ Error: Git working directory is not clean."; \
		echo "   Please commit or stash your changes before releasing."; \
		git status --short; \
		exit 1; \
	else \
		echo "✅ Git working directory is clean."; \
	fi
	
	@# Check if dist/ is in .gitignore
	@echo "📋 Checking .gitignore configuration..."
	@if ! grep -q "^dist/$$" .gitignore; then \
		echo "⚠️  Warning: 'dist/' not found in .gitignore. Adding it..."; \
		echo "dist/" >> .gitignore; \
		git add .gitignore; \
		git commit -m "chore: add dist/ to .gitignore"; \
		echo "✅ Added dist/ to .gitignore and committed."; \
	else \
		echo "✅ dist/ is properly configured in .gitignore."; \
	fi
	
	@# Clean the dist directory
	@echo "🧹 Cleaning distribution directory..."
	@rm -rf dist/
	@echo "✅ Distribution directory cleaned."
	
	@# Run tests
	@echo "🧪 Running tests..."
	@if ! go test ./...; then \
		echo "❌ Error: Tests failed. Fix test issues before releasing."; \
		exit 1; \
	else \
		echo "✅ All tests passed."; \
	fi
	
	@# Check if tag exists
	@echo "🏷️  Checking tag $(TAG)..."
	@if git tag | grep -q "^$(TAG)$$"; then \
		echo "⚠️  Warning: Tag $(TAG) already exists."; \
		read -p "Do you want to delete and recreate it? [y/N] " answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			echo "🔄 Deleting existing tag $(TAG)..."; \
			git tag -d $(TAG); \
			if git ls-remote --tags origin | grep -q "$(TAG)"; then \
				echo "🔄 Deleting remote tag $(TAG)..."; \
				git push --delete origin $(TAG); \
			fi; \
			echo "🔄 Creating new tag $(TAG)..."; \
			if ! git tag -a $(TAG) -m "Release $(TAG)"; then \
				echo "❌ Error: Failed to create tag $(TAG)"; \
				exit 1; \
			fi; \
			echo "✅ Tag $(TAG) created."; \
		else \
			echo "ℹ️  Using existing tag $(TAG)."; \
		fi; \
	else \
		echo "🔄 Creating new tag $(TAG)..."; \
		if ! git tag -a $(TAG) -m "Release $(TAG)"; then \
			echo "❌ Error: Failed to create tag $(TAG)"; \
			exit 1; \
		fi; \
		echo "✅ Tag $(TAG) created."; \
	fi
	
	@# Run GoReleaser
	@echo "🚀 Running GoReleaser..."
	@if [ -f .env ] && [ -z "$$GITHUB_TOKEN" ]; then \
		echo "ℹ️  Using GITHUB_TOKEN from .env file"; \
		TOKEN=$$(grep -E "^GITHUB_TOKEN=" .env | cut -d= -f2); \
		if [ -z "$$TOKEN" ]; then \
			echo "❌ Error: Could not extract GITHUB_TOKEN from .env file"; \
			exit 1; \
		fi; \
		if ! GITHUB_TOKEN="$$TOKEN" goreleaser release --clean; then \
			echo "❌ Error: GoReleaser failed. Check the output above for details."; \
			exit 1; \
		fi; \
	else \
		echo "ℹ️  Using GITHUB_TOKEN from environment"; \
		if ! goreleaser release --clean; then \
			echo "❌ Error: GoReleaser failed. Check the output above for details."; \
			exit 1; \
		fi; \
	fi; \
	echo "✅ Release successful!"
	
	@# Push tag if successful
	@echo "🔄 Pushing tag $(TAG) to remote..."
	@git push origin $(TAG)
	@echo "✅ Tag pushed to remote."
	
	@echo "✨ Release $(TAG) completed successfully!"
	@echo "🌐 Check your GitHub repository for the new release."
	@echo "🍺 Check your Homebrew tap repository to verify the formula was updated."
	@echo ""
	@echo "📦 Users can now install with:"
	@echo "   brew tap nsnarender5511/tap"
	@echo "   brew install crules"

# Show help
help:
	@echo "Available commands:"
	@echo "  build        Build the binary"
	@echo "  test         Run tests"
	@echo "  release-test Test GoReleaser configuration"
	@echo "  clean        Clean build artifacts"
	@echo "  release TAG=v0.0.1  Create and publish a new release with the specified tag"

# Default target
.DEFAULT_GOAL := build 