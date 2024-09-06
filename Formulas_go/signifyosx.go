package main

// SignifyOsx - Cryptographically sign and verify files
// Homepage: https://man.openbsd.org/signify.1

import (
	"fmt"
	
	"os/exec"
)

func installSignifyOsx() {
	// Método 1: Descargar y extraer .tar.gz
	signifyosx_tar_url := "https://github.com/jpouellet/signify-osx/archive/refs/tags/1.4.tar.gz"
	signifyosx_cmd_tar := exec.Command("curl", "-L", signifyosx_tar_url, "-o", "package.tar.gz")
	err := signifyosx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	signifyosx_zip_url := "https://github.com/jpouellet/signify-osx/archive/refs/tags/1.4.zip"
	signifyosx_cmd_zip := exec.Command("curl", "-L", signifyosx_zip_url, "-o", "package.zip")
	err = signifyosx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	signifyosx_bin_url := "https://github.com/jpouellet/signify-osx/archive/refs/tags/1.4.bin"
	signifyosx_cmd_bin := exec.Command("curl", "-L", signifyosx_bin_url, "-o", "binary.bin")
	err = signifyosx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	signifyosx_src_url := "https://github.com/jpouellet/signify-osx/archive/refs/tags/1.4.src.tar.gz"
	signifyosx_cmd_src := exec.Command("curl", "-L", signifyosx_src_url, "-o", "source.tar.gz")
	err = signifyosx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	signifyosx_cmd_direct := exec.Command("./binary")
	err = signifyosx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
