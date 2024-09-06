package main

// HelmDocs - Tool for automatically generating markdown documentation for helm charts
// Homepage: https://github.com/norwoodj/helm-docs

import (
	"fmt"
	
	"os/exec"
)

func installHelmDocs() {
	// Método 1: Descargar y extraer .tar.gz
	helmdocs_tar_url := "https://github.com/norwoodj/helm-docs/archive/refs/tags/v1.14.2.tar.gz"
	helmdocs_cmd_tar := exec.Command("curl", "-L", helmdocs_tar_url, "-o", "package.tar.gz")
	err := helmdocs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	helmdocs_zip_url := "https://github.com/norwoodj/helm-docs/archive/refs/tags/v1.14.2.zip"
	helmdocs_cmd_zip := exec.Command("curl", "-L", helmdocs_zip_url, "-o", "package.zip")
	err = helmdocs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	helmdocs_bin_url := "https://github.com/norwoodj/helm-docs/archive/refs/tags/v1.14.2.bin"
	helmdocs_cmd_bin := exec.Command("curl", "-L", helmdocs_bin_url, "-o", "binary.bin")
	err = helmdocs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	helmdocs_src_url := "https://github.com/norwoodj/helm-docs/archive/refs/tags/v1.14.2.src.tar.gz"
	helmdocs_cmd_src := exec.Command("curl", "-L", helmdocs_src_url, "-o", "source.tar.gz")
	err = helmdocs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	helmdocs_cmd_direct := exec.Command("./binary")
	err = helmdocs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
