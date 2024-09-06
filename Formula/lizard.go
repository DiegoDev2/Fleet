package main

// Lizard - Efficient compressor with very fast decompression
// Homepage: https://github.com/inikep/lizard

import (
	"fmt"
	
	"os/exec"
)

func installLizard() {
	// Método 1: Descargar y extraer .tar.gz
	lizard_tar_url := "https://github.com/inikep/lizard/archive/refs/tags/v1.0.tar.gz"
	lizard_cmd_tar := exec.Command("curl", "-L", lizard_tar_url, "-o", "package.tar.gz")
	err := lizard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lizard_zip_url := "https://github.com/inikep/lizard/archive/refs/tags/v1.0.zip"
	lizard_cmd_zip := exec.Command("curl", "-L", lizard_zip_url, "-o", "package.zip")
	err = lizard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lizard_bin_url := "https://github.com/inikep/lizard/archive/refs/tags/v1.0.bin"
	lizard_cmd_bin := exec.Command("curl", "-L", lizard_bin_url, "-o", "binary.bin")
	err = lizard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lizard_src_url := "https://github.com/inikep/lizard/archive/refs/tags/v1.0.src.tar.gz"
	lizard_cmd_src := exec.Command("curl", "-L", lizard_src_url, "-o", "source.tar.gz")
	err = lizard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lizard_cmd_direct := exec.Command("./binary")
	err = lizard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
