package main

// Sl - Prints a steam locomotive if you type sl instead of ls
// Homepage: https://github.com/mtoyoda/sl

import (
	"fmt"
	
	"os/exec"
)

func installSl() {
	// Método 1: Descargar y extraer .tar.gz
	sl_tar_url := "https://github.com/mtoyoda/sl/archive/refs/tags/5.02.tar.gz"
	sl_cmd_tar := exec.Command("curl", "-L", sl_tar_url, "-o", "package.tar.gz")
	err := sl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sl_zip_url := "https://github.com/mtoyoda/sl/archive/refs/tags/5.02.zip"
	sl_cmd_zip := exec.Command("curl", "-L", sl_zip_url, "-o", "package.zip")
	err = sl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sl_bin_url := "https://github.com/mtoyoda/sl/archive/refs/tags/5.02.bin"
	sl_cmd_bin := exec.Command("curl", "-L", sl_bin_url, "-o", "binary.bin")
	err = sl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sl_src_url := "https://github.com/mtoyoda/sl/archive/refs/tags/5.02.src.tar.gz"
	sl_cmd_src := exec.Command("curl", "-L", sl_src_url, "-o", "source.tar.gz")
	err = sl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sl_cmd_direct := exec.Command("./binary")
	err = sl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
