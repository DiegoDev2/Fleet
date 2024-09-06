package main

// Pacapt - Package manager in the style of Arch's pacman
// Homepage: https://github.com/icy/pacapt

import (
	"fmt"
	
	"os/exec"
)

func installPacapt() {
	// Método 1: Descargar y extraer .tar.gz
	pacapt_tar_url := "https://github.com/icy/pacapt/archive/refs/tags/v3.0.7.tar.gz"
	pacapt_cmd_tar := exec.Command("curl", "-L", pacapt_tar_url, "-o", "package.tar.gz")
	err := pacapt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pacapt_zip_url := "https://github.com/icy/pacapt/archive/refs/tags/v3.0.7.zip"
	pacapt_cmd_zip := exec.Command("curl", "-L", pacapt_zip_url, "-o", "package.zip")
	err = pacapt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pacapt_bin_url := "https://github.com/icy/pacapt/archive/refs/tags/v3.0.7.bin"
	pacapt_cmd_bin := exec.Command("curl", "-L", pacapt_bin_url, "-o", "binary.bin")
	err = pacapt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pacapt_src_url := "https://github.com/icy/pacapt/archive/refs/tags/v3.0.7.src.tar.gz"
	pacapt_cmd_src := exec.Command("curl", "-L", pacapt_src_url, "-o", "source.tar.gz")
	err = pacapt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pacapt_cmd_direct := exec.Command("./binary")
	err = pacapt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
