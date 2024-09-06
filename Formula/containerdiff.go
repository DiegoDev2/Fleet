package main

// ContainerDiff - Diff your Docker containers
// Homepage: https://github.com/GoogleContainerTools/container-diff

import (
	"fmt"
	
	"os/exec"
)

func installContainerDiff() {
	// Método 1: Descargar y extraer .tar.gz
	containerdiff_tar_url := "https://github.com/GoogleContainerTools/container-diff/archive/refs/tags/v0.19.0.tar.gz"
	containerdiff_cmd_tar := exec.Command("curl", "-L", containerdiff_tar_url, "-o", "package.tar.gz")
	err := containerdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	containerdiff_zip_url := "https://github.com/GoogleContainerTools/container-diff/archive/refs/tags/v0.19.0.zip"
	containerdiff_cmd_zip := exec.Command("curl", "-L", containerdiff_zip_url, "-o", "package.zip")
	err = containerdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	containerdiff_bin_url := "https://github.com/GoogleContainerTools/container-diff/archive/refs/tags/v0.19.0.bin"
	containerdiff_cmd_bin := exec.Command("curl", "-L", containerdiff_bin_url, "-o", "binary.bin")
	err = containerdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	containerdiff_src_url := "https://github.com/GoogleContainerTools/container-diff/archive/refs/tags/v0.19.0.src.tar.gz"
	containerdiff_cmd_src := exec.Command("curl", "-L", containerdiff_src_url, "-o", "source.tar.gz")
	err = containerdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	containerdiff_cmd_direct := exec.Command("./binary")
	err = containerdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
