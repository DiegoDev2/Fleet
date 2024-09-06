package main

// Enca - Charset analyzer and converter
// Homepage: https://cihar.com/software/enca/

import (
	"fmt"
	
	"os/exec"
)

func installEnca() {
	// Método 1: Descargar y extraer .tar.gz
	enca_tar_url := "https://dl.cihar.com/enca/enca-1.19.tar.gz"
	enca_cmd_tar := exec.Command("curl", "-L", enca_tar_url, "-o", "package.tar.gz")
	err := enca_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enca_zip_url := "https://dl.cihar.com/enca/enca-1.19.zip"
	enca_cmd_zip := exec.Command("curl", "-L", enca_zip_url, "-o", "package.zip")
	err = enca_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enca_bin_url := "https://dl.cihar.com/enca/enca-1.19.bin"
	enca_cmd_bin := exec.Command("curl", "-L", enca_bin_url, "-o", "binary.bin")
	err = enca_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enca_src_url := "https://dl.cihar.com/enca/enca-1.19.src.tar.gz"
	enca_cmd_src := exec.Command("curl", "-L", enca_src_url, "-o", "source.tar.gz")
	err = enca_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enca_cmd_direct := exec.Command("./binary")
	err = enca_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
