package main

// Devtodo - Command-line task lists
// Homepage: https://swapoff.org/devtodo.html

import (
	"fmt"
	
	"os/exec"
)

func installDevtodo() {
	// Método 1: Descargar y extraer .tar.gz
	devtodo_tar_url := "https://swapoff.org/files/devtodo/devtodo-0.1.20.tar.gz"
	devtodo_cmd_tar := exec.Command("curl", "-L", devtodo_tar_url, "-o", "package.tar.gz")
	err := devtodo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	devtodo_zip_url := "https://swapoff.org/files/devtodo/devtodo-0.1.20.zip"
	devtodo_cmd_zip := exec.Command("curl", "-L", devtodo_zip_url, "-o", "package.zip")
	err = devtodo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	devtodo_bin_url := "https://swapoff.org/files/devtodo/devtodo-0.1.20.bin"
	devtodo_cmd_bin := exec.Command("curl", "-L", devtodo_bin_url, "-o", "binary.bin")
	err = devtodo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	devtodo_src_url := "https://swapoff.org/files/devtodo/devtodo-0.1.20.src.tar.gz"
	devtodo_cmd_src := exec.Command("curl", "-L", devtodo_src_url, "-o", "source.tar.gz")
	err = devtodo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	devtodo_cmd_direct := exec.Command("./binary")
	err = devtodo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
