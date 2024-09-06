package main

// Webpod - Deploy websites and apps anywhere
// Homepage: https://webpod.dev

import (
	"fmt"
	
	"os/exec"
)

func installWebpod() {
	// Método 1: Descargar y extraer .tar.gz
	webpod_tar_url := "https://registry.npmjs.org/webpod/-/webpod-1.0.0.tgz"
	webpod_cmd_tar := exec.Command("curl", "-L", webpod_tar_url, "-o", "package.tar.gz")
	err := webpod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webpod_zip_url := "https://registry.npmjs.org/webpod/-/webpod-1.0.0.tgz"
	webpod_cmd_zip := exec.Command("curl", "-L", webpod_zip_url, "-o", "package.zip")
	err = webpod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webpod_bin_url := "https://registry.npmjs.org/webpod/-/webpod-1.0.0.tgz"
	webpod_cmd_bin := exec.Command("curl", "-L", webpod_bin_url, "-o", "binary.bin")
	err = webpod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webpod_src_url := "https://registry.npmjs.org/webpod/-/webpod-1.0.0.tgz"
	webpod_cmd_src := exec.Command("curl", "-L", webpod_src_url, "-o", "source.tar.gz")
	err = webpod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webpod_cmd_direct := exec.Command("./binary")
	err = webpod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
