package main

// Virtctl - Allows for using more advanced kubevirt features
// Homepage: https://kubevirt.io/

import (
	"fmt"
	
	"os/exec"
)

func installVirtctl() {
	// Método 1: Descargar y extraer .tar.gz
	virtctl_tar_url := "https://github.com/kubevirt/kubevirt/archive/refs/tags/v1.3.1.tar.gz"
	virtctl_cmd_tar := exec.Command("curl", "-L", virtctl_tar_url, "-o", "package.tar.gz")
	err := virtctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtctl_zip_url := "https://github.com/kubevirt/kubevirt/archive/refs/tags/v1.3.1.zip"
	virtctl_cmd_zip := exec.Command("curl", "-L", virtctl_zip_url, "-o", "package.zip")
	err = virtctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtctl_bin_url := "https://github.com/kubevirt/kubevirt/archive/refs/tags/v1.3.1.bin"
	virtctl_cmd_bin := exec.Command("curl", "-L", virtctl_bin_url, "-o", "binary.bin")
	err = virtctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtctl_src_url := "https://github.com/kubevirt/kubevirt/archive/refs/tags/v1.3.1.src.tar.gz"
	virtctl_cmd_src := exec.Command("curl", "-L", virtctl_src_url, "-o", "source.tar.gz")
	err = virtctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtctl_cmd_direct := exec.Command("./binary")
	err = virtctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
