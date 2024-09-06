package main

// XcodeBuildServer - Build server protocol implementation for integrating Xcode with sourcekit-lsp
// Homepage: https://github.com/SolaWing/xcode-build-server

import (
	"fmt"
	
	"os/exec"
)

func installXcodeBuildServer() {
	// Método 1: Descargar y extraer .tar.gz
	xcodebuildserver_tar_url := "https://github.com/SolaWing/xcode-build-server/archive/refs/tags/v1.1.0.tar.gz"
	xcodebuildserver_cmd_tar := exec.Command("curl", "-L", xcodebuildserver_tar_url, "-o", "package.tar.gz")
	err := xcodebuildserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcodebuildserver_zip_url := "https://github.com/SolaWing/xcode-build-server/archive/refs/tags/v1.1.0.zip"
	xcodebuildserver_cmd_zip := exec.Command("curl", "-L", xcodebuildserver_zip_url, "-o", "package.zip")
	err = xcodebuildserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcodebuildserver_bin_url := "https://github.com/SolaWing/xcode-build-server/archive/refs/tags/v1.1.0.bin"
	xcodebuildserver_cmd_bin := exec.Command("curl", "-L", xcodebuildserver_bin_url, "-o", "binary.bin")
	err = xcodebuildserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcodebuildserver_src_url := "https://github.com/SolaWing/xcode-build-server/archive/refs/tags/v1.1.0.src.tar.gz"
	xcodebuildserver_cmd_src := exec.Command("curl", "-L", xcodebuildserver_src_url, "-o", "source.tar.gz")
	err = xcodebuildserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcodebuildserver_cmd_direct := exec.Command("./binary")
	err = xcodebuildserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gzip")
	exec.Command("latte", "install", "gzip").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
