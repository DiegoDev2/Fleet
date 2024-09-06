package main

// Glasskube - Missing Package Manager for Kubernetes
// Homepage: https://glasskube.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGlasskube() {
	// Método 1: Descargar y extraer .tar.gz
	glasskube_tar_url := "https://github.com/glasskube/glasskube/archive/refs/tags/v0.19.1.tar.gz"
	glasskube_cmd_tar := exec.Command("curl", "-L", glasskube_tar_url, "-o", "package.tar.gz")
	err := glasskube_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glasskube_zip_url := "https://github.com/glasskube/glasskube/archive/refs/tags/v0.19.1.zip"
	glasskube_cmd_zip := exec.Command("curl", "-L", glasskube_zip_url, "-o", "package.zip")
	err = glasskube_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glasskube_bin_url := "https://github.com/glasskube/glasskube/archive/refs/tags/v0.19.1.bin"
	glasskube_cmd_bin := exec.Command("curl", "-L", glasskube_bin_url, "-o", "binary.bin")
	err = glasskube_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glasskube_src_url := "https://github.com/glasskube/glasskube/archive/refs/tags/v0.19.1.src.tar.gz"
	glasskube_cmd_src := exec.Command("curl", "-L", glasskube_src_url, "-o", "source.tar.gz")
	err = glasskube_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glasskube_cmd_direct := exec.Command("./binary")
	err = glasskube_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
