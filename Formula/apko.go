package main

// Apko - Build OCI images from APK packages directly without Dockerfile
// Homepage: https://github.com/chainguard-dev/apko

import (
	"fmt"
	
	"os/exec"
)

func installApko() {
	// Método 1: Descargar y extraer .tar.gz
	apko_tar_url := "https://github.com/chainguard-dev/apko/archive/refs/tags/v0.14.7.tar.gz"
	apko_cmd_tar := exec.Command("curl", "-L", apko_tar_url, "-o", "package.tar.gz")
	err := apko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apko_zip_url := "https://github.com/chainguard-dev/apko/archive/refs/tags/v0.14.7.zip"
	apko_cmd_zip := exec.Command("curl", "-L", apko_zip_url, "-o", "package.zip")
	err = apko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apko_bin_url := "https://github.com/chainguard-dev/apko/archive/refs/tags/v0.14.7.bin"
	apko_cmd_bin := exec.Command("curl", "-L", apko_bin_url, "-o", "binary.bin")
	err = apko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apko_src_url := "https://github.com/chainguard-dev/apko/archive/refs/tags/v0.14.7.src.tar.gz"
	apko_cmd_src := exec.Command("curl", "-L", apko_src_url, "-o", "source.tar.gz")
	err = apko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apko_cmd_direct := exec.Command("./binary")
	err = apko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
