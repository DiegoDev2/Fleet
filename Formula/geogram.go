package main

// Geogram - Programming library of geometric algorithms
// Homepage: https://brunolevy.github.io/geogram/

import (
	"fmt"
	
	"os/exec"
)

func installGeogram() {
	// Método 1: Descargar y extraer .tar.gz
	geogram_tar_url := "https://github.com/BrunoLevy/geogram/releases/download/v1.9.0/geogram_1.9.0.tar.gz"
	geogram_cmd_tar := exec.Command("curl", "-L", geogram_tar_url, "-o", "package.tar.gz")
	err := geogram_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geogram_zip_url := "https://github.com/BrunoLevy/geogram/releases/download/v1.9.0/geogram_1.9.0.zip"
	geogram_cmd_zip := exec.Command("curl", "-L", geogram_zip_url, "-o", "package.zip")
	err = geogram_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geogram_bin_url := "https://github.com/BrunoLevy/geogram/releases/download/v1.9.0/geogram_1.9.0.bin"
	geogram_cmd_bin := exec.Command("curl", "-L", geogram_bin_url, "-o", "binary.bin")
	err = geogram_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geogram_src_url := "https://github.com/BrunoLevy/geogram/releases/download/v1.9.0/geogram_1.9.0.src.tar.gz"
	geogram_cmd_src := exec.Command("curl", "-L", geogram_src_url, "-o", "source.tar.gz")
	err = geogram_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geogram_cmd_direct := exec.Command("./binary")
	err = geogram_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
}
