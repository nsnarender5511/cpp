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
		read -p "Delete and recreate tag? [y/N] " answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			echo "🔄 Deleting local tag $(TAG)..."; \
			git tag -d $(TAG) || { echo "❌ Failed to delete local tag"; exit 1; }; \
			if git ls-remote --tags origin refs/tags/$(TAG) | grep -q ""; then \
				echo "🔄 Deleting remote tag $(TAG)..."; \
				git push --delete origin $(TAG) || { echo "❌ Failed to delete remote tag"; exit 1; }; \
			fi; \
			echo "✅ Existing tags deleted successfully."; \
			echo "🔄 Creating new tag $(TAG)..."; \
			git tag -a $(TAG) -m "Release $(TAG)" || { echo "❌ Failed to create tag"; exit 1; }; \
			echo "✅ New tag created successfully."; \
		else \
			echo "❌ Cannot proceed with existing tag. Please choose a different tag or confirm deletion."; \
			exit 1; \
		fi; \
	else \
		echo "🔄 Creating new tag $(TAG)..."; \
		git tag -a $(TAG) -m "Release $(TAG)" || { echo "❌ Failed to create tag"; exit 1; }; \
		echo "✅ Tag created successfully."; \
	fi
	
	@# Use token from .env file or exit
	@echo "🔑 Loading token EXCLUSIVELY from .env file..."
	@if [ ! -f .env ]; then \
		echo "❌ Error: .env file not found!"; \
		exit 1; \
	fi
	@echo "📄 Contents of .env file (with sensitive data masked):"; \
	@cat .env | sed 's/\(GITHUB_TOKEN=\)[^ ]*/\1********/g'; \
	@echo "🔍 Extracting token..."; \
	@TOKEN=$$(grep -E "^\s*GITHUB_TOKEN=" .env | sed 's/^\s*GITHUB_TOKEN=//' | tr -d '\r\n '); \
	@if [ -z "$$TOKEN" ]; then \
		echo "❌ Error: Could not extract GITHUB_TOKEN from .env file"; \
		echo "   Make sure the line in .env is formatted as GITHUB_TOKEN=your_token"; \
		echo "   without spaces at the beginning or end of the line."; \
		exit 1; \
	fi; \
	echo "🔄 Token successfully extracted from .env"; \
	TOKEN_LENGTH=$${#TOKEN}; \
	if [ $$TOKEN_LENGTH -lt 30 ]; then \
		echo "⚠️  Warning: Token looks suspiciously short ($$TOKEN_LENGTH chars)"; \
	fi; \
	echo "📝 Showing first 8 chars of token as a hint..."; \
	TOKEN_PREFIX=$$(echo $$TOKEN | cut -c1-8); \
	echo "🔑 Using token: $$TOKEN_PREFIX**** ($$TOKEN_LENGTH chars)"; \
	echo "🔍 Testing GitHub API access..."; \
	HTTP_CODE=$$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: token $$TOKEN" https://api.github.com/user); \
	if [ "$$HTTP_CODE" = "200" ]; then \
		echo "✅ GitHub authentication successful"; \
	else \
		echo "⚠️  GitHub API returned HTTP $$HTTP_CODE - token may not have correct permissions"; \
		echo "   Token needs 'repo' scope to create releases"; \
		echo "   You can create a new token at: https://github.com/settings/tokens"; \
	fi; \
	echo "🚀 Running GoReleaser with token from .env file..."; \
	if ! GITHUB_TOKEN="$$TOKEN" goreleaser release --clean; then \
		echo "❌ Error: GoReleaser failed. Check the output above for details."; \
		exit 1; \
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