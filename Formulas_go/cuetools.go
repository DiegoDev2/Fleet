package main

// Cuetools - Utilities for .cue and .toc files
// Homepage: https://github.com/svend/cuetools

import (
	"fmt"
	
	"os/exec"
)

func installCuetools() {
	// Método 1: Descargar y extraer .tar.gz
	cuetools_tar_url := "https://github.com/svend/cuetools/archive/refs/tags/1.4.1.tar.gz"
	cuetools_cmd_tar := exec.Command("curl", "-L", cuetools_tar_url, "-o", "package.tar.gz")
	err := cuetools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cuetools_zip_url := "https://github.com/svend/cuetools/archive/refs/tags/1.4.1.zip"
	cuetools_cmd_zip := exec.Command("curl", "-L", cuetools_zip_url, "-o", "package.zip")
	err = cuetools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cuetools_bin_url := "https://github.com/svend/cuetools/archive/refs/tags/1.4.1.bin"
	cuetools_cmd_bin := exec.Command("curl", "-L", cuetools_bin_url, "-o", "binary.bin")
	err = cuetools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cuetools_src_url := "https://github.com/svend/cuetools/archive/refs/tags/1.4.1.src.tar.gz"
	cuetools_cmd_src := exec.Command("curl", "-L", cuetools_src_url, "-o", "source.tar.gz")
	err = cuetools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cuetools_cmd_direct := exec.Command("./binary")
	err = cuetools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
