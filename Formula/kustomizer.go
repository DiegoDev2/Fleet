package main

// Kustomizer - Package manager for distributing Kubernetes configuration as OCI artifacts
// Homepage: https://github.com/stefanprodan/kustomizer

import (
	"fmt"
	
	"os/exec"
)

func installKustomizer() {
	// Método 1: Descargar y extraer .tar.gz
	kustomizer_tar_url := "https://github.com/stefanprodan/kustomizer/archive/refs/tags/v2.2.1.tar.gz"
	kustomizer_cmd_tar := exec.Command("curl", "-L", kustomizer_tar_url, "-o", "package.tar.gz")
	err := kustomizer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kustomizer_zip_url := "https://github.com/stefanprodan/kustomizer/archive/refs/tags/v2.2.1.zip"
	kustomizer_cmd_zip := exec.Command("curl", "-L", kustomizer_zip_url, "-o", "package.zip")
	err = kustomizer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kustomizer_bin_url := "https://github.com/stefanprodan/kustomizer/archive/refs/tags/v2.2.1.bin"
	kustomizer_cmd_bin := exec.Command("curl", "-L", kustomizer_bin_url, "-o", "binary.bin")
	err = kustomizer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kustomizer_src_url := "https://github.com/stefanprodan/kustomizer/archive/refs/tags/v2.2.1.src.tar.gz"
	kustomizer_cmd_src := exec.Command("curl", "-L", kustomizer_src_url, "-o", "source.tar.gz")
	err = kustomizer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kustomizer_cmd_direct := exec.Command("./binary")
	err = kustomizer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
