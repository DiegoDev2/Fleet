package main

// Mighttpd2 - HTTP server
// Homepage: https://kazu-yamamoto.github.io/mighttpd2/

import (
	"fmt"
	
	"os/exec"
)

func installMighttpd2() {
	// Método 1: Descargar y extraer .tar.gz
	mighttpd2_tar_url := "https://hackage.haskell.org/package/mighttpd2-4.0.7/mighttpd2-4.0.7.tar.gz"
	mighttpd2_cmd_tar := exec.Command("curl", "-L", mighttpd2_tar_url, "-o", "package.tar.gz")
	err := mighttpd2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mighttpd2_zip_url := "https://hackage.haskell.org/package/mighttpd2-4.0.7/mighttpd2-4.0.7.zip"
	mighttpd2_cmd_zip := exec.Command("curl", "-L", mighttpd2_zip_url, "-o", "package.zip")
	err = mighttpd2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mighttpd2_bin_url := "https://hackage.haskell.org/package/mighttpd2-4.0.7/mighttpd2-4.0.7.bin"
	mighttpd2_cmd_bin := exec.Command("curl", "-L", mighttpd2_bin_url, "-o", "binary.bin")
	err = mighttpd2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mighttpd2_src_url := "https://hackage.haskell.org/package/mighttpd2-4.0.7/mighttpd2-4.0.7.src.tar.gz"
	mighttpd2_cmd_src := exec.Command("curl", "-L", mighttpd2_src_url, "-o", "source.tar.gz")
	err = mighttpd2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mighttpd2_cmd_direct := exec.Command("./binary")
	err = mighttpd2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
}
