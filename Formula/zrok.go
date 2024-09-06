package main

// Zrok - Geo-scale, next-generation sharing platform built on top of OpenZiti
// Homepage: https://zrok.io

import (
	"fmt"
	
	"os/exec"
)

func installZrok() {
	// Método 1: Descargar y extraer .tar.gz
	zrok_tar_url := "https://github.com/openziti/zrok/archive/refs/tags/v0.4.39.tar.gz"
	zrok_cmd_tar := exec.Command("curl", "-L", zrok_tar_url, "-o", "package.tar.gz")
	err := zrok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zrok_zip_url := "https://github.com/openziti/zrok/archive/refs/tags/v0.4.39.zip"
	zrok_cmd_zip := exec.Command("curl", "-L", zrok_zip_url, "-o", "package.zip")
	err = zrok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zrok_bin_url := "https://github.com/openziti/zrok/archive/refs/tags/v0.4.39.bin"
	zrok_cmd_bin := exec.Command("curl", "-L", zrok_bin_url, "-o", "binary.bin")
	err = zrok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zrok_src_url := "https://github.com/openziti/zrok/archive/refs/tags/v0.4.39.src.tar.gz"
	zrok_cmd_src := exec.Command("curl", "-L", zrok_src_url, "-o", "source.tar.gz")
	err = zrok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zrok_cmd_direct := exec.Command("./binary")
	err = zrok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
