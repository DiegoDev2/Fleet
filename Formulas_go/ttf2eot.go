package main

// Ttf2eot - Convert TTF files to EOT
// Homepage: https://github.com/wget/ttf2eot

import (
	"fmt"
	
	"os/exec"
)

func installTtf2eot() {
	// Método 1: Descargar y extraer .tar.gz
	ttf2eot_tar_url := "https://github.com/wget/ttf2eot/archive/refs/tags/v0.0.3.tar.gz"
	ttf2eot_cmd_tar := exec.Command("curl", "-L", ttf2eot_tar_url, "-o", "package.tar.gz")
	err := ttf2eot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttf2eot_zip_url := "https://github.com/wget/ttf2eot/archive/refs/tags/v0.0.3.zip"
	ttf2eot_cmd_zip := exec.Command("curl", "-L", ttf2eot_zip_url, "-o", "package.zip")
	err = ttf2eot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttf2eot_bin_url := "https://github.com/wget/ttf2eot/archive/refs/tags/v0.0.3.bin"
	ttf2eot_cmd_bin := exec.Command("curl", "-L", ttf2eot_bin_url, "-o", "binary.bin")
	err = ttf2eot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttf2eot_src_url := "https://github.com/wget/ttf2eot/archive/refs/tags/v0.0.3.src.tar.gz"
	ttf2eot_cmd_src := exec.Command("curl", "-L", ttf2eot_src_url, "-o", "source.tar.gz")
	err = ttf2eot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttf2eot_cmd_direct := exec.Command("./binary")
	err = ttf2eot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
