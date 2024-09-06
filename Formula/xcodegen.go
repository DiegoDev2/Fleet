package main

// Xcodegen - Generate your Xcode project from a spec file and your folder structure
// Homepage: https://github.com/yonaskolb/XcodeGen

import (
	"fmt"
	
	"os/exec"
)

func installXcodegen() {
	// Método 1: Descargar y extraer .tar.gz
	xcodegen_tar_url := "https://github.com/yonaskolb/XcodeGen/archive/refs/tags/2.42.0.tar.gz"
	xcodegen_cmd_tar := exec.Command("curl", "-L", xcodegen_tar_url, "-o", "package.tar.gz")
	err := xcodegen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcodegen_zip_url := "https://github.com/yonaskolb/XcodeGen/archive/refs/tags/2.42.0.zip"
	xcodegen_cmd_zip := exec.Command("curl", "-L", xcodegen_zip_url, "-o", "package.zip")
	err = xcodegen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcodegen_bin_url := "https://github.com/yonaskolb/XcodeGen/archive/refs/tags/2.42.0.bin"
	xcodegen_cmd_bin := exec.Command("curl", "-L", xcodegen_bin_url, "-o", "binary.bin")
	err = xcodegen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcodegen_src_url := "https://github.com/yonaskolb/XcodeGen/archive/refs/tags/2.42.0.src.tar.gz"
	xcodegen_cmd_src := exec.Command("curl", "-L", xcodegen_src_url, "-o", "source.tar.gz")
	err = xcodegen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcodegen_cmd_direct := exec.Command("./binary")
	err = xcodegen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
