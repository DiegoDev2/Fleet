package main

// PandocIncludeCode - Pandoc filter for including code from source files
// Homepage: https://github.com/owickstrom/pandoc-include-code

import (
	"fmt"
	
	"os/exec"
)

func installPandocIncludeCode() {
	// Método 1: Descargar y extraer .tar.gz
	pandocincludecode_tar_url := "https://hackage.haskell.org/package/pandoc-include-code-1.5.0.0/pandoc-include-code-1.5.0.0.tar.gz"
	pandocincludecode_cmd_tar := exec.Command("curl", "-L", pandocincludecode_tar_url, "-o", "package.tar.gz")
	err := pandocincludecode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandocincludecode_zip_url := "https://hackage.haskell.org/package/pandoc-include-code-1.5.0.0/pandoc-include-code-1.5.0.0.zip"
	pandocincludecode_cmd_zip := exec.Command("curl", "-L", pandocincludecode_zip_url, "-o", "package.zip")
	err = pandocincludecode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandocincludecode_bin_url := "https://hackage.haskell.org/package/pandoc-include-code-1.5.0.0/pandoc-include-code-1.5.0.0.bin"
	pandocincludecode_cmd_bin := exec.Command("curl", "-L", pandocincludecode_bin_url, "-o", "binary.bin")
	err = pandocincludecode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandocincludecode_src_url := "https://hackage.haskell.org/package/pandoc-include-code-1.5.0.0/pandoc-include-code-1.5.0.0.src.tar.gz"
	pandocincludecode_cmd_src := exec.Command("curl", "-L", pandocincludecode_src_url, "-o", "source.tar.gz")
	err = pandocincludecode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandocincludecode_cmd_direct := exec.Command("./binary")
	err = pandocincludecode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@8.10")
exec.Command("latte", "install", "ghc@8.10")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
}
