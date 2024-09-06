package main

// PandocCrossref - Pandoc filter for numbering and cross-referencing
// Homepage: https://github.com/lierdakil/pandoc-crossref

import (
	"fmt"
	
	"os/exec"
)

func installPandocCrossref() {
	// Método 1: Descargar y extraer .tar.gz
	pandoccrossref_tar_url := "https://github.com/lierdakil/pandoc-crossref/archive/refs/tags/v0.3.17.1c.tar.gz"
	pandoccrossref_cmd_tar := exec.Command("curl", "-L", pandoccrossref_tar_url, "-o", "package.tar.gz")
	err := pandoccrossref_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandoccrossref_zip_url := "https://github.com/lierdakil/pandoc-crossref/archive/refs/tags/v0.3.17.1c.zip"
	pandoccrossref_cmd_zip := exec.Command("curl", "-L", pandoccrossref_zip_url, "-o", "package.zip")
	err = pandoccrossref_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandoccrossref_bin_url := "https://github.com/lierdakil/pandoc-crossref/archive/refs/tags/v0.3.17.1c.bin"
	pandoccrossref_cmd_bin := exec.Command("curl", "-L", pandoccrossref_bin_url, "-o", "binary.bin")
	err = pandoccrossref_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandoccrossref_src_url := "https://github.com/lierdakil/pandoc-crossref/archive/refs/tags/v0.3.17.1c.src.tar.gz"
	pandoccrossref_cmd_src := exec.Command("curl", "-L", pandoccrossref_src_url, "-o", "source.tar.gz")
	err = pandoccrossref_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandoccrossref_cmd_direct := exec.Command("./binary")
	err = pandoccrossref_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
}
