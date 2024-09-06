package main

// Lima - Linux virtual machines
// Homepage: https://lima-vm.io/

import (
	"fmt"
	
	"os/exec"
)

func installLima() {
	// Método 1: Descargar y extraer .tar.gz
	lima_tar_url := "https://github.com/lima-vm/lima/archive/refs/tags/v0.23.2.tar.gz"
	lima_cmd_tar := exec.Command("curl", "-L", lima_tar_url, "-o", "package.tar.gz")
	err := lima_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lima_zip_url := "https://github.com/lima-vm/lima/archive/refs/tags/v0.23.2.zip"
	lima_cmd_zip := exec.Command("curl", "-L", lima_zip_url, "-o", "package.zip")
	err = lima_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lima_bin_url := "https://github.com/lima-vm/lima/archive/refs/tags/v0.23.2.bin"
	lima_cmd_bin := exec.Command("curl", "-L", lima_bin_url, "-o", "binary.bin")
	err = lima_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lima_src_url := "https://github.com/lima-vm/lima/archive/refs/tags/v0.23.2.src.tar.gz"
	lima_cmd_src := exec.Command("curl", "-L", lima_src_url, "-o", "source.tar.gz")
	err = lima_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lima_cmd_direct := exec.Command("./binary")
	err = lima_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: qemu")
	exec.Command("latte", "install", "qemu").Run()
}
