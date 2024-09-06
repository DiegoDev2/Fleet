package main

// Kind - Run local Kubernetes cluster in Docker
// Homepage: https://kind.sigs.k8s.io/

import (
	"fmt"
	
	"os/exec"
)

func installKind() {
	// Método 1: Descargar y extraer .tar.gz
	kind_tar_url := "https://github.com/kubernetes-sigs/kind/archive/refs/tags/v0.24.0.tar.gz"
	kind_cmd_tar := exec.Command("curl", "-L", kind_tar_url, "-o", "package.tar.gz")
	err := kind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kind_zip_url := "https://github.com/kubernetes-sigs/kind/archive/refs/tags/v0.24.0.zip"
	kind_cmd_zip := exec.Command("curl", "-L", kind_zip_url, "-o", "package.zip")
	err = kind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kind_bin_url := "https://github.com/kubernetes-sigs/kind/archive/refs/tags/v0.24.0.bin"
	kind_cmd_bin := exec.Command("curl", "-L", kind_bin_url, "-o", "binary.bin")
	err = kind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kind_src_url := "https://github.com/kubernetes-sigs/kind/archive/refs/tags/v0.24.0.src.tar.gz"
	kind_cmd_src := exec.Command("curl", "-L", kind_src_url, "-o", "source.tar.gz")
	err = kind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kind_cmd_direct := exec.Command("./binary")
	err = kind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: docker")
exec.Command("latte", "install", "docker")
}
