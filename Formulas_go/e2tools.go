package main

// E2tools - Utilities to read, write, and manipulate files in ext2/3/4 filesystems
// Homepage: https://e2tools.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installE2tools() {
	// Método 1: Descargar y extraer .tar.gz
	e2tools_tar_url := "https://github.com/e2tools/e2tools/releases/download/v0.1.0/e2tools-0.1.0.tar.gz"
	e2tools_cmd_tar := exec.Command("curl", "-L", e2tools_tar_url, "-o", "package.tar.gz")
	err := e2tools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	e2tools_zip_url := "https://github.com/e2tools/e2tools/releases/download/v0.1.0/e2tools-0.1.0.zip"
	e2tools_cmd_zip := exec.Command("curl", "-L", e2tools_zip_url, "-o", "package.zip")
	err = e2tools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	e2tools_bin_url := "https://github.com/e2tools/e2tools/releases/download/v0.1.0/e2tools-0.1.0.bin"
	e2tools_cmd_bin := exec.Command("curl", "-L", e2tools_bin_url, "-o", "binary.bin")
	err = e2tools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	e2tools_src_url := "https://github.com/e2tools/e2tools/releases/download/v0.1.0/e2tools-0.1.0.src.tar.gz"
	e2tools_cmd_src := exec.Command("curl", "-L", e2tools_src_url, "-o", "source.tar.gz")
	err = e2tools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	e2tools_cmd_direct := exec.Command("./binary")
	err = e2tools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
}
