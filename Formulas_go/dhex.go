package main

// Dhex - Ncurses based advanced hex editor featuring diff mode and more
// Homepage: https://www.dettus.net/dhex/

import (
	"fmt"
	
	"os/exec"
)

func installDhex() {
	// Método 1: Descargar y extraer .tar.gz
	dhex_tar_url := "https://www.dettus.net/dhex/dhex_0.69.tar.gz"
	dhex_cmd_tar := exec.Command("curl", "-L", dhex_tar_url, "-o", "package.tar.gz")
	err := dhex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhex_zip_url := "https://www.dettus.net/dhex/dhex_0.69.zip"
	dhex_cmd_zip := exec.Command("curl", "-L", dhex_zip_url, "-o", "package.zip")
	err = dhex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhex_bin_url := "https://www.dettus.net/dhex/dhex_0.69.bin"
	dhex_cmd_bin := exec.Command("curl", "-L", dhex_bin_url, "-o", "binary.bin")
	err = dhex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhex_src_url := "https://www.dettus.net/dhex/dhex_0.69.src.tar.gz"
	dhex_cmd_src := exec.Command("curl", "-L", dhex_src_url, "-o", "source.tar.gz")
	err = dhex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhex_cmd_direct := exec.Command("./binary")
	err = dhex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
