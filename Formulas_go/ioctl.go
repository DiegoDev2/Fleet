package main

// Ioctl - Command-line interface for interacting with the IoTeX blockchain
// Homepage: https://docs.iotex.io/

import (
	"fmt"
	
	"os/exec"
)

func installIoctl() {
	// Método 1: Descargar y extraer .tar.gz
	ioctl_tar_url := "https://github.com/iotexproject/iotex-core/archive/refs/tags/v2.0.4.tar.gz"
	ioctl_cmd_tar := exec.Command("curl", "-L", ioctl_tar_url, "-o", "package.tar.gz")
	err := ioctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ioctl_zip_url := "https://github.com/iotexproject/iotex-core/archive/refs/tags/v2.0.4.zip"
	ioctl_cmd_zip := exec.Command("curl", "-L", ioctl_zip_url, "-o", "package.zip")
	err = ioctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ioctl_bin_url := "https://github.com/iotexproject/iotex-core/archive/refs/tags/v2.0.4.bin"
	ioctl_cmd_bin := exec.Command("curl", "-L", ioctl_bin_url, "-o", "binary.bin")
	err = ioctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ioctl_src_url := "https://github.com/iotexproject/iotex-core/archive/refs/tags/v2.0.4.src.tar.gz"
	ioctl_cmd_src := exec.Command("curl", "-L", ioctl_src_url, "-o", "source.tar.gz")
	err = ioctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ioctl_cmd_direct := exec.Command("./binary")
	err = ioctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
