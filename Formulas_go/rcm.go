package main

// Rcm - RC file (dotfile) management
// Homepage: https://thoughtbot.github.io/rcm/rcm.7.html

import (
	"fmt"
	
	"os/exec"
)

func installRcm() {
	// Método 1: Descargar y extraer .tar.gz
	rcm_tar_url := "https://thoughtbot.github.io/rcm/dist/rcm-1.3.6.tar.gz"
	rcm_cmd_tar := exec.Command("curl", "-L", rcm_tar_url, "-o", "package.tar.gz")
	err := rcm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rcm_zip_url := "https://thoughtbot.github.io/rcm/dist/rcm-1.3.6.zip"
	rcm_cmd_zip := exec.Command("curl", "-L", rcm_zip_url, "-o", "package.zip")
	err = rcm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rcm_bin_url := "https://thoughtbot.github.io/rcm/dist/rcm-1.3.6.bin"
	rcm_cmd_bin := exec.Command("curl", "-L", rcm_bin_url, "-o", "binary.bin")
	err = rcm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rcm_src_url := "https://thoughtbot.github.io/rcm/dist/rcm-1.3.6.src.tar.gz"
	rcm_cmd_src := exec.Command("curl", "-L", rcm_src_url, "-o", "source.tar.gz")
	err = rcm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rcm_cmd_direct := exec.Command("./binary")
	err = rcm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
