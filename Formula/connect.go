package main

// Connect - Provides SOCKS and HTTPS proxy support to SSH
// Homepage: https://github.com/gotoh/ssh-connect

import (
	"fmt"
	
	"os/exec"
)

func installConnect() {
	// Método 1: Descargar y extraer .tar.gz
	connect_tar_url := "https://github.com/gotoh/ssh-connect/archive/refs/tags/1.105.tar.gz"
	connect_cmd_tar := exec.Command("curl", "-L", connect_tar_url, "-o", "package.tar.gz")
	err := connect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	connect_zip_url := "https://github.com/gotoh/ssh-connect/archive/refs/tags/1.105.zip"
	connect_cmd_zip := exec.Command("curl", "-L", connect_zip_url, "-o", "package.zip")
	err = connect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	connect_bin_url := "https://github.com/gotoh/ssh-connect/archive/refs/tags/1.105.bin"
	connect_cmd_bin := exec.Command("curl", "-L", connect_bin_url, "-o", "binary.bin")
	err = connect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	connect_src_url := "https://github.com/gotoh/ssh-connect/archive/refs/tags/1.105.src.tar.gz"
	connect_cmd_src := exec.Command("curl", "-L", connect_src_url, "-o", "source.tar.gz")
	err = connect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	connect_cmd_direct := exec.Command("./binary")
	err = connect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
