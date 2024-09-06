package main

// Tdlib - Cross-platform library for building Telegram clients
// Homepage: https://core.telegram.org/tdlib

import (
	"fmt"
	
	"os/exec"
)

func installTdlib() {
	// Método 1: Descargar y extraer .tar.gz
	tdlib_tar_url := "https://github.com/tdlib/td/archive/refs/tags/v1.8.0.tar.gz"
	tdlib_cmd_tar := exec.Command("curl", "-L", tdlib_tar_url, "-o", "package.tar.gz")
	err := tdlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tdlib_zip_url := "https://github.com/tdlib/td/archive/refs/tags/v1.8.0.zip"
	tdlib_cmd_zip := exec.Command("curl", "-L", tdlib_zip_url, "-o", "package.zip")
	err = tdlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tdlib_bin_url := "https://github.com/tdlib/td/archive/refs/tags/v1.8.0.bin"
	tdlib_cmd_bin := exec.Command("curl", "-L", tdlib_bin_url, "-o", "binary.bin")
	err = tdlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tdlib_src_url := "https://github.com/tdlib/td/archive/refs/tags/v1.8.0.src.tar.gz"
	tdlib_cmd_src := exec.Command("curl", "-L", tdlib_src_url, "-o", "source.tar.gz")
	err = tdlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tdlib_cmd_direct := exec.Command("./binary")
	err = tdlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gperf")
	exec.Command("latte", "install", "gperf").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
