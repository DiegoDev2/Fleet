package main

// Abduco - Provides session management: i.e. separate programs from terminals
// Homepage: https://www.brain-dump.org/projects/abduco

import (
	"fmt"
	
	"os/exec"
)

func installAbduco() {
	// Método 1: Descargar y extraer .tar.gz
	abduco_tar_url := "https://github.com/martanne/abduco/releases/download/v0.6/abduco-0.6.tar.gz"
	abduco_cmd_tar := exec.Command("curl", "-L", abduco_tar_url, "-o", "package.tar.gz")
	err := abduco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abduco_zip_url := "https://github.com/martanne/abduco/releases/download/v0.6/abduco-0.6.zip"
	abduco_cmd_zip := exec.Command("curl", "-L", abduco_zip_url, "-o", "package.zip")
	err = abduco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abduco_bin_url := "https://github.com/martanne/abduco/releases/download/v0.6/abduco-0.6.bin"
	abduco_cmd_bin := exec.Command("curl", "-L", abduco_bin_url, "-o", "binary.bin")
	err = abduco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abduco_src_url := "https://github.com/martanne/abduco/releases/download/v0.6/abduco-0.6.src.tar.gz"
	abduco_cmd_src := exec.Command("curl", "-L", abduco_src_url, "-o", "source.tar.gz")
	err = abduco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abduco_cmd_direct := exec.Command("./binary")
	err = abduco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
