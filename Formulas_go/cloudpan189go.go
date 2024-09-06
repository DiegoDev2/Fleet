package main

// Cloudpan189Go - Command-line client tool for Cloud189 web disk
// Homepage: https://github.com/tickstep/cloudpan189-go

import (
	"fmt"
	
	"os/exec"
)

func installCloudpan189Go() {
	// Método 1: Descargar y extraer .tar.gz
	cloudpan189go_tar_url := "https://github.com/tickstep/cloudpan189-go/archive/refs/tags/v0.1.3.tar.gz"
	cloudpan189go_cmd_tar := exec.Command("curl", "-L", cloudpan189go_tar_url, "-o", "package.tar.gz")
	err := cloudpan189go_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudpan189go_zip_url := "https://github.com/tickstep/cloudpan189-go/archive/refs/tags/v0.1.3.zip"
	cloudpan189go_cmd_zip := exec.Command("curl", "-L", cloudpan189go_zip_url, "-o", "package.zip")
	err = cloudpan189go_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudpan189go_bin_url := "https://github.com/tickstep/cloudpan189-go/archive/refs/tags/v0.1.3.bin"
	cloudpan189go_cmd_bin := exec.Command("curl", "-L", cloudpan189go_bin_url, "-o", "binary.bin")
	err = cloudpan189go_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudpan189go_src_url := "https://github.com/tickstep/cloudpan189-go/archive/refs/tags/v0.1.3.src.tar.gz"
	cloudpan189go_cmd_src := exec.Command("curl", "-L", cloudpan189go_src_url, "-o", "source.tar.gz")
	err = cloudpan189go_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudpan189go_cmd_direct := exec.Command("./binary")
	err = cloudpan189go_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
exec.Command("latte", "install", "go@1.22")
}
