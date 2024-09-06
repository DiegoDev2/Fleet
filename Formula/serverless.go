package main

// Serverless - Build applications with serverless architectures
// Homepage: https://www.serverless.com/

import (
	"fmt"
	
	"os/exec"
)

func installServerless() {
	// Método 1: Descargar y extraer .tar.gz
	serverless_tar_url := "https://github.com/serverless/serverless/archive/refs/tags/v3.39.0.tar.gz"
	serverless_cmd_tar := exec.Command("curl", "-L", serverless_tar_url, "-o", "package.tar.gz")
	err := serverless_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	serverless_zip_url := "https://github.com/serverless/serverless/archive/refs/tags/v3.39.0.zip"
	serverless_cmd_zip := exec.Command("curl", "-L", serverless_zip_url, "-o", "package.zip")
	err = serverless_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	serverless_bin_url := "https://github.com/serverless/serverless/archive/refs/tags/v3.39.0.bin"
	serverless_cmd_bin := exec.Command("curl", "-L", serverless_bin_url, "-o", "binary.bin")
	err = serverless_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	serverless_src_url := "https://github.com/serverless/serverless/archive/refs/tags/v3.39.0.src.tar.gz"
	serverless_cmd_src := exec.Command("curl", "-L", serverless_src_url, "-o", "source.tar.gz")
	err = serverless_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	serverless_cmd_direct := exec.Command("./binary")
	err = serverless_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
