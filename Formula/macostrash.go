package main

// MacosTrash - Move files and folders to the trash
// Homepage: https://github.com/sindresorhus/macos-trash

import (
	"fmt"
	
	"os/exec"
)

func installMacosTrash() {
	// Método 1: Descargar y extraer .tar.gz
	macostrash_tar_url := "https://github.com/sindresorhus/macos-trash/archive/refs/tags/v1.2.0.tar.gz"
	macostrash_cmd_tar := exec.Command("curl", "-L", macostrash_tar_url, "-o", "package.tar.gz")
	err := macostrash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	macostrash_zip_url := "https://github.com/sindresorhus/macos-trash/archive/refs/tags/v1.2.0.zip"
	macostrash_cmd_zip := exec.Command("curl", "-L", macostrash_zip_url, "-o", "package.zip")
	err = macostrash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	macostrash_bin_url := "https://github.com/sindresorhus/macos-trash/archive/refs/tags/v1.2.0.bin"
	macostrash_cmd_bin := exec.Command("curl", "-L", macostrash_bin_url, "-o", "binary.bin")
	err = macostrash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	macostrash_src_url := "https://github.com/sindresorhus/macos-trash/archive/refs/tags/v1.2.0.src.tar.gz"
	macostrash_cmd_src := exec.Command("curl", "-L", macostrash_src_url, "-o", "source.tar.gz")
	err = macostrash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	macostrash_cmd_direct := exec.Command("./binary")
	err = macostrash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
