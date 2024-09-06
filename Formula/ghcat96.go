package main

// GhcAT96 - Glorious Glasgow Haskell Compilation System
// Homepage: https://haskell.org/ghc/

import (
	"fmt"
	
	"os/exec"
)

func installGhcAT96() {
	// Método 1: Descargar y extraer .tar.gz
	ghcat96_tar_url := "https://downloads.haskell.org/~ghc/9.6.6/ghc-9.6.6-src.tar.xz"
	ghcat96_cmd_tar := exec.Command("curl", "-L", ghcat96_tar_url, "-o", "package.tar.gz")
	err := ghcat96_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghcat96_zip_url := "https://downloads.haskell.org/~ghc/9.6.6/ghc-9.6.6-src.tar.xz"
	ghcat96_cmd_zip := exec.Command("curl", "-L", ghcat96_zip_url, "-o", "package.zip")
	err = ghcat96_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghcat96_bin_url := "https://downloads.haskell.org/~ghc/9.6.6/ghc-9.6.6-src.tar.xz"
	ghcat96_cmd_bin := exec.Command("curl", "-L", ghcat96_bin_url, "-o", "binary.bin")
	err = ghcat96_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghcat96_src_url := "https://downloads.haskell.org/~ghc/9.6.6/ghc-9.6.6-src.tar.xz"
	ghcat96_cmd_src := exec.Command("curl", "-L", ghcat96_src_url, "-o", "source.tar.gz")
	err = ghcat96_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghcat96_cmd_direct := exec.Command("./binary")
	err = ghcat96_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
