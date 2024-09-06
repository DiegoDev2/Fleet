package main

// Semtag - Semantic tagging script for git
// Homepage: https://github.com/nico2sh/semtag

import (
	"fmt"
	
	"os/exec"
)

func installSemtag() {
	// Método 1: Descargar y extraer .tar.gz
	semtag_tar_url := "https://github.com/nico2sh/semtag/archive/refs/tags/v0.1.1.tar.gz"
	semtag_cmd_tar := exec.Command("curl", "-L", semtag_tar_url, "-o", "package.tar.gz")
	err := semtag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	semtag_zip_url := "https://github.com/nico2sh/semtag/archive/refs/tags/v0.1.1.zip"
	semtag_cmd_zip := exec.Command("curl", "-L", semtag_zip_url, "-o", "package.zip")
	err = semtag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	semtag_bin_url := "https://github.com/nico2sh/semtag/archive/refs/tags/v0.1.1.bin"
	semtag_cmd_bin := exec.Command("curl", "-L", semtag_bin_url, "-o", "binary.bin")
	err = semtag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	semtag_src_url := "https://github.com/nico2sh/semtag/archive/refs/tags/v0.1.1.src.tar.gz"
	semtag_cmd_src := exec.Command("curl", "-L", semtag_src_url, "-o", "source.tar.gz")
	err = semtag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	semtag_cmd_direct := exec.Command("./binary")
	err = semtag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
