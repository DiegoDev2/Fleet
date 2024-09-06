package main

// Swiftgen - Swift code generator for assets, storyboards, Localizable.strings, etc.
// Homepage: https://github.com/SwiftGen/SwiftGen

import (
	"fmt"
	
	"os/exec"
)

func installSwiftgen() {
	// Método 1: Descargar y extraer .tar.gz
	swiftgen_tar_url := "https://github.com/SwiftGen/SwiftGen/archive/refs/tags/6.6.3.tar.gz"
	swiftgen_cmd_tar := exec.Command("curl", "-L", swiftgen_tar_url, "-o", "package.tar.gz")
	err := swiftgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftgen_zip_url := "https://github.com/SwiftGen/SwiftGen/archive/refs/tags/6.6.3.zip"
	swiftgen_cmd_zip := exec.Command("curl", "-L", swiftgen_zip_url, "-o", "package.zip")
	err = swiftgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftgen_bin_url := "https://github.com/SwiftGen/SwiftGen/archive/refs/tags/6.6.3.bin"
	swiftgen_cmd_bin := exec.Command("curl", "-L", swiftgen_bin_url, "-o", "binary.bin")
	err = swiftgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftgen_src_url := "https://github.com/SwiftGen/SwiftGen/archive/refs/tags/6.6.3.src.tar.gz"
	swiftgen_cmd_src := exec.Command("curl", "-L", swiftgen_src_url, "-o", "source.tar.gz")
	err = swiftgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftgen_cmd_direct := exec.Command("./binary")
	err = swiftgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
