package main

// Kompose - Tool to move from `docker-compose` to Kubernetes
// Homepage: https://kompose.io/

import (
	"fmt"
	
	"os/exec"
)

func installKompose() {
	// Método 1: Descargar y extraer .tar.gz
	kompose_tar_url := "https://github.com/kubernetes/kompose/archive/refs/tags/v1.34.0.tar.gz"
	kompose_cmd_tar := exec.Command("curl", "-L", kompose_tar_url, "-o", "package.tar.gz")
	err := kompose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kompose_zip_url := "https://github.com/kubernetes/kompose/archive/refs/tags/v1.34.0.zip"
	kompose_cmd_zip := exec.Command("curl", "-L", kompose_zip_url, "-o", "package.zip")
	err = kompose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kompose_bin_url := "https://github.com/kubernetes/kompose/archive/refs/tags/v1.34.0.bin"
	kompose_cmd_bin := exec.Command("curl", "-L", kompose_bin_url, "-o", "binary.bin")
	err = kompose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kompose_src_url := "https://github.com/kubernetes/kompose/archive/refs/tags/v1.34.0.src.tar.gz"
	kompose_cmd_src := exec.Command("curl", "-L", kompose_src_url, "-o", "source.tar.gz")
	err = kompose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kompose_cmd_direct := exec.Command("./binary")
	err = kompose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
