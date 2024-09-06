package main

// Diffoci - Diff for Docker and OCI container images
// Homepage: https://github.com/reproducible-containers/diffoci

import (
	"fmt"
	
	"os/exec"
)

func installDiffoci() {
	// Método 1: Descargar y extraer .tar.gz
	diffoci_tar_url := "https://github.com/reproducible-containers/diffoci/archive/refs/tags/v0.1.5.tar.gz"
	diffoci_cmd_tar := exec.Command("curl", "-L", diffoci_tar_url, "-o", "package.tar.gz")
	err := diffoci_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffoci_zip_url := "https://github.com/reproducible-containers/diffoci/archive/refs/tags/v0.1.5.zip"
	diffoci_cmd_zip := exec.Command("curl", "-L", diffoci_zip_url, "-o", "package.zip")
	err = diffoci_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffoci_bin_url := "https://github.com/reproducible-containers/diffoci/archive/refs/tags/v0.1.5.bin"
	diffoci_cmd_bin := exec.Command("curl", "-L", diffoci_bin_url, "-o", "binary.bin")
	err = diffoci_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffoci_src_url := "https://github.com/reproducible-containers/diffoci/archive/refs/tags/v0.1.5.src.tar.gz"
	diffoci_cmd_src := exec.Command("curl", "-L", diffoci_src_url, "-o", "source.tar.gz")
	err = diffoci_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffoci_cmd_direct := exec.Command("./binary")
	err = diffoci_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
