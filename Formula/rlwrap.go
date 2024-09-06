package main

// Rlwrap - Readline wrapper: adds readline support to tools that lack it
// Homepage: https://github.com/hanslub42/rlwrap

import (
	"fmt"
	
	"os/exec"
)

func installRlwrap() {
	// Método 1: Descargar y extraer .tar.gz
	rlwrap_tar_url := "https://github.com/hanslub42/rlwrap/archive/refs/tags/0.46.1.tar.gz"
	rlwrap_cmd_tar := exec.Command("curl", "-L", rlwrap_tar_url, "-o", "package.tar.gz")
	err := rlwrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rlwrap_zip_url := "https://github.com/hanslub42/rlwrap/archive/refs/tags/0.46.1.zip"
	rlwrap_cmd_zip := exec.Command("curl", "-L", rlwrap_zip_url, "-o", "package.zip")
	err = rlwrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rlwrap_bin_url := "https://github.com/hanslub42/rlwrap/archive/refs/tags/0.46.1.bin"
	rlwrap_cmd_bin := exec.Command("curl", "-L", rlwrap_bin_url, "-o", "binary.bin")
	err = rlwrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rlwrap_src_url := "https://github.com/hanslub42/rlwrap/archive/refs/tags/0.46.1.src.tar.gz"
	rlwrap_cmd_src := exec.Command("curl", "-L", rlwrap_src_url, "-o", "source.tar.gz")
	err = rlwrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rlwrap_cmd_direct := exec.Command("./binary")
	err = rlwrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
