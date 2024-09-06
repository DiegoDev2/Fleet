package main

// Enkits - C and C++ Task Scheduler for creating parallel programs
// Homepage: https://github.com/dougbinks/enkiTS

import (
	"fmt"
	
	"os/exec"
)

func installEnkits() {
	// Método 1: Descargar y extraer .tar.gz
	enkits_tar_url := "https://github.com/dougbinks/enkiTS/archive/refs/tags/v1.11.tar.gz"
	enkits_cmd_tar := exec.Command("curl", "-L", enkits_tar_url, "-o", "package.tar.gz")
	err := enkits_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enkits_zip_url := "https://github.com/dougbinks/enkiTS/archive/refs/tags/v1.11.zip"
	enkits_cmd_zip := exec.Command("curl", "-L", enkits_zip_url, "-o", "package.zip")
	err = enkits_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enkits_bin_url := "https://github.com/dougbinks/enkiTS/archive/refs/tags/v1.11.bin"
	enkits_cmd_bin := exec.Command("curl", "-L", enkits_bin_url, "-o", "binary.bin")
	err = enkits_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enkits_src_url := "https://github.com/dougbinks/enkiTS/archive/refs/tags/v1.11.src.tar.gz"
	enkits_cmd_src := exec.Command("curl", "-L", enkits_src_url, "-o", "source.tar.gz")
	err = enkits_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enkits_cmd_direct := exec.Command("./binary")
	err = enkits_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
