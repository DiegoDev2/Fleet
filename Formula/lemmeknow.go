package main

// Lemmeknow - Fastest way to identify anything!
// Homepage: https://github.com/swanandx/lemmeknow

import (
	"fmt"
	
	"os/exec"
)

func installLemmeknow() {
	// Método 1: Descargar y extraer .tar.gz
	lemmeknow_tar_url := "https://github.com/swanandx/lemmeknow/archive/refs/tags/v0.8.0.tar.gz"
	lemmeknow_cmd_tar := exec.Command("curl", "-L", lemmeknow_tar_url, "-o", "package.tar.gz")
	err := lemmeknow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lemmeknow_zip_url := "https://github.com/swanandx/lemmeknow/archive/refs/tags/v0.8.0.zip"
	lemmeknow_cmd_zip := exec.Command("curl", "-L", lemmeknow_zip_url, "-o", "package.zip")
	err = lemmeknow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lemmeknow_bin_url := "https://github.com/swanandx/lemmeknow/archive/refs/tags/v0.8.0.bin"
	lemmeknow_cmd_bin := exec.Command("curl", "-L", lemmeknow_bin_url, "-o", "binary.bin")
	err = lemmeknow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lemmeknow_src_url := "https://github.com/swanandx/lemmeknow/archive/refs/tags/v0.8.0.src.tar.gz"
	lemmeknow_cmd_src := exec.Command("curl", "-L", lemmeknow_src_url, "-o", "source.tar.gz")
	err = lemmeknow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lemmeknow_cmd_direct := exec.Command("./binary")
	err = lemmeknow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
