package main

// Fribidi - Implementation of the Unicode BiDi algorithm
// Homepage: https://github.com/fribidi/fribidi

import (
	"fmt"
	
	"os/exec"
)

func installFribidi() {
	// Método 1: Descargar y extraer .tar.gz
	fribidi_tar_url := "https://github.com/fribidi/fribidi/releases/download/v1.0.15/fribidi-1.0.15.tar.xz"
	fribidi_cmd_tar := exec.Command("curl", "-L", fribidi_tar_url, "-o", "package.tar.gz")
	err := fribidi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fribidi_zip_url := "https://github.com/fribidi/fribidi/releases/download/v1.0.15/fribidi-1.0.15.tar.xz"
	fribidi_cmd_zip := exec.Command("curl", "-L", fribidi_zip_url, "-o", "package.zip")
	err = fribidi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fribidi_bin_url := "https://github.com/fribidi/fribidi/releases/download/v1.0.15/fribidi-1.0.15.tar.xz"
	fribidi_cmd_bin := exec.Command("curl", "-L", fribidi_bin_url, "-o", "binary.bin")
	err = fribidi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fribidi_src_url := "https://github.com/fribidi/fribidi/releases/download/v1.0.15/fribidi-1.0.15.tar.xz"
	fribidi_cmd_src := exec.Command("curl", "-L", fribidi_src_url, "-o", "source.tar.gz")
	err = fribidi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fribidi_cmd_direct := exec.Command("./binary")
	err = fribidi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
