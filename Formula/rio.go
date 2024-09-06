package main

// Rio - Hardware-accelerated GPU terminal emulator powered by WebGPU
// Homepage: https://raphamorim.io/rio/

import (
	"fmt"
	
	"os/exec"
)

func installRio() {
	// Método 1: Descargar y extraer .tar.gz
	rio_tar_url := "https://github.com/raphamorim/rio/archive/refs/tags/v0.1.11.tar.gz"
	rio_cmd_tar := exec.Command("curl", "-L", rio_tar_url, "-o", "package.tar.gz")
	err := rio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rio_zip_url := "https://github.com/raphamorim/rio/archive/refs/tags/v0.1.11.zip"
	rio_cmd_zip := exec.Command("curl", "-L", rio_zip_url, "-o", "package.zip")
	err = rio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rio_bin_url := "https://github.com/raphamorim/rio/archive/refs/tags/v0.1.11.bin"
	rio_cmd_bin := exec.Command("curl", "-L", rio_bin_url, "-o", "binary.bin")
	err = rio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rio_src_url := "https://github.com/raphamorim/rio/archive/refs/tags/v0.1.11.src.tar.gz"
	rio_cmd_src := exec.Command("curl", "-L", rio_src_url, "-o", "source.tar.gz")
	err = rio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rio_cmd_direct := exec.Command("./binary")
	err = rio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
