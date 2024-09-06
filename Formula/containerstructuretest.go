package main

// ContainerStructureTest - Validate the structure of your container images
// Homepage: https://github.com/GoogleContainerTools/container-structure-test

import (
	"fmt"
	
	"os/exec"
)

func installContainerStructureTest() {
	// Método 1: Descargar y extraer .tar.gz
	containerstructuretest_tar_url := "https://github.com/GoogleContainerTools/container-structure-test/archive/refs/tags/v1.19.1.tar.gz"
	containerstructuretest_cmd_tar := exec.Command("curl", "-L", containerstructuretest_tar_url, "-o", "package.tar.gz")
	err := containerstructuretest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	containerstructuretest_zip_url := "https://github.com/GoogleContainerTools/container-structure-test/archive/refs/tags/v1.19.1.zip"
	containerstructuretest_cmd_zip := exec.Command("curl", "-L", containerstructuretest_zip_url, "-o", "package.zip")
	err = containerstructuretest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	containerstructuretest_bin_url := "https://github.com/GoogleContainerTools/container-structure-test/archive/refs/tags/v1.19.1.bin"
	containerstructuretest_cmd_bin := exec.Command("curl", "-L", containerstructuretest_bin_url, "-o", "binary.bin")
	err = containerstructuretest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	containerstructuretest_src_url := "https://github.com/GoogleContainerTools/container-structure-test/archive/refs/tags/v1.19.1.src.tar.gz"
	containerstructuretest_cmd_src := exec.Command("curl", "-L", containerstructuretest_src_url, "-o", "source.tar.gz")
	err = containerstructuretest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	containerstructuretest_cmd_direct := exec.Command("./binary")
	err = containerstructuretest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
