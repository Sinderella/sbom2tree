package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

func main() {
	// Define the -s flag for searching
	searchTerm := flag.String("s", "", "Search term for filtering dependencies")
	flag.Parse()

	// Ensure the user provides a file path as an argument
	if flag.NArg() < 1 {
		fmt.Printf("Usage: %s [-s search-term] <path-to-bom-file>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := flag.Arg(0)

	// Open the specified file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Initialize a new BOM (Bill of Materials) object
	bom := new(cdx.BOM)

	// Determine the file format and create the appropriate decoder
	var decoder cdx.BOMDecoder
	switch {
	case strings.HasSuffix(filePath, ".json"):
		decoder = cdx.NewBOMDecoder(file, cdx.BOMFileFormatJSON)
	case strings.HasSuffix(filePath, ".xml"):
		decoder = cdx.NewBOMDecoder(file, cdx.BOMFileFormatXML)
	default:
		log.Fatalf("Unsupported file format. Please provide a JSON or XML file.")
	}

	// Decode the BOM file into the BOM object
	if err = decoder.Decode(bom); err != nil {
		log.Fatalf("Failed to decode BOM: %v", err)
	}

	printDependencies(bom, *searchTerm)
}

func printDependencies(bom *cdx.BOM, searchTerm string) {
	// Print the dependencies in a tree-like structure
	for _, dependency := range *bom.Dependencies {
		shouldPrint := searchTerm == "" || strings.Contains(dependency.Ref, searchTerm)
		if !shouldPrint && dependency.Dependencies != nil {
			for _, subDependency := range *dependency.Dependencies {
				if strings.Contains(subDependency, searchTerm) {
					shouldPrint = true
					break
				}
			}
		}
		if shouldPrint {
			fmt.Printf("|-- %s\n", dependency.Ref)
			if dependency.Dependencies != nil {
				for _, subDependency := range *dependency.Dependencies {
					if searchTerm == "" || strings.Contains(subDependency, searchTerm) {
						fmt.Printf("|   |-- %s\n", subDependency)
					}
				}
			}
		}
	}
}
