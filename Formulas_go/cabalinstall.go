package main

// CabalInstall - Command-line interface for Cabal and Hackage
// Homepage: https://www.haskell.org/cabal/

import (
	"fmt"
	
	"os/exec"
)

func installCabalInstall() {
	// Método 1: Descargar y extraer .tar.gz
	cabalinstall_tar_url := "https://hackage.haskell.org/package/cabal-install-3.12.1.0/cabal-install-3.12.1.0.tar.gz"
	cabalinstall_cmd_tar := exec.Command("curl", "-L", cabalinstall_tar_url, "-o", "package.tar.gz")
	err := cabalinstall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cabalinstall_zip_url := "https://hackage.haskell.org/package/cabal-install-3.12.1.0/cabal-install-3.12.1.0.zip"
	cabalinstall_cmd_zip := exec.Command("curl", "-L", cabalinstall_zip_url, "-o", "package.zip")
	err = cabalinstall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cabalinstall_bin_url := "https://hackage.haskell.org/package/cabal-install-3.12.1.0/cabal-install-3.12.1.0.bin"
	cabalinstall_cmd_bin := exec.Command("curl", "-L", cabalinstall_bin_url, "-o", "binary.bin")
	err = cabalinstall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cabalinstall_src_url := "https://hackage.haskell.org/package/cabal-install-3.12.1.0/cabal-install-3.12.1.0.src.tar.gz"
	cabalinstall_cmd_src := exec.Command("curl", "-L", cabalinstall_src_url, "-o", "source.tar.gz")
	err = cabalinstall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cabalinstall_cmd_direct := exec.Command("./binary")
	err = cabalinstall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
}
