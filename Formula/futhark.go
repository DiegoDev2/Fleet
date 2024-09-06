package main

// Futhark - Data-parallel functional programming language
// Homepage: https://futhark-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installFuthark() {
	// Método 1: Descargar y extraer .tar.gz
	futhark_tar_url := "https://github.com/diku-dk/futhark/archive/refs/tags/v0.25.21.tar.gz"
	futhark_cmd_tar := exec.Command("curl", "-L", futhark_tar_url, "-o", "package.tar.gz")
	err := futhark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	futhark_zip_url := "https://github.com/diku-dk/futhark/archive/refs/tags/v0.25.21.zip"
	futhark_cmd_zip := exec.Command("curl", "-L", futhark_zip_url, "-o", "package.zip")
	err = futhark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	futhark_bin_url := "https://github.com/diku-dk/futhark/archive/refs/tags/v0.25.21.bin"
	futhark_cmd_bin := exec.Command("curl", "-L", futhark_bin_url, "-o", "binary.bin")
	err = futhark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	futhark_src_url := "https://github.com/diku-dk/futhark/archive/refs/tags/v0.25.21.src.tar.gz"
	futhark_cmd_src := exec.Command("curl", "-L", futhark_src_url, "-o", "source.tar.gz")
	err = futhark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	futhark_cmd_direct := exec.Command("./binary")
	err = futhark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
}
