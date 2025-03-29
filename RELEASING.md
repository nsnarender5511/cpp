# Release Process

This document describes the process for creating a new release of crules.

## Prerequisites

1. [GoReleaser](https://goreleaser.com/install/) installed locally
2. Git access to the repository with push permissions
3. GitHub personal access token with `repo` scope (for GoReleaser)
4. Create these repository dependencies:
   - https://github.com/nsnarender5511/homebrew-tap (for Homebrew formulas)
   - https://github.com/nsnarender5511/scoop-bucket (for Scoop manifests)

## Release Steps

1. **Ensure all changes are committed and pushed to the main branch**

2. **Update CHANGELOG.md**
   - Move changes from "Unreleased" to a new version section
   - Update the release date
   - Commit the changes:
     ```bash
     git add CHANGELOG.md
     git commit -m "docs: prepare for release vX.Y.Z"
     git push
     ```

3. **Create and push a new tag**
   ```bash
   git tag -a vX.Y.Z -m "Release vX.Y.Z"
   git push origin vX.Y.Z
   ```

4. **Monitor the GitHub Actions workflow**
   - Go to the Actions tab in GitHub
   - Ensure the release workflow completes successfully

5. **Verify the release**
   - Check the GitHub Releases page
   - Verify the Homebrew formula was updated
   - Verify the Scoop manifest was updated
   - Test installation via the different methods

## Testing Before Release

To test the release process locally:

```bash
# Check the GoReleaser config
make release-test
```

This will:
1. Validate the GoReleaser configuration
2. Build a snapshot release locally
3. Skip the actual publishing step

## Troubleshooting

If the release fails:

1. Check the GitHub Actions logs for errors
2. Make necessary corrections
3. Delete the tag locally and remotely:
   ```bash
   git tag -d vX.Y.Z
   git push --delete origin vX.Y.Z
   ```
4. Try again with the fixed setup

## Post-Release

1. Update version references in documentation if needed
2. Create an "Unreleased" section in CHANGELOG.md for future changes
3. Announce the release in appropriate channels 