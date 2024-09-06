package main

// SwiftSh - Scripting with easy zero-conf dependency imports
// Homepage: https://github.com/mxcl/swift-sh

import (
	"fmt"
	
	"os/exec"
)

func installSwiftSh() {
	// Método 1: Descargar y extraer .tar.gz
	swiftsh_tar_url := "https://github.com/mxcl/swift-sh/archive/refs/tags/2.4.0.tar.gz"
	swiftsh_cmd_tar := exec.Command("curl", "-L", swiftsh_tar_url, "-o", "package.tar.gz")
	err := swiftsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftsh_zip_url := "https://github.com/mxcl/swift-sh/archive/refs/tags/2.4.0.zip"
	swiftsh_cmd_zip := exec.Command("curl", "-L", swiftsh_zip_url, "-o", "package.zip")
	err = swiftsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftsh_bin_url := "https://github.com/mxcl/swift-sh/archive/refs/tags/2.4.0.bin"
	swiftsh_cmd_bin := exec.Command("curl", "-L", swiftsh_bin_url, "-o", "binary.bin")
	err = swiftsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftsh_src_url := "https://github.com/mxcl/swift-sh/archive/refs/tags/2.4.0.src.tar.gz"
	swiftsh_cmd_src := exec.Command("curl", "-L", swiftsh_src_url, "-o", "source.tar.gz")
	err = swiftsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftsh_cmd_direct := exec.Command("./binary")
	err = swiftsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
