package main

// Clair - Vulnerability Static Analysis for Containers
// Homepage: https://github.com/quay/clair

import (
	"fmt"
	
	"os/exec"
)

func installClair() {
	// Método 1: Descargar y extraer .tar.gz
	clair_tar_url := "https://github.com/quay/clair/archive/refs/tags/v4.7.4.tar.gz"
	clair_cmd_tar := exec.Command("curl", "-L", clair_tar_url, "-o", "package.tar.gz")
	err := clair_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clair_zip_url := "https://github.com/quay/clair/archive/refs/tags/v4.7.4.zip"
	clair_cmd_zip := exec.Command("curl", "-L", clair_zip_url, "-o", "package.zip")
	err = clair_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clair_bin_url := "https://github.com/quay/clair/archive/refs/tags/v4.7.4.bin"
	clair_cmd_bin := exec.Command("curl", "-L", clair_bin_url, "-o", "binary.bin")
	err = clair_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clair_src_url := "https://github.com/quay/clair/archive/refs/tags/v4.7.4.src.tar.gz"
	clair_cmd_src := exec.Command("curl", "-L", clair_src_url, "-o", "source.tar.gz")
	err = clair_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clair_cmd_direct := exec.Command("./binary")
	err = clair_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
