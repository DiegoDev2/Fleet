package main

// DhallLspServer - Language Server Protocol (LSP) server for Dhall
// Homepage: https://github.com/dhall-lang/dhall-haskell/tree/master/dhall-lsp-server

import (
	"fmt"
	
	"os/exec"
)

func installDhallLspServer() {
	// Método 1: Descargar y extraer .tar.gz
	dhalllspserver_tar_url := "https://hackage.haskell.org/package/dhall-lsp-server-1.1.3/dhall-lsp-server-1.1.3.tar.gz"
	dhalllspserver_cmd_tar := exec.Command("curl", "-L", dhalllspserver_tar_url, "-o", "package.tar.gz")
	err := dhalllspserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhalllspserver_zip_url := "https://hackage.haskell.org/package/dhall-lsp-server-1.1.3/dhall-lsp-server-1.1.3.zip"
	dhalllspserver_cmd_zip := exec.Command("curl", "-L", dhalllspserver_zip_url, "-o", "package.zip")
	err = dhalllspserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhalllspserver_bin_url := "https://hackage.haskell.org/package/dhall-lsp-server-1.1.3/dhall-lsp-server-1.1.3.bin"
	dhalllspserver_cmd_bin := exec.Command("curl", "-L", dhalllspserver_bin_url, "-o", "binary.bin")
	err = dhalllspserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhalllspserver_src_url := "https://hackage.haskell.org/package/dhall-lsp-server-1.1.3/dhall-lsp-server-1.1.3.src.tar.gz"
	dhalllspserver_cmd_src := exec.Command("curl", "-L", dhalllspserver_src_url, "-o", "source.tar.gz")
	err = dhalllspserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhalllspserver_cmd_direct := exec.Command("./binary")
	err = dhalllspserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.4")
exec.Command("latte", "install", "ghc@9.4")
}
