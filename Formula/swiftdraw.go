package main

// Swiftdraw - Convert SVG into PDF, PNG, JPEG or SF Symbol
// Homepage: https://github.com/swhitty/SwiftDraw

import (
	"fmt"
	
	"os/exec"
)

func installSwiftdraw() {
	// Método 1: Descargar y extraer .tar.gz
	swiftdraw_tar_url := "https://github.com/swhitty/SwiftDraw/archive/refs/tags/0.17.0.tar.gz"
	swiftdraw_cmd_tar := exec.Command("curl", "-L", swiftdraw_tar_url, "-o", "package.tar.gz")
	err := swiftdraw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftdraw_zip_url := "https://github.com/swhitty/SwiftDraw/archive/refs/tags/0.17.0.zip"
	swiftdraw_cmd_zip := exec.Command("curl", "-L", swiftdraw_zip_url, "-o", "package.zip")
	err = swiftdraw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftdraw_bin_url := "https://github.com/swhitty/SwiftDraw/archive/refs/tags/0.17.0.bin"
	swiftdraw_cmd_bin := exec.Command("curl", "-L", swiftdraw_bin_url, "-o", "binary.bin")
	err = swiftdraw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftdraw_src_url := "https://github.com/swhitty/SwiftDraw/archive/refs/tags/0.17.0.src.tar.gz"
	swiftdraw_cmd_src := exec.Command("curl", "-L", swiftdraw_src_url, "-o", "source.tar.gz")
	err = swiftdraw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftdraw_cmd_direct := exec.Command("./binary")
	err = swiftdraw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
