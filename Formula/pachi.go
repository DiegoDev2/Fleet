package main

// Pachi - Software for the Board Game of Go/Weiqi/Baduk
// Homepage: https://pachi.or.cz/

import (
	"fmt"
	
	"os/exec"
)

func installPachi() {
	// Método 1: Descargar y extraer .tar.gz
	pachi_tar_url := "https://github.com/pasky/pachi/archive/refs/tags/pachi-12.84.tar.gz"
	pachi_cmd_tar := exec.Command("curl", "-L", pachi_tar_url, "-o", "package.tar.gz")
	err := pachi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pachi_zip_url := "https://github.com/pasky/pachi/archive/refs/tags/pachi-12.84.zip"
	pachi_cmd_zip := exec.Command("curl", "-L", pachi_zip_url, "-o", "package.zip")
	err = pachi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pachi_bin_url := "https://github.com/pasky/pachi/archive/refs/tags/pachi-12.84.bin"
	pachi_cmd_bin := exec.Command("curl", "-L", pachi_bin_url, "-o", "binary.bin")
	err = pachi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pachi_src_url := "https://github.com/pasky/pachi/archive/refs/tags/pachi-12.84.src.tar.gz"
	pachi_cmd_src := exec.Command("curl", "-L", pachi_src_url, "-o", "source.tar.gz")
	err = pachi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pachi_cmd_direct := exec.Command("./binary")
	err = pachi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
