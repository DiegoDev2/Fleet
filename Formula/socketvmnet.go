package main

// SocketVmnet - Daemon to provide vmnet.framework support for rootless QEMU
// Homepage: https://github.com/lima-vm/socket_vmnet

import (
	"fmt"
	
	"os/exec"
)

func installSocketVmnet() {
	// Método 1: Descargar y extraer .tar.gz
	socketvmnet_tar_url := "https://github.com/lima-vm/socket_vmnet/archive/refs/tags/v1.1.4.tar.gz"
	socketvmnet_cmd_tar := exec.Command("curl", "-L", socketvmnet_tar_url, "-o", "package.tar.gz")
	err := socketvmnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	socketvmnet_zip_url := "https://github.com/lima-vm/socket_vmnet/archive/refs/tags/v1.1.4.zip"
	socketvmnet_cmd_zip := exec.Command("curl", "-L", socketvmnet_zip_url, "-o", "package.zip")
	err = socketvmnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	socketvmnet_bin_url := "https://github.com/lima-vm/socket_vmnet/archive/refs/tags/v1.1.4.bin"
	socketvmnet_cmd_bin := exec.Command("curl", "-L", socketvmnet_bin_url, "-o", "binary.bin")
	err = socketvmnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	socketvmnet_src_url := "https://github.com/lima-vm/socket_vmnet/archive/refs/tags/v1.1.4.src.tar.gz"
	socketvmnet_cmd_src := exec.Command("curl", "-L", socketvmnet_src_url, "-o", "source.tar.gz")
	err = socketvmnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	socketvmnet_cmd_direct := exec.Command("./binary")
	err = socketvmnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
