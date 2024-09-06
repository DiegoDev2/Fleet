package main

// Microsocks - Tiny, portable SOCKS5 server with very moderate resource usage
// Homepage: https://github.com/rofl0r/microsocks

import (
	"fmt"
	
	"os/exec"
)

func installMicrosocks() {
	// Método 1: Descargar y extraer .tar.gz
	microsocks_tar_url := "https://github.com/rofl0r/microsocks/archive/refs/tags/v1.0.4.tar.gz"
	microsocks_cmd_tar := exec.Command("curl", "-L", microsocks_tar_url, "-o", "package.tar.gz")
	err := microsocks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	microsocks_zip_url := "https://github.com/rofl0r/microsocks/archive/refs/tags/v1.0.4.zip"
	microsocks_cmd_zip := exec.Command("curl", "-L", microsocks_zip_url, "-o", "package.zip")
	err = microsocks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	microsocks_bin_url := "https://github.com/rofl0r/microsocks/archive/refs/tags/v1.0.4.bin"
	microsocks_cmd_bin := exec.Command("curl", "-L", microsocks_bin_url, "-o", "binary.bin")
	err = microsocks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	microsocks_src_url := "https://github.com/rofl0r/microsocks/archive/refs/tags/v1.0.4.src.tar.gz"
	microsocks_cmd_src := exec.Command("curl", "-L", microsocks_src_url, "-o", "source.tar.gz")
	err = microsocks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	microsocks_cmd_direct := exec.Command("./binary")
	err = microsocks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
