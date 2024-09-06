package main

// Yetris - Customizable Tetris for the terminal
// Homepage: https://github.com/alexdantas/yetris

import (
	"fmt"
	
	"os/exec"
)

func installYetris() {
	// Método 1: Descargar y extraer .tar.gz
	yetris_tar_url := "https://github.com/alexdantas/yetris/archive/refs/tags/v2.3.0.tar.gz"
	yetris_cmd_tar := exec.Command("curl", "-L", yetris_tar_url, "-o", "package.tar.gz")
	err := yetris_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yetris_zip_url := "https://github.com/alexdantas/yetris/archive/refs/tags/v2.3.0.zip"
	yetris_cmd_zip := exec.Command("curl", "-L", yetris_zip_url, "-o", "package.zip")
	err = yetris_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yetris_bin_url := "https://github.com/alexdantas/yetris/archive/refs/tags/v2.3.0.bin"
	yetris_cmd_bin := exec.Command("curl", "-L", yetris_bin_url, "-o", "binary.bin")
	err = yetris_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yetris_src_url := "https://github.com/alexdantas/yetris/archive/refs/tags/v2.3.0.src.tar.gz"
	yetris_cmd_src := exec.Command("curl", "-L", yetris_src_url, "-o", "source.tar.gz")
	err = yetris_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yetris_cmd_direct := exec.Command("./binary")
	err = yetris_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
