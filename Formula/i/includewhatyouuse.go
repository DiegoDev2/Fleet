package main

import "fmt"

// IncludeWhatYouUseFormulaFormula representa una fórmula en Go.
type IncludeWhatYouUseFormulaFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg IncludeWhatYouUseFormulaFormula) Print() {
	fmt.Printf("Name: IncludeWhatYouUse\\n")
	fmt.Printf("Description: %s\\n", pkg.Description)
	fmt.Printf("Homepage: %s\\n", pkg.Homepage)
	fmt.Printf("URL: %s\\n", pkg.URL)
	fmt.Printf("Sha256: %s\\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\\n", pkg.Dependencies)
}

func main() {
	// Crear una instancia de IncludeWhatYouUseFormulaFormula
	pkg := IncludeWhatYouUseFormulaFormula{
		Description:  "Descripción de IncludeWhatYouUse",
		Homepage:     "https://example.com",
		URL:          "https://example.com/example-1.0.0.tar.gz",
		Sha256:       "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		Dependencies: []string{"dep1", "dep2"},
	}

	// Imprimir la información de la fórmula
	pkg.Print()
}
