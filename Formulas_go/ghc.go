package main

// Ghc - Glorious Glasgow Haskell Compilation System
// Homepage: https://haskell.org/ghc/

import (
	"fmt"
	
	"os/exec"
)

func installGhc() {
	// Método 1: Descargar y extraer .tar.gz
	ghc_tar_url := "https://downloads.haskell.org/~ghc/9.10.1/ghc-9.10.1-src.tar.xz"
	ghc_cmd_tar := exec.Command("curl", "-L", ghc_tar_url, "-o", "package.tar.gz")
	err := ghc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghc_zip_url := "https://downloads.haskell.org/~ghc/9.10.1/ghc-9.10.1-src.tar.xz"
	ghc_cmd_zip := exec.Command("curl", "-L", ghc_zip_url, "-o", "package.zip")
	err = ghc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghc_bin_url := "https://downloads.haskell.org/~ghc/9.10.1/ghc-9.10.1-src.tar.xz"
	ghc_cmd_bin := exec.Command("curl", "-L", ghc_bin_url, "-o", "binary.bin")
	err = ghc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghc_src_url := "https://downloads.haskell.org/~ghc/9.10.1/ghc-9.10.1-src.tar.xz"
	ghc_cmd_src := exec.Command("curl", "-L", ghc_src_url, "-o", "source.tar.gz")
	err = ghc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghc_cmd_direct := exec.Command("./binary")
	err = ghc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
