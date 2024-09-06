package main

// Crow - Fast and Easy to use microframework for the web
// Homepage: https://crowcpp.org

import (
	"fmt"
	
	"os/exec"
)

func installCrow() {
	// Método 1: Descargar y extraer .tar.gz
	crow_tar_url := "https://github.com/CrowCpp/Crow/archive/refs/tags/v1.2.0.tar.gz"
	crow_cmd_tar := exec.Command("curl", "-L", crow_tar_url, "-o", "package.tar.gz")
	err := crow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crow_zip_url := "https://github.com/CrowCpp/Crow/archive/refs/tags/v1.2.0.zip"
	crow_cmd_zip := exec.Command("curl", "-L", crow_zip_url, "-o", "package.zip")
	err = crow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crow_bin_url := "https://github.com/CrowCpp/Crow/archive/refs/tags/v1.2.0.bin"
	crow_cmd_bin := exec.Command("curl", "-L", crow_bin_url, "-o", "binary.bin")
	err = crow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crow_src_url := "https://github.com/CrowCpp/Crow/archive/refs/tags/v1.2.0.src.tar.gz"
	crow_cmd_src := exec.Command("curl", "-L", crow_src_url, "-o", "source.tar.gz")
	err = crow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crow_cmd_direct := exec.Command("./binary")
	err = crow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: asio")
exec.Command("latte", "install", "asio")
}
