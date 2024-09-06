package main

// LizardAnalyzer - Extensible Cyclomatic Complexity Analyzer
// Homepage: http://www.lizard.ws

import (
	"fmt"
	
	"os/exec"
)

func installLizardAnalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	lizardanalyzer_tar_url := "https://files.pythonhosted.org/packages/ef/70/bbb7c6b5d1b29acca0cd13582a7303fc528e6dbf40d0026861f9aa7f3ff0/lizard-1.17.10.tar.gz"
	lizardanalyzer_cmd_tar := exec.Command("curl", "-L", lizardanalyzer_tar_url, "-o", "package.tar.gz")
	err := lizardanalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lizardanalyzer_zip_url := "https://files.pythonhosted.org/packages/ef/70/bbb7c6b5d1b29acca0cd13582a7303fc528e6dbf40d0026861f9aa7f3ff0/lizard-1.17.10.zip"
	lizardanalyzer_cmd_zip := exec.Command("curl", "-L", lizardanalyzer_zip_url, "-o", "package.zip")
	err = lizardanalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lizardanalyzer_bin_url := "https://files.pythonhosted.org/packages/ef/70/bbb7c6b5d1b29acca0cd13582a7303fc528e6dbf40d0026861f9aa7f3ff0/lizard-1.17.10.bin"
	lizardanalyzer_cmd_bin := exec.Command("curl", "-L", lizardanalyzer_bin_url, "-o", "binary.bin")
	err = lizardanalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lizardanalyzer_src_url := "https://files.pythonhosted.org/packages/ef/70/bbb7c6b5d1b29acca0cd13582a7303fc528e6dbf40d0026861f9aa7f3ff0/lizard-1.17.10.src.tar.gz"
	lizardanalyzer_cmd_src := exec.Command("curl", "-L", lizardanalyzer_src_url, "-o", "source.tar.gz")
	err = lizardanalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lizardanalyzer_cmd_direct := exec.Command("./binary")
	err = lizardanalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
