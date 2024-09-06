package main

// Packer - Tool for creating identical machine images for multiple platforms
// Homepage: https://packer.io

import (
	"fmt"
	
	"os/exec"
)

func installPacker() {
	// Método 1: Descargar y extraer .tar.gz
	packer_tar_url := "https://github.com/hashicorp/packer/archive/refs/tags/v1.9.4.tar.gz"
	packer_cmd_tar := exec.Command("curl", "-L", packer_tar_url, "-o", "package.tar.gz")
	err := packer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	packer_zip_url := "https://github.com/hashicorp/packer/archive/refs/tags/v1.9.4.zip"
	packer_cmd_zip := exec.Command("curl", "-L", packer_zip_url, "-o", "package.zip")
	err = packer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	packer_bin_url := "https://github.com/hashicorp/packer/archive/refs/tags/v1.9.4.bin"
	packer_cmd_bin := exec.Command("curl", "-L", packer_bin_url, "-o", "binary.bin")
	err = packer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	packer_src_url := "https://github.com/hashicorp/packer/archive/refs/tags/v1.9.4.src.tar.gz"
	packer_cmd_src := exec.Command("curl", "-L", packer_src_url, "-o", "source.tar.gz")
	err = packer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	packer_cmd_direct := exec.Command("./binary")
	err = packer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
