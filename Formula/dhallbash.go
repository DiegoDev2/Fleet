package main

// DhallBash - Compile Dhall to Bash
// Homepage: https://github.com/dhall-lang/dhall-haskell/tree/master/dhall-bash

import (
	"fmt"
	
	"os/exec"
)

func installDhallBash() {
	// Método 1: Descargar y extraer .tar.gz
	dhallbash_tar_url := "https://hackage.haskell.org/package/dhall-bash-1.0.41/dhall-bash-1.0.41.tar.gz"
	dhallbash_cmd_tar := exec.Command("curl", "-L", dhallbash_tar_url, "-o", "package.tar.gz")
	err := dhallbash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhallbash_zip_url := "https://hackage.haskell.org/package/dhall-bash-1.0.41/dhall-bash-1.0.41.zip"
	dhallbash_cmd_zip := exec.Command("curl", "-L", dhallbash_zip_url, "-o", "package.zip")
	err = dhallbash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhallbash_bin_url := "https://hackage.haskell.org/package/dhall-bash-1.0.41/dhall-bash-1.0.41.bin"
	dhallbash_cmd_bin := exec.Command("curl", "-L", dhallbash_bin_url, "-o", "binary.bin")
	err = dhallbash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhallbash_src_url := "https://hackage.haskell.org/package/dhall-bash-1.0.41/dhall-bash-1.0.41.src.tar.gz"
	dhallbash_cmd_src := exec.Command("curl", "-L", dhallbash_src_url, "-o", "source.tar.gz")
	err = dhallbash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhallbash_cmd_direct := exec.Command("./binary")
	err = dhallbash_cmd_direct.Run()
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
