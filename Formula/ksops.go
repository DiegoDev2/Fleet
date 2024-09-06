package main

// Ksops - Flexible Kustomize Plugin for SOPS Encrypted Resources
// Homepage: https://github.com/viaduct-ai/kustomize-sops

import (
	"fmt"
	
	"os/exec"
)

func installKsops() {
	// Método 1: Descargar y extraer .tar.gz
	ksops_tar_url := "https://github.com/viaduct-ai/kustomize-sops/archive/refs/tags/v4.3.2.tar.gz"
	ksops_cmd_tar := exec.Command("curl", "-L", ksops_tar_url, "-o", "package.tar.gz")
	err := ksops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ksops_zip_url := "https://github.com/viaduct-ai/kustomize-sops/archive/refs/tags/v4.3.2.zip"
	ksops_cmd_zip := exec.Command("curl", "-L", ksops_zip_url, "-o", "package.zip")
	err = ksops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ksops_bin_url := "https://github.com/viaduct-ai/kustomize-sops/archive/refs/tags/v4.3.2.bin"
	ksops_cmd_bin := exec.Command("curl", "-L", ksops_bin_url, "-o", "binary.bin")
	err = ksops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ksops_src_url := "https://github.com/viaduct-ai/kustomize-sops/archive/refs/tags/v4.3.2.src.tar.gz"
	ksops_cmd_src := exec.Command("curl", "-L", ksops_src_url, "-o", "source.tar.gz")
	err = ksops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ksops_cmd_direct := exec.Command("./binary")
	err = ksops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
