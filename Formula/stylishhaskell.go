package main

// StylishHaskell - Haskell code prettifier
// Homepage: https://github.com/haskell/stylish-haskell

import (
	"fmt"
	
	"os/exec"
)

func installStylishHaskell() {
	// Método 1: Descargar y extraer .tar.gz
	stylishhaskell_tar_url := "https://github.com/haskell/stylish-haskell/archive/refs/tags/v0.14.6.0.tar.gz"
	stylishhaskell_cmd_tar := exec.Command("curl", "-L", stylishhaskell_tar_url, "-o", "package.tar.gz")
	err := stylishhaskell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stylishhaskell_zip_url := "https://github.com/haskell/stylish-haskell/archive/refs/tags/v0.14.6.0.zip"
	stylishhaskell_cmd_zip := exec.Command("curl", "-L", stylishhaskell_zip_url, "-o", "package.zip")
	err = stylishhaskell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stylishhaskell_bin_url := "https://github.com/haskell/stylish-haskell/archive/refs/tags/v0.14.6.0.bin"
	stylishhaskell_cmd_bin := exec.Command("curl", "-L", stylishhaskell_bin_url, "-o", "binary.bin")
	err = stylishhaskell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stylishhaskell_src_url := "https://github.com/haskell/stylish-haskell/archive/refs/tags/v0.14.6.0.src.tar.gz"
	stylishhaskell_cmd_src := exec.Command("curl", "-L", stylishhaskell_src_url, "-o", "source.tar.gz")
	err = stylishhaskell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stylishhaskell_cmd_direct := exec.Command("./binary")
	err = stylishhaskell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
}
