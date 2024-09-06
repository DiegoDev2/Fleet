package main

// Hblock - Adblocker that creates a hosts file from multiple sources
// Homepage: https://hblock.molinero.dev/

import (
	"fmt"
	
	"os/exec"
)

func installHblock() {
	// Método 1: Descargar y extraer .tar.gz
	hblock_tar_url := "https://github.com/hectorm/hblock/archive/refs/tags/v3.4.5.tar.gz"
	hblock_cmd_tar := exec.Command("curl", "-L", hblock_tar_url, "-o", "package.tar.gz")
	err := hblock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hblock_zip_url := "https://github.com/hectorm/hblock/archive/refs/tags/v3.4.5.zip"
	hblock_cmd_zip := exec.Command("curl", "-L", hblock_zip_url, "-o", "package.zip")
	err = hblock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hblock_bin_url := "https://github.com/hectorm/hblock/archive/refs/tags/v3.4.5.bin"
	hblock_cmd_bin := exec.Command("curl", "-L", hblock_bin_url, "-o", "binary.bin")
	err = hblock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hblock_src_url := "https://github.com/hectorm/hblock/archive/refs/tags/v3.4.5.src.tar.gz"
	hblock_cmd_src := exec.Command("curl", "-L", hblock_src_url, "-o", "source.tar.gz")
	err = hblock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hblock_cmd_direct := exec.Command("./binary")
	err = hblock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
