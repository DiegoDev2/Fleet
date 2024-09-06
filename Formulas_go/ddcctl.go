package main

// Ddcctl - DDC monitor controls (brightness) for Mac OSX command-line
// Homepage: https://github.com/kfix/ddcctl

import (
	"fmt"
	
	"os/exec"
)

func installDdcctl() {
	// Método 1: Descargar y extraer .tar.gz
	ddcctl_tar_url := "https://github.com/kfix/ddcctl/archive/refs/tags/v1.tar.gz"
	ddcctl_cmd_tar := exec.Command("curl", "-L", ddcctl_tar_url, "-o", "package.tar.gz")
	err := ddcctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddcctl_zip_url := "https://github.com/kfix/ddcctl/archive/refs/tags/v1.zip"
	ddcctl_cmd_zip := exec.Command("curl", "-L", ddcctl_zip_url, "-o", "package.zip")
	err = ddcctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddcctl_bin_url := "https://github.com/kfix/ddcctl/archive/refs/tags/v1.bin"
	ddcctl_cmd_bin := exec.Command("curl", "-L", ddcctl_bin_url, "-o", "binary.bin")
	err = ddcctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddcctl_src_url := "https://github.com/kfix/ddcctl/archive/refs/tags/v1.src.tar.gz"
	ddcctl_cmd_src := exec.Command("curl", "-L", ddcctl_src_url, "-o", "source.tar.gz")
	err = ddcctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddcctl_cmd_direct := exec.Command("./binary")
	err = ddcctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
