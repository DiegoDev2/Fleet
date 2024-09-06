package main

// Triton - Joyent Triton CLI
// Homepage: https://www.npmjs.com/package/triton

import (
	"fmt"
	
	"os/exec"
)

func installTriton() {
	// Método 1: Descargar y extraer .tar.gz
	triton_tar_url := "https://registry.npmjs.org/triton/-/triton-7.17.0.tgz"
	triton_cmd_tar := exec.Command("curl", "-L", triton_tar_url, "-o", "package.tar.gz")
	err := triton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	triton_zip_url := "https://registry.npmjs.org/triton/-/triton-7.17.0.tgz"
	triton_cmd_zip := exec.Command("curl", "-L", triton_zip_url, "-o", "package.zip")
	err = triton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	triton_bin_url := "https://registry.npmjs.org/triton/-/triton-7.17.0.tgz"
	triton_cmd_bin := exec.Command("curl", "-L", triton_bin_url, "-o", "binary.bin")
	err = triton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	triton_src_url := "https://registry.npmjs.org/triton/-/triton-7.17.0.tgz"
	triton_cmd_src := exec.Command("curl", "-L", triton_src_url, "-o", "source.tar.gz")
	err = triton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	triton_cmd_direct := exec.Command("./binary")
	err = triton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
