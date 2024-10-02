# sbom2tree

`sbom2tree` is a command-line tool for parsing and displaying the dependencies from a Software Bill of Materials (SBOM) file in a tree-like structure. It supports both JSON and XML formats.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)

## Installation

### Using Docker

You can build and run the `sbom2tree` tool using Docker:

```sh
docker build -t sbom2tree .
docker run --rm -v $(pwd):/app sbom2tree /app/path-to-bom-file
```

### From Source

To build the tool from source, you need to have Go installed:

```sh
git clone https://github.com/Sinderella/sbom2tree.git
cd sbom2tree
go build ./cmd/sbom2tree
```

## Usage

To use the `sbom2tree` tool, provide the path to the SBOM file as an argument. You can also use the `-s` flag to filter dependencies by a search term:

```sh
./sbom2tree [-s search-term] <path-to-bom-file>
```

### Example

```sh
./sbom2tree -s Json <path-to-bom-file>
```

## Development

### Building

To build the project, you can use the `just` command:

```sh
just build
```
