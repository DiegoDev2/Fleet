package main

// Cc65 - 6502 C compiler
// Homepage: https://cc65.github.io/cc65/

import (
	"fmt"
	
	"os/exec"
)

func installCc65() {
	// Método 1: Descargar y extraer .tar.gz
	cc65_tar_url := "https://github.com/cc65/cc65/archive/refs/tags/V2.19.tar.gz"
	cc65_cmd_tar := exec.Command("curl", "-L", cc65_tar_url, "-o", "package.tar.gz")
	err := cc65_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cc65_zip_url := "https://github.com/cc65/cc65/archive/refs/tags/V2.19.zip"
	cc65_cmd_zip := exec.Command("curl", "-L", cc65_zip_url, "-o", "package.zip")
	err = cc65_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cc65_bin_url := "https://github.com/cc65/cc65/archive/refs/tags/V2.19.bin"
	cc65_cmd_bin := exec.Command("curl", "-L", cc65_bin_url, "-o", "binary.bin")
	err = cc65_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cc65_src_url := "https://github.com/cc65/cc65/archive/refs/tags/V2.19.src.tar.gz"
	cc65_cmd_src := exec.Command("curl", "-L", cc65_src_url, "-o", "source.tar.gz")
	err = cc65_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cc65_cmd_direct := exec.Command("./binary")
	err = cc65_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
