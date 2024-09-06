package main

// Bnfc - BNF Converter
// Homepage: https://bnfc.digitalgrammars.com/

import (
	"fmt"
	
	"os/exec"
)

func installBnfc() {
	// Método 1: Descargar y extraer .tar.gz
	bnfc_tar_url := "https://github.com/BNFC/bnfc/archive/refs/tags/v2.9.5.tar.gz"
	bnfc_cmd_tar := exec.Command("curl", "-L", bnfc_tar_url, "-o", "package.tar.gz")
	err := bnfc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bnfc_zip_url := "https://github.com/BNFC/bnfc/archive/refs/tags/v2.9.5.zip"
	bnfc_cmd_zip := exec.Command("curl", "-L", bnfc_zip_url, "-o", "package.zip")
	err = bnfc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bnfc_bin_url := "https://github.com/BNFC/bnfc/archive/refs/tags/v2.9.5.bin"
	bnfc_cmd_bin := exec.Command("curl", "-L", bnfc_bin_url, "-o", "binary.bin")
	err = bnfc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bnfc_src_url := "https://github.com/BNFC/bnfc/archive/refs/tags/v2.9.5.src.tar.gz"
	bnfc_cmd_src := exec.Command("curl", "-L", bnfc_src_url, "-o", "source.tar.gz")
	err = bnfc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bnfc_cmd_direct := exec.Command("./binary")
	err = bnfc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: agda")
	exec.Command("latte", "install", "agda").Run()
	fmt.Println("Instalando dependencia: antlr")
	exec.Command("latte", "install", "antlr").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
