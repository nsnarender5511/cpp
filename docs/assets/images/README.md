# Image Assets Reference

This directory contains all image assets used in the crules documentation. This README serves as a reference guide to help maintain consistent naming and organization of image assets.

## Directory Structure

```
images/
├── examples/     # Images used in example documentation
├── logos/        # Project logos and branding
├── screenshots/  # UI screenshots
└── ui/           # UI components and elements
```

## Naming Conventions

1. Use lowercase file names
2. Use hyphens to separate words
3. Include a descriptor of the content
4. Use appropriate file extensions (.png, .jpg, .svg)

Examples:
- `agent-selection-interface.png`
- `crules-logo-dark.svg`
- `sync-command-output.png`

## Required Images

The following images are required for the documentation:

### Logos

- `logo.png` - The main crules logo (color)
- `logo-dark.png` - Dark-themed logo
- `logo-light.png` - Light-themed logo
- `favicon.ico` - Favicon for the documentation site

### UI Screenshots

- `agent-selection-interface.png` - The agent selection UI
- `main-interface.png` - The main crules interface
- `command-line-interface.png` - Example of the CLI interface

### Example Screenshots

- `sync-command-example.png` - Example of the sync command output
- `init-command-example.png` - Example of the init command output
- `agent-command-example.png` - Example of the agent command output

## Image Specifications

- **Screenshots**: 
  - Resolution: 1920x1080 or 2560x1440
  - Format: PNG
  - Max size: 500KB

- **Logos**:
  - Format: SVG (preferred) or PNG
  - Background: Transparent
  - PNG resolution: at least 512x512

- **Icons**:
  - Format: SVG (preferred) or PNG
  - Size: 64x64 or 128x128
  - Background: Transparent

## Adding New Images

When adding new images:

1. Follow the naming conventions
2. Place the image in the appropriate subdirectory
3. Optimize the image for web (compress PNGs, optimize SVGs)
4. Update this README if creating a new category of images

## Image Refresh Guidelines

Screenshots should be refreshed:

- When the UI changes significantly
- At least once per major version
- When features shown in the screenshot are modified

## Usage Guidelines

When using images in documentation:

- Reference images with relative paths
- Include descriptive alt text for accessibility
- Prefer SVG for diagrams and logos
- Use PNG for screenshots
- Keep file sizes small for better page load times

---

Last updated: 2023-10-30 