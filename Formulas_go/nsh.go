package main

// Nsh - Fish-like, POSIX-compatible shell
// Homepage: https://github.com/nuta/nsh

import (
	"fmt"
	
	"os/exec"
)

func installNsh() {
	// Método 1: Descargar y extraer .tar.gz
	nsh_tar_url := "https://github.com/nuta/nsh/archive/refs/tags/v0.4.2.tar.gz"
	nsh_cmd_tar := exec.Command("curl", "-L", nsh_tar_url, "-o", "package.tar.gz")
	err := nsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsh_zip_url := "https://github.com/nuta/nsh/archive/refs/tags/v0.4.2.zip"
	nsh_cmd_zip := exec.Command("curl", "-L", nsh_zip_url, "-o", "package.zip")
	err = nsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsh_bin_url := "https://github.com/nuta/nsh/archive/refs/tags/v0.4.2.bin"
	nsh_cmd_bin := exec.Command("curl", "-L", nsh_bin_url, "-o", "binary.bin")
	err = nsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsh_src_url := "https://github.com/nuta/nsh/archive/refs/tags/v0.4.2.src.tar.gz"
	nsh_cmd_src := exec.Command("curl", "-L", nsh_src_url, "-o", "source.tar.gz")
	err = nsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsh_cmd_direct := exec.Command("./binary")
	err = nsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
