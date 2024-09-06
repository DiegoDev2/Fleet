package main

// Kubehound - Tool for building Kubernetes attack paths
// Homepage: https://kubehound.io

import (
	"fmt"
	
	"os/exec"
)

func installKubehound() {
	// Método 1: Descargar y extraer .tar.gz
	kubehound_tar_url := "https://github.com/DataDog/KubeHound/archive/refs/tags/v1.4.0.tar.gz"
	kubehound_cmd_tar := exec.Command("curl", "-L", kubehound_tar_url, "-o", "package.tar.gz")
	err := kubehound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubehound_zip_url := "https://github.com/DataDog/KubeHound/archive/refs/tags/v1.4.0.zip"
	kubehound_cmd_zip := exec.Command("curl", "-L", kubehound_zip_url, "-o", "package.zip")
	err = kubehound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubehound_bin_url := "https://github.com/DataDog/KubeHound/archive/refs/tags/v1.4.0.bin"
	kubehound_cmd_bin := exec.Command("curl", "-L", kubehound_bin_url, "-o", "binary.bin")
	err = kubehound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubehound_src_url := "https://github.com/DataDog/KubeHound/archive/refs/tags/v1.4.0.src.tar.gz"
	kubehound_cmd_src := exec.Command("curl", "-L", kubehound_src_url, "-o", "source.tar.gz")
	err = kubehound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubehound_cmd_direct := exec.Command("./binary")
	err = kubehound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
