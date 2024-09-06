package main

// Blueutil - Get/set bluetooth power and discoverable state
// Homepage: https://github.com/toy/blueutil

import (
	"fmt"
	
	"os/exec"
)

func installBlueutil() {
	// Método 1: Descargar y extraer .tar.gz
	blueutil_tar_url := "https://github.com/toy/blueutil/archive/refs/tags/v2.10.0.tar.gz"
	blueutil_cmd_tar := exec.Command("curl", "-L", blueutil_tar_url, "-o", "package.tar.gz")
	err := blueutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blueutil_zip_url := "https://github.com/toy/blueutil/archive/refs/tags/v2.10.0.zip"
	blueutil_cmd_zip := exec.Command("curl", "-L", blueutil_zip_url, "-o", "package.zip")
	err = blueutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blueutil_bin_url := "https://github.com/toy/blueutil/archive/refs/tags/v2.10.0.bin"
	blueutil_cmd_bin := exec.Command("curl", "-L", blueutil_bin_url, "-o", "binary.bin")
	err = blueutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blueutil_src_url := "https://github.com/toy/blueutil/archive/refs/tags/v2.10.0.src.tar.gz"
	blueutil_cmd_src := exec.Command("curl", "-L", blueutil_src_url, "-o", "source.tar.gz")
	err = blueutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blueutil_cmd_direct := exec.Command("./binary")
	err = blueutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
