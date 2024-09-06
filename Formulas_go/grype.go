package main

// Grype - Vulnerability scanner for container images and filesystems
// Homepage: https://github.com/anchore/grype

import (
	"fmt"
	
	"os/exec"
)

func installGrype() {
	// Método 1: Descargar y extraer .tar.gz
	grype_tar_url := "https://github.com/anchore/grype/archive/refs/tags/v0.80.0.tar.gz"
	grype_cmd_tar := exec.Command("curl", "-L", grype_tar_url, "-o", "package.tar.gz")
	err := grype_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grype_zip_url := "https://github.com/anchore/grype/archive/refs/tags/v0.80.0.zip"
	grype_cmd_zip := exec.Command("curl", "-L", grype_zip_url, "-o", "package.zip")
	err = grype_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grype_bin_url := "https://github.com/anchore/grype/archive/refs/tags/v0.80.0.bin"
	grype_cmd_bin := exec.Command("curl", "-L", grype_bin_url, "-o", "binary.bin")
	err = grype_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grype_src_url := "https://github.com/anchore/grype/archive/refs/tags/v0.80.0.src.tar.gz"
	grype_cmd_src := exec.Command("curl", "-L", grype_src_url, "-o", "source.tar.gz")
	err = grype_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grype_cmd_direct := exec.Command("./binary")
	err = grype_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
