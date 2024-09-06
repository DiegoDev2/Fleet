package main

// Kyverno - Kubernetes Native Policy Management
// Homepage: https://kyverno.io/

import (
	"fmt"
	
	"os/exec"
)

func installKyverno() {
	// Método 1: Descargar y extraer .tar.gz
	kyverno_tar_url := "https://github.com/kyverno/kyverno/archive/refs/tags/v1.12.5.tar.gz"
	kyverno_cmd_tar := exec.Command("curl", "-L", kyverno_tar_url, "-o", "package.tar.gz")
	err := kyverno_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kyverno_zip_url := "https://github.com/kyverno/kyverno/archive/refs/tags/v1.12.5.zip"
	kyverno_cmd_zip := exec.Command("curl", "-L", kyverno_zip_url, "-o", "package.zip")
	err = kyverno_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kyverno_bin_url := "https://github.com/kyverno/kyverno/archive/refs/tags/v1.12.5.bin"
	kyverno_cmd_bin := exec.Command("curl", "-L", kyverno_bin_url, "-o", "binary.bin")
	err = kyverno_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kyverno_src_url := "https://github.com/kyverno/kyverno/archive/refs/tags/v1.12.5.src.tar.gz"
	kyverno_cmd_src := exec.Command("curl", "-L", kyverno_src_url, "-o", "source.tar.gz")
	err = kyverno_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kyverno_cmd_direct := exec.Command("./binary")
	err = kyverno_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
