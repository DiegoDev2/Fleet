package main

// Vf - Enhanced version of `cd` command
// Homepage: https://github.com/glejeune/vf

import (
	"fmt"
	
	"os/exec"
)

func installVf() {
	// Método 1: Descargar y extraer .tar.gz
	vf_tar_url := "https://github.com/glejeune/vf/archive/refs/tags/0.0.1.tar.gz"
	vf_cmd_tar := exec.Command("curl", "-L", vf_tar_url, "-o", "package.tar.gz")
	err := vf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vf_zip_url := "https://github.com/glejeune/vf/archive/refs/tags/0.0.1.zip"
	vf_cmd_zip := exec.Command("curl", "-L", vf_zip_url, "-o", "package.zip")
	err = vf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vf_bin_url := "https://github.com/glejeune/vf/archive/refs/tags/0.0.1.bin"
	vf_cmd_bin := exec.Command("curl", "-L", vf_bin_url, "-o", "binary.bin")
	err = vf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vf_src_url := "https://github.com/glejeune/vf/archive/refs/tags/0.0.1.src.tar.gz"
	vf_cmd_src := exec.Command("curl", "-L", vf_src_url, "-o", "source.tar.gz")
	err = vf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vf_cmd_direct := exec.Command("./binary")
	err = vf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
