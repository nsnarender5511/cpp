.PHONY: build test release-test clean release release-github

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

# Trigger a GitHub Actions release by creating and pushing a tag
release-github:
	@echo "🚀 Starting GitHub Actions release process..."
	
	@# Check if tag is provided
	@if [ -z "$(TAG)" ]; then \
		echo "❌ Error: TAG variable is required. Use 'make release-github TAG=v0.0.1'"; \
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
	
	@# Check if tag exists locally
	@echo "🏷️  Checking tag $(TAG)..."
	@if git tag | grep -q "^$(TAG)$$"; then \
		echo "⚠️  Warning: Tag $(TAG) already exists locally."; \
		read -p "Delete and recreate local tag? [y/N] " answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			echo "🔄 Deleting local tag $(TAG)..."; \
			git tag -d $(TAG) || { echo "❌ Failed to delete local tag"; exit 1; }; \
			echo "✅ Local tag deleted successfully."; \
		else \
			echo "❌ Cannot proceed with existing tag. Please choose a different tag or confirm deletion."; \
			exit 1; \
		fi; \
	fi
	
	@# Check if tag exists remotely
	@echo "🔍 Checking if tag exists on remote..."
	@if git ls-remote --tags origin refs/tags/$(TAG) | grep -q ""; then \
		echo "⚠️  Warning: Tag $(TAG) already exists on remote."; \
		read -p "Delete remote tag and release? [y/N] " answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			echo "🔄 Deleting remote tag $(TAG)..."; \
			git push --delete origin $(TAG) || { echo "❌ Failed to delete remote tag"; exit 1; }; \
			echo "✅ Remote tag deleted successfully."; \
			echo "🔄 Attempting to delete GitHub release..."; \
			GITHUB_REPO=$$(git remote get-url origin | sed 's/.*github.com[:/]\(.*\).git/\1/'); \
			if [ -n "$$GITHUB_TOKEN" ]; then \
				echo "🔍 Checking for existing GitHub release..."; \
				RELEASE_ID=$$(curl -s -H "Authorization: token $$GITHUB_TOKEN" \
					"https://api.github.com/repos/$$GITHUB_REPO/releases/tags/$(TAG)" | \
					grep -o '"id": [0-9]*' | head -1 | grep -o '[0-9]*'); \
				if [ -n "$$RELEASE_ID" ]; then \
					echo "🗑️  Deleting GitHub release ID: $$RELEASE_ID"; \
					DELETE_RESULT=$$(curl -s -X DELETE -H "Authorization: token $$GITHUB_TOKEN" \
						"https://api.github.com/repos/$$GITHUB_REPO/releases/$$RELEASE_ID"); \
					echo "✅ GitHub release deleted successfully."; \
				else \
					echo "ℹ️  No existing GitHub release found for tag $(TAG)."; \
				fi; \
			else \
				echo "⚠️  No GITHUB_TOKEN found in environment"; \
				echo "⚠️  Cannot delete GitHub release automatically."; \
				echo "⚠️  Please delete any existing release for tag $(TAG) manually on GitHub."; \
				echo "⚠️  Visit: https://github.com/$$GITHUB_REPO/releases"; \
			fi; \
		else \
			echo "❌ Cannot proceed with existing remote tag. Please choose a different tag."; \
			exit 1; \
		fi; \
	fi
	
	@# Create and push tag
	@echo "🔄 Creating new tag $(TAG)..."
	@git tag -a $(TAG) -m "Release $(TAG)" || { echo "❌ Failed to create tag"; exit 1; }
	@echo "✅ Tag created locally."
	
	@echo "🚀 Pushing tag to trigger GitHub Actions workflow..."
	@git push origin $(TAG) || { echo "❌ Failed to push tag"; exit 1; }
	@echo "✅ Tag pushed successfully."
	@echo "🎉 GitHub Actions release workflow triggered!"
	@echo "   Visit https://github.com/$(shell git remote get-url origin | sed 's/.*github.com[:/]\(.*\).git/\1/')/actions to monitor the release progress."

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
	
	@# Run GoReleaser
	@echo "🚀 Running GoReleaser..."
	@# Debug token information
	@if [ -f .env ] && [ -z "$$GITHUB_TOKEN" ]; then \
		echo "ℹ️  Using GITHUB_TOKEN from .env file"; \
		echo "📄 Contents of .env file (with sensitive data masked):"; \
		cat .env | sed 's/\(GITHUB_TOKEN=\)[^ ]*/\1********/g'; \
		echo "🔍 Extracting token..."; \
		TOKEN=$$(grep -E "^\s*GITHUB_TOKEN=" .env | sed 's/^\s*GITHUB_TOKEN=//'); \
		echo "🔄 Extracted token status: $${TOKEN:+found}$${TOKEN:-not found}"; \
		if [ -z "$$TOKEN" ]; then \
			echo "❌ Error: Could not extract GITHUB_TOKEN from .env file"; \
			echo "   Make sure the line in .env is formatted as GITHUB_TOKEN=your_token"; \
			echo "   without spaces at the beginning of the line."; \
			exit 1; \
		fi; \
		TOKEN_LENGTH=$${#TOKEN}; \
		if [ $$TOKEN_LENGTH -lt 30 ]; then \
			echo "⚠️  Warning: Token looks suspiciously short ($$TOKEN_LENGTH chars)"; \
		fi; \
		echo "📝 Showing first 4 chars of token as a hint..."; \
		TOKEN_PREFIX=$$(echo $$TOKEN | cut -c1-4); \
		echo "🔑 Using token: $$TOKEN_PREFIX**** ($$TOKEN_LENGTH chars)"; \
		echo "🔍 Testing GitHub API access..."; \
		HTTP_CODE=$$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: token $$TOKEN" https://api.github.com/user); \
		if [ "$$HTTP_CODE" = "200" ]; then \
			echo "✅ GitHub authentication successful"; \
		else \
			echo "⚠️  GitHub API returned HTTP $$HTTP_CODE - token may not have correct permissions"; \
			echo "   Token needs 'repo' scope to create releases"; \
		fi; \
		if ! GITHUB_TOKEN="$$TOKEN" goreleaser release --clean; then \
			echo "❌ Error: GoReleaser failed. Check the output above for details."; \
			exit 1; \
		fi; \
	else \
		echo "ℹ️  Using GITHUB_TOKEN from environment"; \
		if [ -n "$$GITHUB_TOKEN" ]; then \
			TOKEN_LENGTH=$${#GITHUB_TOKEN}; \
			if [ $$TOKEN_LENGTH -lt 30 ]; then \
				echo "⚠️  Warning: Token looks suspiciously short ($$TOKEN_LENGTH chars)"; \
			fi; \
			echo "📝 Showing first 4 chars of token as a hint..."; \
			TOKEN_PREFIX=$$(echo $$GITHUB_TOKEN | cut -c1-4); \
			echo "🔑 Using token: $$TOKEN_PREFIX**** ($$TOKEN_LENGTH chars)"; \
			echo "🔍 Testing GitHub API access..."; \
			HTTP_CODE=$$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: token $$GITHUB_TOKEN" https://api.github.com/user); \
			if [ "$$HTTP_CODE" = "200" ]; then \
				echo "✅ GitHub authentication successful"; \
			else \
				echo "⚠️  GitHub API returned HTTP $$HTTP_CODE - token may not have correct permissions"; \
				echo "   Token needs 'repo' scope to create releases"; \
			fi; \
		else \
			echo "❌ Error: GITHUB_TOKEN is empty"; \
			exit 1; \
		fi; \
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
	@echo "  release-github TAG=v0.0.1  Trigger a GitHub Actions release by creating and pushing a tag"

# Default target
.DEFAULT_GOAL := build 