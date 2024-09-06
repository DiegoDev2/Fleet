package main

// Yadm - Yet Another Dotfiles Manager
// Homepage: https://yadm.io/

import (
	"fmt"
	
	"os/exec"
)

func installYadm() {
	// Método 1: Descargar y extraer .tar.gz
	yadm_tar_url := "https://github.com/TheLocehiliosan/yadm/archive/refs/tags/3.2.2.tar.gz"
	yadm_cmd_tar := exec.Command("curl", "-L", yadm_tar_url, "-o", "package.tar.gz")
	err := yadm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yadm_zip_url := "https://github.com/TheLocehiliosan/yadm/archive/refs/tags/3.2.2.zip"
	yadm_cmd_zip := exec.Command("curl", "-L", yadm_zip_url, "-o", "package.zip")
	err = yadm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yadm_bin_url := "https://github.com/TheLocehiliosan/yadm/archive/refs/tags/3.2.2.bin"
	yadm_cmd_bin := exec.Command("curl", "-L", yadm_bin_url, "-o", "binary.bin")
	err = yadm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yadm_src_url := "https://github.com/TheLocehiliosan/yadm/archive/refs/tags/3.2.2.src.tar.gz"
	yadm_cmd_src := exec.Command("curl", "-L", yadm_src_url, "-o", "source.tar.gz")
	err = yadm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yadm_cmd_direct := exec.Command("./binary")
	err = yadm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
