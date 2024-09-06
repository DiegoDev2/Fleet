package main

// VrcGet - Open Source alternative of Command-line client of VRChat Package Manager
// Homepage: https://github.com/vrc-get/vrc-get

import (
	"fmt"
	
	"os/exec"
)

func installVrcGet() {
	// Método 1: Descargar y extraer .tar.gz
	vrcget_tar_url := "https://github.com/vrc-get/vrc-get/archive/refs/tags/v1.8.1.tar.gz"
	vrcget_cmd_tar := exec.Command("curl", "-L", vrcget_tar_url, "-o", "package.tar.gz")
	err := vrcget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vrcget_zip_url := "https://github.com/vrc-get/vrc-get/archive/refs/tags/v1.8.1.zip"
	vrcget_cmd_zip := exec.Command("curl", "-L", vrcget_zip_url, "-o", "package.zip")
	err = vrcget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vrcget_bin_url := "https://github.com/vrc-get/vrc-get/archive/refs/tags/v1.8.1.bin"
	vrcget_cmd_bin := exec.Command("curl", "-L", vrcget_bin_url, "-o", "binary.bin")
	err = vrcget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vrcget_src_url := "https://github.com/vrc-get/vrc-get/archive/refs/tags/v1.8.1.src.tar.gz"
	vrcget_cmd_src := exec.Command("curl", "-L", vrcget_src_url, "-o", "source.tar.gz")
	err = vrcget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vrcget_cmd_direct := exec.Command("./binary")
	err = vrcget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
