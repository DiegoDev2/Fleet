package main

// Spdlog - Super fast C++ logging library
// Homepage: https://github.com/gabime/spdlog

import (
	"fmt"
	
	"os/exec"
)

func installSpdlog() {
	// Método 1: Descargar y extraer .tar.gz
	spdlog_tar_url := "https://github.com/gabime/spdlog/archive/refs/tags/v1.14.1.tar.gz"
	spdlog_cmd_tar := exec.Command("curl", "-L", spdlog_tar_url, "-o", "package.tar.gz")
	err := spdlog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spdlog_zip_url := "https://github.com/gabime/spdlog/archive/refs/tags/v1.14.1.zip"
	spdlog_cmd_zip := exec.Command("curl", "-L", spdlog_zip_url, "-o", "package.zip")
	err = spdlog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spdlog_bin_url := "https://github.com/gabime/spdlog/archive/refs/tags/v1.14.1.bin"
	spdlog_cmd_bin := exec.Command("curl", "-L", spdlog_bin_url, "-o", "binary.bin")
	err = spdlog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spdlog_src_url := "https://github.com/gabime/spdlog/archive/refs/tags/v1.14.1.src.tar.gz"
	spdlog_cmd_src := exec.Command("curl", "-L", spdlog_src_url, "-o", "source.tar.gz")
	err = spdlog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spdlog_cmd_direct := exec.Command("./binary")
	err = spdlog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
}
