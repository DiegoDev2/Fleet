package main

// Bear - Generate compilation database for clang tooling
// Homepage: https://github.com/rizsotto/Bear

import (
	"fmt"
	
	"os/exec"
)

func installBear() {
	// Método 1: Descargar y extraer .tar.gz
	bear_tar_url := "https://github.com/rizsotto/Bear/archive/refs/tags/3.1.4.tar.gz"
	bear_cmd_tar := exec.Command("curl", "-L", bear_tar_url, "-o", "package.tar.gz")
	err := bear_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bear_zip_url := "https://github.com/rizsotto/Bear/archive/refs/tags/3.1.4.zip"
	bear_cmd_zip := exec.Command("curl", "-L", bear_zip_url, "-o", "package.zip")
	err = bear_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bear_bin_url := "https://github.com/rizsotto/Bear/archive/refs/tags/3.1.4.bin"
	bear_cmd_bin := exec.Command("curl", "-L", bear_bin_url, "-o", "binary.bin")
	err = bear_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bear_src_url := "https://github.com/rizsotto/Bear/archive/refs/tags/3.1.4.src.tar.gz"
	bear_cmd_src := exec.Command("curl", "-L", bear_src_url, "-o", "source.tar.gz")
	err = bear_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bear_cmd_direct := exec.Command("./binary")
	err = bear_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: grpc")
	exec.Command("latte", "install", "grpc").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: spdlog")
	exec.Command("latte", "install", "spdlog").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
