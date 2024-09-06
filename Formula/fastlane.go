package main

// Fastlane - Easiest way to build and release mobile apps
// Homepage: https://fastlane.tools

import (
	"fmt"
	
	"os/exec"
)

func installFastlane() {
	// Método 1: Descargar y extraer .tar.gz
	fastlane_tar_url := "https://github.com/fastlane/fastlane/archive/refs/tags/2.222.0.tar.gz"
	fastlane_cmd_tar := exec.Command("curl", "-L", fastlane_tar_url, "-o", "package.tar.gz")
	err := fastlane_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastlane_zip_url := "https://github.com/fastlane/fastlane/archive/refs/tags/2.222.0.zip"
	fastlane_cmd_zip := exec.Command("curl", "-L", fastlane_zip_url, "-o", "package.zip")
	err = fastlane_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastlane_bin_url := "https://github.com/fastlane/fastlane/archive/refs/tags/2.222.0.bin"
	fastlane_cmd_bin := exec.Command("curl", "-L", fastlane_bin_url, "-o", "binary.bin")
	err = fastlane_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastlane_src_url := "https://github.com/fastlane/fastlane/archive/refs/tags/2.222.0.src.tar.gz"
	fastlane_cmd_src := exec.Command("curl", "-L", fastlane_src_url, "-o", "source.tar.gz")
	err = fastlane_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastlane_cmd_direct := exec.Command("./binary")
	err = fastlane_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
}
