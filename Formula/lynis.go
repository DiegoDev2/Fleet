package main

// Lynis - Security and system auditing tool to harden systems
// Homepage: https://cisofy.com/lynis/

import (
	"fmt"
	
	"os/exec"
)

func installLynis() {
	// Método 1: Descargar y extraer .tar.gz
	lynis_tar_url := "https://github.com/CISOfy/lynis/archive/refs/tags/3.1.1.tar.gz"
	lynis_cmd_tar := exec.Command("curl", "-L", lynis_tar_url, "-o", "package.tar.gz")
	err := lynis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lynis_zip_url := "https://github.com/CISOfy/lynis/archive/refs/tags/3.1.1.zip"
	lynis_cmd_zip := exec.Command("curl", "-L", lynis_zip_url, "-o", "package.zip")
	err = lynis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lynis_bin_url := "https://github.com/CISOfy/lynis/archive/refs/tags/3.1.1.bin"
	lynis_cmd_bin := exec.Command("curl", "-L", lynis_bin_url, "-o", "binary.bin")
	err = lynis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lynis_src_url := "https://github.com/CISOfy/lynis/archive/refs/tags/3.1.1.src.tar.gz"
	lynis_cmd_src := exec.Command("curl", "-L", lynis_src_url, "-o", "source.tar.gz")
	err = lynis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lynis_cmd_direct := exec.Command("./binary")
	err = lynis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
