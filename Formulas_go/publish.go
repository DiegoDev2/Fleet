package main

// Publish - Static site generator for Swift developers
// Homepage: https://github.com/JohnSundell/Publish

import (
	"fmt"
	
	"os/exec"
)

func installPublish() {
	// Método 1: Descargar y extraer .tar.gz
	publish_tar_url := "https://github.com/JohnSundell/Publish/archive/refs/tags/0.9.0.tar.gz"
	publish_cmd_tar := exec.Command("curl", "-L", publish_tar_url, "-o", "package.tar.gz")
	err := publish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	publish_zip_url := "https://github.com/JohnSundell/Publish/archive/refs/tags/0.9.0.zip"
	publish_cmd_zip := exec.Command("curl", "-L", publish_zip_url, "-o", "package.zip")
	err = publish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	publish_bin_url := "https://github.com/JohnSundell/Publish/archive/refs/tags/0.9.0.bin"
	publish_cmd_bin := exec.Command("curl", "-L", publish_bin_url, "-o", "binary.bin")
	err = publish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	publish_src_url := "https://github.com/JohnSundell/Publish/archive/refs/tags/0.9.0.src.tar.gz"
	publish_cmd_src := exec.Command("curl", "-L", publish_src_url, "-o", "source.tar.gz")
	err = publish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	publish_cmd_direct := exec.Command("./binary")
	err = publish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
