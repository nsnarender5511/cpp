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
			git tag -a $(TAG) -m "Release $(TAG)"; \
			echo "✅ Tag $(TAG) created."; \
		else \
			echo "ℹ️  Using existing tag $(TAG)."; \
		fi; \
	else \
		echo "🔄 Creating new tag $(TAG)..."; \
		git tag -a $(TAG) -m "Release $(TAG)"; \
		echo "✅ Tag $(TAG) created."; \
	fi
	
	@# Check for GITHUB_TOKEN from environment or .env file
	@echo "🔑 Checking for GITHUB_TOKEN..."
	@if [ -z "$$GITHUB_TOKEN" ]; then \
		echo "ℹ️  GITHUB_TOKEN not found in environment, checking .env file..."; \
		if [ -f .env ]; then \
			echo "📄 Found .env file, trying to load GITHUB_TOKEN..."; \
			export GITHUB_TOKEN=$$(grep -E "^GITHUB_TOKEN=" .env | cut -d= -f2); \
			if [ -z "$$GITHUB_TOKEN" ]; then \
				echo "❌ Error: GITHUB_TOKEN not found in .env file."; \
				echo "   Add GITHUB_TOKEN=your_token to your .env file or export it."; \
				exit 1; \
			else \
				echo "✅ GITHUB_TOKEN loaded from .env file."; \
			fi; \
		else \
			echo "❌ Error: GITHUB_TOKEN environment variable is not set and .env file not found."; \
			echo "   Either:"; \
			echo "   - Export your GitHub token with: export GITHUB_TOKEN=your_token"; \
			echo "   - Create a .env file with: GITHUB_TOKEN=your_token"; \
			exit 1; \
		fi; \
	else \
		echo "✅ GITHUB_TOKEN is set in environment."; \
	fi
	
	@# Run GoReleaser
	@echo "🚀 Running GoReleaser..."
	@if ! ( \
		if [ -f .env ] && [ -z "$$GITHUB_TOKEN" ]; then \
			GITHUB_TOKEN=$$(grep -E "^GITHUB_TOKEN=" .env | cut -d= -f2) goreleaser release --clean; \
		else \
			goreleaser release --clean; \
		fi \
	); then \
		echo "❌ Error: GoReleaser failed. Check the output above for details."; \
		exit 1; \
	else \
		echo "✅ Release successful!"; \
	fi
	
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