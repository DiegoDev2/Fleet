package main

// Inja - Template engine for modern C++
// Homepage: https://pantor.github.io/inja/

import (
	"fmt"
	
	"os/exec"
)

func installInja() {
	// Método 1: Descargar y extraer .tar.gz
	inja_tar_url := "https://github.com/pantor/inja/archive/refs/tags/v3.4.0.tar.gz"
	inja_cmd_tar := exec.Command("curl", "-L", inja_tar_url, "-o", "package.tar.gz")
	err := inja_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inja_zip_url := "https://github.com/pantor/inja/archive/refs/tags/v3.4.0.zip"
	inja_cmd_zip := exec.Command("curl", "-L", inja_zip_url, "-o", "package.zip")
	err = inja_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inja_bin_url := "https://github.com/pantor/inja/archive/refs/tags/v3.4.0.bin"
	inja_cmd_bin := exec.Command("curl", "-L", inja_bin_url, "-o", "binary.bin")
	err = inja_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inja_src_url := "https://github.com/pantor/inja/archive/refs/tags/v3.4.0.src.tar.gz"
	inja_cmd_src := exec.Command("curl", "-L", inja_src_url, "-o", "source.tar.gz")
	err = inja_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inja_cmd_direct := exec.Command("./binary")
	err = inja_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
}
