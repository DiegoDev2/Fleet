package main

// Exa - Modern replacement for 'ls'
// Homepage: https://the.exa.website

import (
	"fmt"
	
	"os/exec"
)

func installExa() {
	// Método 1: Descargar y extraer .tar.gz
	exa_tar_url := "https://github.com/ogham/exa/archive/refs/tags/v0.10.1.tar.gz"
	exa_cmd_tar := exec.Command("curl", "-L", exa_tar_url, "-o", "package.tar.gz")
	err := exa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exa_zip_url := "https://github.com/ogham/exa/archive/refs/tags/v0.10.1.zip"
	exa_cmd_zip := exec.Command("curl", "-L", exa_zip_url, "-o", "package.zip")
	err = exa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exa_bin_url := "https://github.com/ogham/exa/archive/refs/tags/v0.10.1.bin"
	exa_cmd_bin := exec.Command("curl", "-L", exa_bin_url, "-o", "binary.bin")
	err = exa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exa_src_url := "https://github.com/ogham/exa/archive/refs/tags/v0.10.1.src.tar.gz"
	exa_cmd_src := exec.Command("curl", "-L", exa_src_url, "-o", "source.tar.gz")
	err = exa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exa_cmd_direct := exec.Command("./binary")
	err = exa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
}
