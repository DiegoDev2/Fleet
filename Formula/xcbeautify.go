package main

// Xcbeautify - Little beautifier tool for xcodebuild
// Homepage: https://github.com/cpisciotta/xcbeautify

import (
	"fmt"
	
	"os/exec"
)

func installXcbeautify() {
	// Método 1: Descargar y extraer .tar.gz
	xcbeautify_tar_url := "https://github.com/cpisciotta/xcbeautify/archive/refs/tags/2.11.0.tar.gz"
	xcbeautify_cmd_tar := exec.Command("curl", "-L", xcbeautify_tar_url, "-o", "package.tar.gz")
	err := xcbeautify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbeautify_zip_url := "https://github.com/cpisciotta/xcbeautify/archive/refs/tags/2.11.0.zip"
	xcbeautify_cmd_zip := exec.Command("curl", "-L", xcbeautify_zip_url, "-o", "package.zip")
	err = xcbeautify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbeautify_bin_url := "https://github.com/cpisciotta/xcbeautify/archive/refs/tags/2.11.0.bin"
	xcbeautify_cmd_bin := exec.Command("curl", "-L", xcbeautify_bin_url, "-o", "binary.bin")
	err = xcbeautify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbeautify_src_url := "https://github.com/cpisciotta/xcbeautify/archive/refs/tags/2.11.0.src.tar.gz"
	xcbeautify_cmd_src := exec.Command("curl", "-L", xcbeautify_src_url, "-o", "source.tar.gz")
	err = xcbeautify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbeautify_cmd_direct := exec.Command("./binary")
	err = xcbeautify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
