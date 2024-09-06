package main

// Kommit - More detailed commit messages without committing!
// Homepage: https://github.com/vigo/kommit

import (
	"fmt"
	
	"os/exec"
)

func installKommit() {
	// Método 1: Descargar y extraer .tar.gz
	kommit_tar_url := "https://github.com/vigo/kommit/archive/refs/tags/v1.1.0.tar.gz"
	kommit_cmd_tar := exec.Command("curl", "-L", kommit_tar_url, "-o", "package.tar.gz")
	err := kommit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kommit_zip_url := "https://github.com/vigo/kommit/archive/refs/tags/v1.1.0.zip"
	kommit_cmd_zip := exec.Command("curl", "-L", kommit_zip_url, "-o", "package.zip")
	err = kommit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kommit_bin_url := "https://github.com/vigo/kommit/archive/refs/tags/v1.1.0.bin"
	kommit_cmd_bin := exec.Command("curl", "-L", kommit_bin_url, "-o", "binary.bin")
	err = kommit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kommit_src_url := "https://github.com/vigo/kommit/archive/refs/tags/v1.1.0.src.tar.gz"
	kommit_cmd_src := exec.Command("curl", "-L", kommit_src_url, "-o", "source.tar.gz")
	err = kommit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kommit_cmd_direct := exec.Command("./binary")
	err = kommit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
