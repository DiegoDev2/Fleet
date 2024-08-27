package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	rubyFilesDir = "/Users/diegodev/LattePkg/Formula/a/formulas" // Ruta donde están los archivos Ruby
	goFilesDir   = "GoStructs"                                   // Directorio donde se guardarán los archivos .go
)

// Estructura base para una fórmula en Go
const goStructTemplate = `package main

import "fmt"

// %[1]sFormula representa una fórmula en Go.
type %[1]sFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg %[1]sFormula) Print() {
	fmt.Printf("Name: %[2]s\\n")
	fmt.Printf("Description: %%s\\n", pkg.Description)
	fmt.Printf("Homepage: %%s\\n", pkg.Homepage)
	fmt.Printf("URL: %%s\\n", pkg.URL)
	fmt.Printf("Sha256: %%s\\n", pkg.Sha256)
	fmt.Printf("Dependencies: %%v\\n", pkg.Dependencies)
}

func main() {
	// Crear una instancia de %[1]sFormula
	pkg := %[1]sFormula{
		Description:  "Descripción de %[2]s",
		Homepage:     "https://example.com",
		URL:          "https://example.com/example-1.0.0.tar.gz",
		Sha256:       "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		Dependencies: []string{"dep1", "dep2"},
	}

	// Imprimir la información de la fórmula
	pkg.Print()
}
`

// Extrae el nombre de la fórmula de un archivo Ruby
func extractFormulaName(content string) string {
	re := regexp.MustCompile(`class\s+(\w+)\s+<\s+Formula`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// Genera un archivo Go con una estructura para la fórmula
func generateGoFile(name string) error {
	structName := name + "Formula"
	content := fmt.Sprintf(goStructTemplate, structName, name)

	// Crear el directorio de destino si no existe
	if err := os.MkdirAll(goFilesDir, os.ModePerm); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s.go", goFilesDir, strings.ToLower(name))
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

func main() {
	files, err := ioutil.ReadDir(rubyFilesDir)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".rb") {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", rubyFilesDir, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
			continue
		}

		name := extractFormulaName(string(content))
		if name != "" {
			fmt.Printf("Generando archivo Go para: %s\n", name)
			if err := generateGoFile(name); err != nil {
				fmt.Printf("Error al generar el archivo Go para %s: %v\n", name, err)
			}
		} else {
			fmt.Printf("No se pudo extraer el nombre de la fórmula de %s\n", file.Name())
		}
	}

	fmt.Println("Generación de archivos Go completada.")
}
