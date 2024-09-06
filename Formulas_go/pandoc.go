package main

// Pandoc - Swiss-army knife of markup format conversion
// Homepage: https://pandoc.org/

import (
	"fmt"
	
	"os/exec"
)

func installPandoc() {
	// Método 1: Descargar y extraer .tar.gz
	pandoc_tar_url := "https://github.com/jgm/pandoc/archive/refs/tags/3.3.tar.gz"
	pandoc_cmd_tar := exec.Command("curl", "-L", pandoc_tar_url, "-o", "package.tar.gz")
	err := pandoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandoc_zip_url := "https://github.com/jgm/pandoc/archive/refs/tags/3.3.zip"
	pandoc_cmd_zip := exec.Command("curl", "-L", pandoc_zip_url, "-o", "package.zip")
	err = pandoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandoc_bin_url := "https://github.com/jgm/pandoc/archive/refs/tags/3.3.bin"
	pandoc_cmd_bin := exec.Command("curl", "-L", pandoc_bin_url, "-o", "binary.bin")
	err = pandoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandoc_src_url := "https://github.com/jgm/pandoc/archive/refs/tags/3.3.src.tar.gz"
	pandoc_cmd_src := exec.Command("curl", "-L", pandoc_src_url, "-o", "source.tar.gz")
	err = pandoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandoc_cmd_direct := exec.Command("./binary")
	err = pandoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.8")
exec.Command("latte", "install", "ghc@9.8")
}
