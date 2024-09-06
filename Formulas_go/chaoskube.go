package main

// Chaoskube - Periodically kills random pods in your Kubernetes cluster
// Homepage: https://github.com/linki/chaoskube

import (
	"fmt"
	
	"os/exec"
)

func installChaoskube() {
	// Método 1: Descargar y extraer .tar.gz
	chaoskube_tar_url := "https://github.com/linki/chaoskube/archive/refs/tags/v0.33.0.tar.gz"
	chaoskube_cmd_tar := exec.Command("curl", "-L", chaoskube_tar_url, "-o", "package.tar.gz")
	err := chaoskube_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chaoskube_zip_url := "https://github.com/linki/chaoskube/archive/refs/tags/v0.33.0.zip"
	chaoskube_cmd_zip := exec.Command("curl", "-L", chaoskube_zip_url, "-o", "package.zip")
	err = chaoskube_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chaoskube_bin_url := "https://github.com/linki/chaoskube/archive/refs/tags/v0.33.0.bin"
	chaoskube_cmd_bin := exec.Command("curl", "-L", chaoskube_bin_url, "-o", "binary.bin")
	err = chaoskube_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chaoskube_src_url := "https://github.com/linki/chaoskube/archive/refs/tags/v0.33.0.src.tar.gz"
	chaoskube_cmd_src := exec.Command("curl", "-L", chaoskube_src_url, "-o", "source.tar.gz")
	err = chaoskube_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chaoskube_cmd_direct := exec.Command("./binary")
	err = chaoskube_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
