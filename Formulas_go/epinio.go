package main

// Epinio - CLI for Epinio, the Application Development Engine for Kubernetes
// Homepage: https://epinio.io/

import (
	"fmt"
	
	"os/exec"
)

func installEpinio() {
	// Método 1: Descargar y extraer .tar.gz
	epinio_tar_url := "https://github.com/epinio/epinio/archive/refs/tags/v1.11.0.tar.gz"
	epinio_cmd_tar := exec.Command("curl", "-L", epinio_tar_url, "-o", "package.tar.gz")
	err := epinio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epinio_zip_url := "https://github.com/epinio/epinio/archive/refs/tags/v1.11.0.zip"
	epinio_cmd_zip := exec.Command("curl", "-L", epinio_zip_url, "-o", "package.zip")
	err = epinio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epinio_bin_url := "https://github.com/epinio/epinio/archive/refs/tags/v1.11.0.bin"
	epinio_cmd_bin := exec.Command("curl", "-L", epinio_bin_url, "-o", "binary.bin")
	err = epinio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epinio_src_url := "https://github.com/epinio/epinio/archive/refs/tags/v1.11.0.src.tar.gz"
	epinio_cmd_src := exec.Command("curl", "-L", epinio_src_url, "-o", "source.tar.gz")
	err = epinio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epinio_cmd_direct := exec.Command("./binary")
	err = epinio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
