package main

// Eureka - CLI tool to input and store your ideas without leaving the terminal
// Homepage: https://github.com/simeg/eureka

import (
	"fmt"
	
	"os/exec"
)

func installEureka() {
	// Método 1: Descargar y extraer .tar.gz
	eureka_tar_url := "https://github.com/simeg/eureka/archive/refs/tags/v2.0.0.tar.gz"
	eureka_cmd_tar := exec.Command("curl", "-L", eureka_tar_url, "-o", "package.tar.gz")
	err := eureka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eureka_zip_url := "https://github.com/simeg/eureka/archive/refs/tags/v2.0.0.zip"
	eureka_cmd_zip := exec.Command("curl", "-L", eureka_zip_url, "-o", "package.zip")
	err = eureka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eureka_bin_url := "https://github.com/simeg/eureka/archive/refs/tags/v2.0.0.bin"
	eureka_cmd_bin := exec.Command("curl", "-L", eureka_bin_url, "-o", "binary.bin")
	err = eureka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eureka_src_url := "https://github.com/simeg/eureka/archive/refs/tags/v2.0.0.src.tar.gz"
	eureka_cmd_src := exec.Command("curl", "-L", eureka_src_url, "-o", "source.tar.gz")
	err = eureka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eureka_cmd_direct := exec.Command("./binary")
	err = eureka_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
