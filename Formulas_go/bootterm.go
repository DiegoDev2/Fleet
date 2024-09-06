package main

// Bootterm - Simple, reliable and powerful terminal to ease connection to serial ports
// Homepage: https://github.com/wtarreau/bootterm

import (
	"fmt"
	
	"os/exec"
)

func installBootterm() {
	// Método 1: Descargar y extraer .tar.gz
	bootterm_tar_url := "https://github.com/wtarreau/bootterm/archive/refs/tags/v0.5.tar.gz"
	bootterm_cmd_tar := exec.Command("curl", "-L", bootterm_tar_url, "-o", "package.tar.gz")
	err := bootterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bootterm_zip_url := "https://github.com/wtarreau/bootterm/archive/refs/tags/v0.5.zip"
	bootterm_cmd_zip := exec.Command("curl", "-L", bootterm_zip_url, "-o", "package.zip")
	err = bootterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bootterm_bin_url := "https://github.com/wtarreau/bootterm/archive/refs/tags/v0.5.bin"
	bootterm_cmd_bin := exec.Command("curl", "-L", bootterm_bin_url, "-o", "binary.bin")
	err = bootterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bootterm_src_url := "https://github.com/wtarreau/bootterm/archive/refs/tags/v0.5.src.tar.gz"
	bootterm_cmd_src := exec.Command("curl", "-L", bootterm_src_url, "-o", "source.tar.gz")
	err = bootterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bootterm_cmd_direct := exec.Command("./binary")
	err = bootterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
