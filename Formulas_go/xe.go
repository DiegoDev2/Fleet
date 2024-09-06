package main

// Xe - Simple xargs and apply replacement
// Homepage: https://github.com/leahneukirchen/xe

import (
	"fmt"
	
	"os/exec"
)

func installXe() {
	// Método 1: Descargar y extraer .tar.gz
	xe_tar_url := "https://github.com/leahneukirchen/xe/archive/refs/tags/v1.0.tar.gz"
	xe_cmd_tar := exec.Command("curl", "-L", xe_tar_url, "-o", "package.tar.gz")
	err := xe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xe_zip_url := "https://github.com/leahneukirchen/xe/archive/refs/tags/v1.0.zip"
	xe_cmd_zip := exec.Command("curl", "-L", xe_zip_url, "-o", "package.zip")
	err = xe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xe_bin_url := "https://github.com/leahneukirchen/xe/archive/refs/tags/v1.0.bin"
	xe_cmd_bin := exec.Command("curl", "-L", xe_bin_url, "-o", "binary.bin")
	err = xe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xe_src_url := "https://github.com/leahneukirchen/xe/archive/refs/tags/v1.0.src.tar.gz"
	xe_cmd_src := exec.Command("curl", "-L", xe_src_url, "-o", "source.tar.gz")
	err = xe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xe_cmd_direct := exec.Command("./binary")
	err = xe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
