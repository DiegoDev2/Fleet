package main

// MacosTermSize - Get the terminal window size on macOS
// Homepage: https://github.com/sindresorhus/macos-term-size

import (
	"fmt"
	
	"os/exec"
)

func installMacosTermSize() {
	// Método 1: Descargar y extraer .tar.gz
	macostermsize_tar_url := "https://github.com/sindresorhus/macos-term-size/archive/refs/tags/v1.0.0.tar.gz"
	macostermsize_cmd_tar := exec.Command("curl", "-L", macostermsize_tar_url, "-o", "package.tar.gz")
	err := macostermsize_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	macostermsize_zip_url := "https://github.com/sindresorhus/macos-term-size/archive/refs/tags/v1.0.0.zip"
	macostermsize_cmd_zip := exec.Command("curl", "-L", macostermsize_zip_url, "-o", "package.zip")
	err = macostermsize_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	macostermsize_bin_url := "https://github.com/sindresorhus/macos-term-size/archive/refs/tags/v1.0.0.bin"
	macostermsize_cmd_bin := exec.Command("curl", "-L", macostermsize_bin_url, "-o", "binary.bin")
	err = macostermsize_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	macostermsize_src_url := "https://github.com/sindresorhus/macos-term-size/archive/refs/tags/v1.0.0.src.tar.gz"
	macostermsize_cmd_src := exec.Command("curl", "-L", macostermsize_src_url, "-o", "source.tar.gz")
	err = macostermsize_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	macostermsize_cmd_direct := exec.Command("./binary")
	err = macostermsize_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
