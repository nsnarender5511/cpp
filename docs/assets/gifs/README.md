# GIF Assets Reference

This directory contains animated GIF assets used in the cursor++ documentation. This README serves as a reference guide to help maintain consistent naming and organization of GIF assets.

## Directory Structure

```
gifs/
├── installation/  # Installation process animations
├── usage/         # General usage animations
└── workflows/     # Workflow demonstrations
```

## Naming Conventions

1. Use lowercase file names
2. Use hyphens to separate words
3. Include a descriptor of the content
4. Use `.gif` extension

Examples:
- `agent-selection-process.gif`
- `installing-on-macos.gif`
- `sync-workflow.gif`

## Required GIFs

The following GIFs are required for the documentation:

### Installation

- `installing-on-macos.gif` - Installing cursor++ on macOS
- `installing-on-linux.gif` - Installing cursor++ on Linux
- `installing-on-windows.gif` - Installing cursor++ on Windows

### Usage

- `initializing-project.gif` - Initializing a project with cursor++
- `selecting-agent.gif` - Selecting an agent to use
- `syncing-rules.gif` - Synchronizing rules across projects

### Workflows

- `feature-development-workflow.gif` - Complete feature development workflow
- `bug-fixing-workflow.gif` - Bug fixing workflow
- `rule-creation-workflow.gif` - Creating and sharing a custom rule

## GIF Specifications

- **Resolution**: 1280x720 (minimum)
- **Frame Rate**: 15-24 fps
- **Max Duration**: 15-30 seconds
- **Max Size**: 2MB
- **Compression**: Optimize for web viewing

## Creating High-Quality GIFs

For consistent, high-quality GIFs:

1. Record at a higher resolution (1920x1080) and scale down
2. Use a neutral/standard terminal theme (dark background, light text)
3. Type at a readable pace
4. Highlight important actions with cursor movements
5. Keep GIFs focused on a single task or concept
6. Include a brief pause at beginning and end (1-2 seconds)

## Tools for Creating GIFs

Recommended tools:
- **macOS**: 
  - LICEcap (https://www.cockos.com/licecap/)
  - Kap (https://getkap.co/)
- **Windows**: 
  - ScreenToGif (https://www.screentogif.com/)
  - ShareX (https://getsharex.com/)
- **Linux**: 
  - Peek (https://github.com/phw/peek)
  - Silent Cast (https://github.com/colinkeenan/silentcast)

## Optimization

Before committing GIFs to the repository:
1. Optimize using tools like `gifsicle` or online services like EZGif
2. Reduce colors if possible
3. Consider using compressed video formats (MP4) with GIF fallback for documentation website

## Adding New GIFs

When adding new GIFs:

1. Follow the naming conventions
2. Place the GIF in the appropriate subdirectory
3. Optimize the GIF for web
4. Update this README if creating a new category

## GIF Refresh Guidelines

GIFs should be refreshed:

- When the UI changes significantly
- At least once per major version
- When features shown in the GIF are modified
- When command output format changes

## Usage Guidelines

When using GIFs in documentation:

- Reference GIFs with relative paths
- Include descriptive alt text for accessibility
- Place GIFs strategically to demonstrate complex interactions
- Supplement GIFs with text descriptions for search indexing
- Consider accessibility by providing text alternatives

---

Last updated: 2023-10-30 