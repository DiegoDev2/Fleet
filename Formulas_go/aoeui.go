package main

// Aoeui - Lightweight text editor optimized for Dvorak and QWERTY keyboards
// Homepage: https://code.google.com/archive/p/aoeui/

import (
	"fmt"
	
	"os/exec"
)

func installAoeui() {
	// Método 1: Descargar y extraer .tar.gz
	aoeui_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/aoeui/aoeui-1.6.tgz"
	aoeui_cmd_tar := exec.Command("curl", "-L", aoeui_tar_url, "-o", "package.tar.gz")
	err := aoeui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aoeui_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/aoeui/aoeui-1.6.tgz"
	aoeui_cmd_zip := exec.Command("curl", "-L", aoeui_zip_url, "-o", "package.zip")
	err = aoeui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aoeui_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/aoeui/aoeui-1.6.tgz"
	aoeui_cmd_bin := exec.Command("curl", "-L", aoeui_bin_url, "-o", "binary.bin")
	err = aoeui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aoeui_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/aoeui/aoeui-1.6.tgz"
	aoeui_cmd_src := exec.Command("curl", "-L", aoeui_src_url, "-o", "source.tar.gz")
	err = aoeui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aoeui_cmd_direct := exec.Command("./binary")
	err = aoeui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
