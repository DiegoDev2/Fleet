package main

// GhcAT810 - Glorious Glasgow Haskell Compilation System
// Homepage: https://haskell.org/ghc/

import (
	"fmt"
	
	"os/exec"
)

func installGhcAT810() {
	// Método 1: Descargar y extraer .tar.gz
	ghcat810_tar_url := "https://downloads.haskell.org/~ghc/8.10.7/ghc-8.10.7-src.tar.xz"
	ghcat810_cmd_tar := exec.Command("curl", "-L", ghcat810_tar_url, "-o", "package.tar.gz")
	err := ghcat810_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghcat810_zip_url := "https://downloads.haskell.org/~ghc/8.10.7/ghc-8.10.7-src.tar.xz"
	ghcat810_cmd_zip := exec.Command("curl", "-L", ghcat810_zip_url, "-o", "package.zip")
	err = ghcat810_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghcat810_bin_url := "https://downloads.haskell.org/~ghc/8.10.7/ghc-8.10.7-src.tar.xz"
	ghcat810_cmd_bin := exec.Command("curl", "-L", ghcat810_bin_url, "-o", "binary.bin")
	err = ghcat810_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghcat810_src_url := "https://downloads.haskell.org/~ghc/8.10.7/ghc-8.10.7-src.tar.xz"
	ghcat810_cmd_src := exec.Command("curl", "-L", ghcat810_src_url, "-o", "source.tar.gz")
	err = ghcat810_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghcat810_cmd_direct := exec.Command("./binary")
	err = ghcat810_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.10")
exec.Command("latte", "install", "python@3.10")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: llvm@12")
exec.Command("latte", "install", "llvm@12")
}
