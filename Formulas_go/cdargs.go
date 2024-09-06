package main

// Cdargs - Directory bookmarking system - Enhanced cd utilities
// Homepage: https://github.com/cbxbiker61/cdargs

import (
	"fmt"
	
	"os/exec"
)

func installCdargs() {
	// Método 1: Descargar y extraer .tar.gz
	cdargs_tar_url := "https://github.com/cbxbiker61/cdargs/archive/refs/tags/2.1.tar.gz"
	cdargs_cmd_tar := exec.Command("curl", "-L", cdargs_tar_url, "-o", "package.tar.gz")
	err := cdargs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdargs_zip_url := "https://github.com/cbxbiker61/cdargs/archive/refs/tags/2.1.zip"
	cdargs_cmd_zip := exec.Command("curl", "-L", cdargs_zip_url, "-o", "package.zip")
	err = cdargs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdargs_bin_url := "https://github.com/cbxbiker61/cdargs/archive/refs/tags/2.1.bin"
	cdargs_cmd_bin := exec.Command("curl", "-L", cdargs_bin_url, "-o", "binary.bin")
	err = cdargs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdargs_src_url := "https://github.com/cbxbiker61/cdargs/archive/refs/tags/2.1.src.tar.gz"
	cdargs_cmd_src := exec.Command("curl", "-L", cdargs_src_url, "-o", "source.tar.gz")
	err = cdargs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdargs_cmd_direct := exec.Command("./binary")
	err = cdargs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
