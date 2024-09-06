package main

// Vkectl - Command-Line Interface for VKE(VolcanoEngine Kubernetes Engine)
// Homepage: https://github.com/volcengine/vkectl

import (
	"fmt"
	
	"os/exec"
)

func installVkectl() {
	// Método 1: Descargar y extraer .tar.gz
	vkectl_tar_url := "https://github.com/volcengine/vkectl/archive/refs/tags/v0.1.1.tar.gz"
	vkectl_cmd_tar := exec.Command("curl", "-L", vkectl_tar_url, "-o", "package.tar.gz")
	err := vkectl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vkectl_zip_url := "https://github.com/volcengine/vkectl/archive/refs/tags/v0.1.1.zip"
	vkectl_cmd_zip := exec.Command("curl", "-L", vkectl_zip_url, "-o", "package.zip")
	err = vkectl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vkectl_bin_url := "https://github.com/volcengine/vkectl/archive/refs/tags/v0.1.1.bin"
	vkectl_cmd_bin := exec.Command("curl", "-L", vkectl_bin_url, "-o", "binary.bin")
	err = vkectl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vkectl_src_url := "https://github.com/volcengine/vkectl/archive/refs/tags/v0.1.1.src.tar.gz"
	vkectl_cmd_src := exec.Command("curl", "-L", vkectl_src_url, "-o", "source.tar.gz")
	err = vkectl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vkectl_cmd_direct := exec.Command("./binary")
	err = vkectl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.17")
	exec.Command("latte", "install", "go@1.17").Run()
}
