package main

// Ttdl - Terminal Todo List Manager
// Homepage: https://github.com/VladimirMarkelov/ttdl

import (
	"fmt"
	
	"os/exec"
)

func installTtdl() {
	// Método 1: Descargar y extraer .tar.gz
	ttdl_tar_url := "https://github.com/VladimirMarkelov/ttdl/archive/refs/tags/v4.4.1.tar.gz"
	ttdl_cmd_tar := exec.Command("curl", "-L", ttdl_tar_url, "-o", "package.tar.gz")
	err := ttdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttdl_zip_url := "https://github.com/VladimirMarkelov/ttdl/archive/refs/tags/v4.4.1.zip"
	ttdl_cmd_zip := exec.Command("curl", "-L", ttdl_zip_url, "-o", "package.zip")
	err = ttdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttdl_bin_url := "https://github.com/VladimirMarkelov/ttdl/archive/refs/tags/v4.4.1.bin"
	ttdl_cmd_bin := exec.Command("curl", "-L", ttdl_bin_url, "-o", "binary.bin")
	err = ttdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttdl_src_url := "https://github.com/VladimirMarkelov/ttdl/archive/refs/tags/v4.4.1.src.tar.gz"
	ttdl_cmd_src := exec.Command("curl", "-L", ttdl_src_url, "-o", "source.tar.gz")
	err = ttdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttdl_cmd_direct := exec.Command("./binary")
	err = ttdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
