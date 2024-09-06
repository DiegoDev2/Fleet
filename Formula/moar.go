package main

// Moar - Nice to use pager for humans
// Homepage: https://github.com/walles/moar

import (
	"fmt"
	
	"os/exec"
)

func installMoar() {
	// Método 1: Descargar y extraer .tar.gz
	moar_tar_url := "https://github.com/walles/moar/archive/refs/tags/v1.26.0.tar.gz"
	moar_cmd_tar := exec.Command("curl", "-L", moar_tar_url, "-o", "package.tar.gz")
	err := moar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moar_zip_url := "https://github.com/walles/moar/archive/refs/tags/v1.26.0.zip"
	moar_cmd_zip := exec.Command("curl", "-L", moar_zip_url, "-o", "package.zip")
	err = moar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moar_bin_url := "https://github.com/walles/moar/archive/refs/tags/v1.26.0.bin"
	moar_cmd_bin := exec.Command("curl", "-L", moar_bin_url, "-o", "binary.bin")
	err = moar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moar_src_url := "https://github.com/walles/moar/archive/refs/tags/v1.26.0.src.tar.gz"
	moar_cmd_src := exec.Command("curl", "-L", moar_src_url, "-o", "source.tar.gz")
	err = moar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moar_cmd_direct := exec.Command("./binary")
	err = moar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
