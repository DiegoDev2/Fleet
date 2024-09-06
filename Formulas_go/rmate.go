package main

// Rmate - Edit files from an SSH session in TextMate
// Homepage: https://github.com/textmate/rmate

import (
	"fmt"
	
	"os/exec"
)

func installRmate() {
	// Método 1: Descargar y extraer .tar.gz
	rmate_tar_url := "https://github.com/textmate/rmate/archive/refs/tags/v1.5.8.tar.gz"
	rmate_cmd_tar := exec.Command("curl", "-L", rmate_tar_url, "-o", "package.tar.gz")
	err := rmate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rmate_zip_url := "https://github.com/textmate/rmate/archive/refs/tags/v1.5.8.zip"
	rmate_cmd_zip := exec.Command("curl", "-L", rmate_zip_url, "-o", "package.zip")
	err = rmate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rmate_bin_url := "https://github.com/textmate/rmate/archive/refs/tags/v1.5.8.bin"
	rmate_cmd_bin := exec.Command("curl", "-L", rmate_bin_url, "-o", "binary.bin")
	err = rmate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rmate_src_url := "https://github.com/textmate/rmate/archive/refs/tags/v1.5.8.src.tar.gz"
	rmate_cmd_src := exec.Command("curl", "-L", rmate_src_url, "-o", "source.tar.gz")
	err = rmate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rmate_cmd_direct := exec.Command("./binary")
	err = rmate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
