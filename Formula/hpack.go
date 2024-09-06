package main

// Hpack - Modern format for Haskell packages
// Homepage: https://github.com/sol/hpack

import (
	"fmt"
	
	"os/exec"
)

func installHpack() {
	// Método 1: Descargar y extraer .tar.gz
	hpack_tar_url := "https://github.com/sol/hpack/archive/refs/tags/0.37.0.tar.gz"
	hpack_cmd_tar := exec.Command("curl", "-L", hpack_tar_url, "-o", "package.tar.gz")
	err := hpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hpack_zip_url := "https://github.com/sol/hpack/archive/refs/tags/0.37.0.zip"
	hpack_cmd_zip := exec.Command("curl", "-L", hpack_zip_url, "-o", "package.zip")
	err = hpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hpack_bin_url := "https://github.com/sol/hpack/archive/refs/tags/0.37.0.bin"
	hpack_cmd_bin := exec.Command("curl", "-L", hpack_bin_url, "-o", "binary.bin")
	err = hpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hpack_src_url := "https://github.com/sol/hpack/archive/refs/tags/0.37.0.src.tar.gz"
	hpack_cmd_src := exec.Command("curl", "-L", hpack_src_url, "-o", "source.tar.gz")
	err = hpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hpack_cmd_direct := exec.Command("./binary")
	err = hpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
}
