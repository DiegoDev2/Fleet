package main

// Xaric - IRC client
// Homepage: https://xaric.org/

import (
	"fmt"
	
	"os/exec"
)

func installXaric() {
	// Método 1: Descargar y extraer .tar.gz
	xaric_tar_url := "https://xaric.org/software/xaric/releases/xaric-0.13.9.tar.gz"
	xaric_cmd_tar := exec.Command("curl", "-L", xaric_tar_url, "-o", "package.tar.gz")
	err := xaric_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xaric_zip_url := "https://xaric.org/software/xaric/releases/xaric-0.13.9.zip"
	xaric_cmd_zip := exec.Command("curl", "-L", xaric_zip_url, "-o", "package.zip")
	err = xaric_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xaric_bin_url := "https://xaric.org/software/xaric/releases/xaric-0.13.9.bin"
	xaric_cmd_bin := exec.Command("curl", "-L", xaric_bin_url, "-o", "binary.bin")
	err = xaric_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xaric_src_url := "https://xaric.org/software/xaric/releases/xaric-0.13.9.src.tar.gz"
	xaric_cmd_src := exec.Command("curl", "-L", xaric_src_url, "-o", "source.tar.gz")
	err = xaric_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xaric_cmd_direct := exec.Command("./binary")
	err = xaric_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
