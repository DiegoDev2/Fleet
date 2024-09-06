package main

// Nef - Steroids for Xcode Playgrounds
// Homepage: https://nef.bow-swift.io

import (
	"fmt"
	
	"os/exec"
)

func installNef() {
	// Método 1: Descargar y extraer .tar.gz
	nef_tar_url := "https://github.com/bow-swift/nef/archive/refs/tags/0.7.1.tar.gz"
	nef_cmd_tar := exec.Command("curl", "-L", nef_tar_url, "-o", "package.tar.gz")
	err := nef_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nef_zip_url := "https://github.com/bow-swift/nef/archive/refs/tags/0.7.1.zip"
	nef_cmd_zip := exec.Command("curl", "-L", nef_zip_url, "-o", "package.zip")
	err = nef_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nef_bin_url := "https://github.com/bow-swift/nef/archive/refs/tags/0.7.1.bin"
	nef_cmd_bin := exec.Command("curl", "-L", nef_bin_url, "-o", "binary.bin")
	err = nef_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nef_src_url := "https://github.com/bow-swift/nef/archive/refs/tags/0.7.1.src.tar.gz"
	nef_cmd_src := exec.Command("curl", "-L", nef_src_url, "-o", "source.tar.gz")
	err = nef_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nef_cmd_direct := exec.Command("./binary")
	err = nef_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
