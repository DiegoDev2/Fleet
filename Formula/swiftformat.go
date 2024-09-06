package main

// Swiftformat - Formatting tool for reformatting Swift code
// Homepage: https://github.com/nicklockwood/SwiftFormat

import (
	"fmt"
	
	"os/exec"
)

func installSwiftformat() {
	// Método 1: Descargar y extraer .tar.gz
	swiftformat_tar_url := "https://github.com/nicklockwood/SwiftFormat/archive/refs/tags/0.54.3.tar.gz"
	swiftformat_cmd_tar := exec.Command("curl", "-L", swiftformat_tar_url, "-o", "package.tar.gz")
	err := swiftformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftformat_zip_url := "https://github.com/nicklockwood/SwiftFormat/archive/refs/tags/0.54.3.zip"
	swiftformat_cmd_zip := exec.Command("curl", "-L", swiftformat_zip_url, "-o", "package.zip")
	err = swiftformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftformat_bin_url := "https://github.com/nicklockwood/SwiftFormat/archive/refs/tags/0.54.3.bin"
	swiftformat_cmd_bin := exec.Command("curl", "-L", swiftformat_bin_url, "-o", "binary.bin")
	err = swiftformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftformat_src_url := "https://github.com/nicklockwood/SwiftFormat/archive/refs/tags/0.54.3.src.tar.gz"
	swiftformat_cmd_src := exec.Command("curl", "-L", swiftformat_src_url, "-o", "source.tar.gz")
	err = swiftformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftformat_cmd_direct := exec.Command("./binary")
	err = swiftformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
