package main

// HopenpgpTools - Command-line tools for OpenPGP-related operations
// Homepage: https://hackage.haskell.org/package/hopenpgp-tools

import (
	"fmt"
	
	"os/exec"
)

func installHopenpgpTools() {
	// Método 1: Descargar y extraer .tar.gz
	hopenpgptools_tar_url := "https://hackage.haskell.org/package/hopenpgp-tools-0.23.8/hopenpgp-tools-0.23.8.tar.gz"
	hopenpgptools_cmd_tar := exec.Command("curl", "-L", hopenpgptools_tar_url, "-o", "package.tar.gz")
	err := hopenpgptools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hopenpgptools_zip_url := "https://hackage.haskell.org/package/hopenpgp-tools-0.23.8/hopenpgp-tools-0.23.8.zip"
	hopenpgptools_cmd_zip := exec.Command("curl", "-L", hopenpgptools_zip_url, "-o", "package.zip")
	err = hopenpgptools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hopenpgptools_bin_url := "https://hackage.haskell.org/package/hopenpgp-tools-0.23.8/hopenpgp-tools-0.23.8.bin"
	hopenpgptools_cmd_bin := exec.Command("curl", "-L", hopenpgptools_bin_url, "-o", "binary.bin")
	err = hopenpgptools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hopenpgptools_src_url := "https://hackage.haskell.org/package/hopenpgp-tools-0.23.8/hopenpgp-tools-0.23.8.src.tar.gz"
	hopenpgptools_cmd_src := exec.Command("curl", "-L", hopenpgptools_src_url, "-o", "source.tar.gz")
	err = hopenpgptools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hopenpgptools_cmd_direct := exec.Command("./binary")
	err = hopenpgptools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.4")
exec.Command("latte", "install", "ghc@9.4")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
}
