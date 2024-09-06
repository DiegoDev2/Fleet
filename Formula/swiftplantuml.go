package main

// Swiftplantuml - Generate UML class diagrams from Swift sources
// Homepage: https://github.com/MarcoEidinger/SwiftPlantUML

import (
	"fmt"
	
	"os/exec"
)

func installSwiftplantuml() {
	// Método 1: Descargar y extraer .tar.gz
	swiftplantuml_tar_url := "https://github.com/MarcoEidinger/SwiftPlantUML/archive/refs/tags/0.8.1.tar.gz"
	swiftplantuml_cmd_tar := exec.Command("curl", "-L", swiftplantuml_tar_url, "-o", "package.tar.gz")
	err := swiftplantuml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftplantuml_zip_url := "https://github.com/MarcoEidinger/SwiftPlantUML/archive/refs/tags/0.8.1.zip"
	swiftplantuml_cmd_zip := exec.Command("curl", "-L", swiftplantuml_zip_url, "-o", "package.zip")
	err = swiftplantuml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftplantuml_bin_url := "https://github.com/MarcoEidinger/SwiftPlantUML/archive/refs/tags/0.8.1.bin"
	swiftplantuml_cmd_bin := exec.Command("curl", "-L", swiftplantuml_bin_url, "-o", "binary.bin")
	err = swiftplantuml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftplantuml_src_url := "https://github.com/MarcoEidinger/SwiftPlantUML/archive/refs/tags/0.8.1.src.tar.gz"
	swiftplantuml_cmd_src := exec.Command("curl", "-L", swiftplantuml_src_url, "-o", "source.tar.gz")
	err = swiftplantuml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftplantuml_cmd_direct := exec.Command("./binary")
	err = swiftplantuml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
