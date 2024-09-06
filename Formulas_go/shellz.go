package main

// Shellz - Small utility to track and control custom shellz
// Homepage: https://github.com/evilsocket/shellz

import (
	"fmt"
	
	"os/exec"
)

func installShellz() {
	// Método 1: Descargar y extraer .tar.gz
	shellz_tar_url := "https://github.com/evilsocket/shellz/archive/refs/tags/v1.6.0.tar.gz"
	shellz_cmd_tar := exec.Command("curl", "-L", shellz_tar_url, "-o", "package.tar.gz")
	err := shellz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellz_zip_url := "https://github.com/evilsocket/shellz/archive/refs/tags/v1.6.0.zip"
	shellz_cmd_zip := exec.Command("curl", "-L", shellz_zip_url, "-o", "package.zip")
	err = shellz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellz_bin_url := "https://github.com/evilsocket/shellz/archive/refs/tags/v1.6.0.bin"
	shellz_cmd_bin := exec.Command("curl", "-L", shellz_bin_url, "-o", "binary.bin")
	err = shellz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellz_src_url := "https://github.com/evilsocket/shellz/archive/refs/tags/v1.6.0.src.tar.gz"
	shellz_cmd_src := exec.Command("curl", "-L", shellz_src_url, "-o", "source.tar.gz")
	err = shellz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellz_cmd_direct := exec.Command("./binary")
	err = shellz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
