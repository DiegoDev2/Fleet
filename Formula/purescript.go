package main

// Purescript - Strongly typed programming language that compiles to JavaScript
// Homepage: https://www.purescript.org/

import (
	"fmt"
	
	"os/exec"
)

func installPurescript() {
	// Método 1: Descargar y extraer .tar.gz
	purescript_tar_url := "https://hackage.haskell.org/package/purescript-0.15.15/purescript-0.15.15.tar.gz"
	purescript_cmd_tar := exec.Command("curl", "-L", purescript_tar_url, "-o", "package.tar.gz")
	err := purescript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	purescript_zip_url := "https://hackage.haskell.org/package/purescript-0.15.15/purescript-0.15.15.zip"
	purescript_cmd_zip := exec.Command("curl", "-L", purescript_zip_url, "-o", "package.zip")
	err = purescript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	purescript_bin_url := "https://hackage.haskell.org/package/purescript-0.15.15/purescript-0.15.15.bin"
	purescript_cmd_bin := exec.Command("curl", "-L", purescript_bin_url, "-o", "binary.bin")
	err = purescript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	purescript_src_url := "https://hackage.haskell.org/package/purescript-0.15.15/purescript-0.15.15.src.tar.gz"
	purescript_cmd_src := exec.Command("curl", "-L", purescript_src_url, "-o", "source.tar.gz")
	err = purescript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	purescript_cmd_direct := exec.Command("./binary")
	err = purescript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.6")
	exec.Command("latte", "install", "ghc@9.6").Run()
}
