package main

// Hyx - Powerful hex editor for the console
// Homepage: https://yx7.cc/code/

import (
	"fmt"
	
	"os/exec"
)

func installHyx() {
	// Método 1: Descargar y extraer .tar.gz
	hyx_tar_url := "https://yx7.cc/code/hyx/hyx-2024.02.29.tar.xz"
	hyx_cmd_tar := exec.Command("curl", "-L", hyx_tar_url, "-o", "package.tar.gz")
	err := hyx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyx_zip_url := "https://yx7.cc/code/hyx/hyx-2024.02.29.tar.xz"
	hyx_cmd_zip := exec.Command("curl", "-L", hyx_zip_url, "-o", "package.zip")
	err = hyx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyx_bin_url := "https://yx7.cc/code/hyx/hyx-2024.02.29.tar.xz"
	hyx_cmd_bin := exec.Command("curl", "-L", hyx_bin_url, "-o", "binary.bin")
	err = hyx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyx_src_url := "https://yx7.cc/code/hyx/hyx-2024.02.29.tar.xz"
	hyx_cmd_src := exec.Command("curl", "-L", hyx_src_url, "-o", "source.tar.gz")
	err = hyx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyx_cmd_direct := exec.Command("./binary")
	err = hyx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
