package main

// Svgo - Nodejs-based tool for optimizing SVG vector graphics files
// Homepage: https://github.com/svg/svgo

import (
	"fmt"
	
	"os/exec"
)

func installSvgo() {
	// Método 1: Descargar y extraer .tar.gz
	svgo_tar_url := "https://github.com/svg/svgo/archive/refs/tags/v3.3.2.tar.gz"
	svgo_cmd_tar := exec.Command("curl", "-L", svgo_tar_url, "-o", "package.tar.gz")
	err := svgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svgo_zip_url := "https://github.com/svg/svgo/archive/refs/tags/v3.3.2.zip"
	svgo_cmd_zip := exec.Command("curl", "-L", svgo_zip_url, "-o", "package.zip")
	err = svgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svgo_bin_url := "https://github.com/svg/svgo/archive/refs/tags/v3.3.2.bin"
	svgo_cmd_bin := exec.Command("curl", "-L", svgo_bin_url, "-o", "binary.bin")
	err = svgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svgo_src_url := "https://github.com/svg/svgo/archive/refs/tags/v3.3.2.src.tar.gz"
	svgo_cmd_src := exec.Command("curl", "-L", svgo_src_url, "-o", "source.tar.gz")
	err = svgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svgo_cmd_direct := exec.Command("./binary")
	err = svgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
