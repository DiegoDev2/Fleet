package main

// V - Z for vim
// Homepage: https://github.com/rupa/v

import (
	"fmt"
	
	"os/exec"
)

func installV() {
	// Método 1: Descargar y extraer .tar.gz
	v_tar_url := "https://github.com/rupa/v/archive/refs/tags/v1.1.tar.gz"
	v_cmd_tar := exec.Command("curl", "-L", v_tar_url, "-o", "package.tar.gz")
	err := v_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	v_zip_url := "https://github.com/rupa/v/archive/refs/tags/v1.1.zip"
	v_cmd_zip := exec.Command("curl", "-L", v_zip_url, "-o", "package.zip")
	err = v_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	v_bin_url := "https://github.com/rupa/v/archive/refs/tags/v1.1.bin"
	v_cmd_bin := exec.Command("curl", "-L", v_bin_url, "-o", "binary.bin")
	err = v_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	v_src_url := "https://github.com/rupa/v/archive/refs/tags/v1.1.src.tar.gz"
	v_cmd_src := exec.Command("curl", "-L", v_src_url, "-o", "source.tar.gz")
	err = v_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	v_cmd_direct := exec.Command("./binary")
	err = v_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
