package main

// SwiftOutdated - Check for outdated Swift package manager dependencies
// Homepage: https://github.com/kiliankoe/swift-outdated

import (
	"fmt"
	
	"os/exec"
)

func installSwiftOutdated() {
	// Método 1: Descargar y extraer .tar.gz
	swiftoutdated_tar_url := "https://github.com/kiliankoe/swift-outdated/archive/refs/tags/0.9.0.tar.gz"
	swiftoutdated_cmd_tar := exec.Command("curl", "-L", swiftoutdated_tar_url, "-o", "package.tar.gz")
	err := swiftoutdated_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftoutdated_zip_url := "https://github.com/kiliankoe/swift-outdated/archive/refs/tags/0.9.0.zip"
	swiftoutdated_cmd_zip := exec.Command("curl", "-L", swiftoutdated_zip_url, "-o", "package.zip")
	err = swiftoutdated_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftoutdated_bin_url := "https://github.com/kiliankoe/swift-outdated/archive/refs/tags/0.9.0.bin"
	swiftoutdated_cmd_bin := exec.Command("curl", "-L", swiftoutdated_bin_url, "-o", "binary.bin")
	err = swiftoutdated_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftoutdated_src_url := "https://github.com/kiliankoe/swift-outdated/archive/refs/tags/0.9.0.src.tar.gz"
	swiftoutdated_cmd_src := exec.Command("curl", "-L", swiftoutdated_src_url, "-o", "source.tar.gz")
	err = swiftoutdated_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftoutdated_cmd_direct := exec.Command("./binary")
	err = swiftoutdated_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
