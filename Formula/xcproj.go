package main

// Xcproj - Manipulate Xcode project files
// Homepage: https://github.com/0xced/xcproj

import (
	"fmt"
	
	"os/exec"
)

func installXcproj() {
	// Método 1: Descargar y extraer .tar.gz
	xcproj_tar_url := "https://github.com/0xced/xcproj/archive/refs/tags/0.2.1.tar.gz"
	xcproj_cmd_tar := exec.Command("curl", "-L", xcproj_tar_url, "-o", "package.tar.gz")
	err := xcproj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcproj_zip_url := "https://github.com/0xced/xcproj/archive/refs/tags/0.2.1.zip"
	xcproj_cmd_zip := exec.Command("curl", "-L", xcproj_zip_url, "-o", "package.zip")
	err = xcproj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcproj_bin_url := "https://github.com/0xced/xcproj/archive/refs/tags/0.2.1.bin"
	xcproj_cmd_bin := exec.Command("curl", "-L", xcproj_bin_url, "-o", "binary.bin")
	err = xcproj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcproj_src_url := "https://github.com/0xced/xcproj/archive/refs/tags/0.2.1.src.tar.gz"
	xcproj_cmd_src := exec.Command("curl", "-L", xcproj_src_url, "-o", "source.tar.gz")
	err = xcproj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcproj_cmd_direct := exec.Command("./binary")
	err = xcproj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
