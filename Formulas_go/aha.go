package main

// Aha - ANSI HTML adapter
// Homepage: https://github.com/theZiz/aha

import (
	"fmt"
	
	"os/exec"
)

func installAha() {
	// Método 1: Descargar y extraer .tar.gz
	aha_tar_url := "https://github.com/theZiz/aha/archive/refs/tags/0.5.1.tar.gz"
	aha_cmd_tar := exec.Command("curl", "-L", aha_tar_url, "-o", "package.tar.gz")
	err := aha_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aha_zip_url := "https://github.com/theZiz/aha/archive/refs/tags/0.5.1.zip"
	aha_cmd_zip := exec.Command("curl", "-L", aha_zip_url, "-o", "package.zip")
	err = aha_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aha_bin_url := "https://github.com/theZiz/aha/archive/refs/tags/0.5.1.bin"
	aha_cmd_bin := exec.Command("curl", "-L", aha_bin_url, "-o", "binary.bin")
	err = aha_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aha_src_url := "https://github.com/theZiz/aha/archive/refs/tags/0.5.1.src.tar.gz"
	aha_cmd_src := exec.Command("curl", "-L", aha_src_url, "-o", "source.tar.gz")
	err = aha_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aha_cmd_direct := exec.Command("./binary")
	err = aha_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
