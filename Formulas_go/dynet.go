package main

// Dynet - Dynamic Neural Network Toolkit
// Homepage: https://github.com/clab/dynet

import (
	"fmt"
	
	"os/exec"
)

func installDynet() {
	// Método 1: Descargar y extraer .tar.gz
	dynet_tar_url := "https://github.com/clab/dynet/archive/refs/tags/2.1.2.tar.gz"
	dynet_cmd_tar := exec.Command("curl", "-L", dynet_tar_url, "-o", "package.tar.gz")
	err := dynet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynet_zip_url := "https://github.com/clab/dynet/archive/refs/tags/2.1.2.zip"
	dynet_cmd_zip := exec.Command("curl", "-L", dynet_zip_url, "-o", "package.zip")
	err = dynet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynet_bin_url := "https://github.com/clab/dynet/archive/refs/tags/2.1.2.bin"
	dynet_cmd_bin := exec.Command("curl", "-L", dynet_bin_url, "-o", "binary.bin")
	err = dynet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynet_src_url := "https://github.com/clab/dynet/archive/refs/tags/2.1.2.src.tar.gz"
	dynet_cmd_src := exec.Command("curl", "-L", dynet_src_url, "-o", "source.tar.gz")
	err = dynet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynet_cmd_direct := exec.Command("./binary")
	err = dynet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
}
