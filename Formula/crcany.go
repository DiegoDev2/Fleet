package main

// Crcany - Compute any CRC, a bit at a time, a byte at a time, and a word at a time
// Homepage: https://github.com/madler/crcany

import (
	"fmt"
	
	"os/exec"
)

func installCrcany() {
	// Método 1: Descargar y extraer .tar.gz
	crcany_tar_url := "https://github.com/madler/crcany/archive/refs/tags/v2.1.tar.gz"
	crcany_cmd_tar := exec.Command("curl", "-L", crcany_tar_url, "-o", "package.tar.gz")
	err := crcany_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crcany_zip_url := "https://github.com/madler/crcany/archive/refs/tags/v2.1.zip"
	crcany_cmd_zip := exec.Command("curl", "-L", crcany_zip_url, "-o", "package.zip")
	err = crcany_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crcany_bin_url := "https://github.com/madler/crcany/archive/refs/tags/v2.1.bin"
	crcany_cmd_bin := exec.Command("curl", "-L", crcany_bin_url, "-o", "binary.bin")
	err = crcany_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crcany_src_url := "https://github.com/madler/crcany/archive/refs/tags/v2.1.src.tar.gz"
	crcany_cmd_src := exec.Command("curl", "-L", crcany_src_url, "-o", "source.tar.gz")
	err = crcany_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crcany_cmd_direct := exec.Command("./binary")
	err = crcany_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
