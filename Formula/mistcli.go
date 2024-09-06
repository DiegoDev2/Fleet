package main

// MistCli - Mac command-line tool that automatically downloads macOS Firmwares / Installers
// Homepage: https://github.com/ninxsoft/mist-cli

import (
	"fmt"
	
	"os/exec"
)

func installMistCli() {
	// Método 1: Descargar y extraer .tar.gz
	mistcli_tar_url := "https://github.com/ninxsoft/mist-cli/archive/refs/tags/v2.1.1.tar.gz"
	mistcli_cmd_tar := exec.Command("curl", "-L", mistcli_tar_url, "-o", "package.tar.gz")
	err := mistcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mistcli_zip_url := "https://github.com/ninxsoft/mist-cli/archive/refs/tags/v2.1.1.zip"
	mistcli_cmd_zip := exec.Command("curl", "-L", mistcli_zip_url, "-o", "package.zip")
	err = mistcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mistcli_bin_url := "https://github.com/ninxsoft/mist-cli/archive/refs/tags/v2.1.1.bin"
	mistcli_cmd_bin := exec.Command("curl", "-L", mistcli_bin_url, "-o", "binary.bin")
	err = mistcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mistcli_src_url := "https://github.com/ninxsoft/mist-cli/archive/refs/tags/v2.1.1.src.tar.gz"
	mistcli_cmd_src := exec.Command("curl", "-L", mistcli_src_url, "-o", "source.tar.gz")
	err = mistcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mistcli_cmd_direct := exec.Command("./binary")
	err = mistcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
