package main

// Httpry - Packet sniffer for displaying and logging HTTP traffic
// Homepage: https://github.com/jbittel/httpry

import (
	"fmt"
	
	"os/exec"
)

func installHttpry() {
	// Método 1: Descargar y extraer .tar.gz
	httpry_tar_url := "https://github.com/jbittel/httpry/archive/refs/tags/httpry-0.1.8.tar.gz"
	httpry_cmd_tar := exec.Command("curl", "-L", httpry_tar_url, "-o", "package.tar.gz")
	err := httpry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpry_zip_url := "https://github.com/jbittel/httpry/archive/refs/tags/httpry-0.1.8.zip"
	httpry_cmd_zip := exec.Command("curl", "-L", httpry_zip_url, "-o", "package.zip")
	err = httpry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpry_bin_url := "https://github.com/jbittel/httpry/archive/refs/tags/httpry-0.1.8.bin"
	httpry_cmd_bin := exec.Command("curl", "-L", httpry_bin_url, "-o", "binary.bin")
	err = httpry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpry_src_url := "https://github.com/jbittel/httpry/archive/refs/tags/httpry-0.1.8.src.tar.gz"
	httpry_cmd_src := exec.Command("curl", "-L", httpry_src_url, "-o", "source.tar.gz")
	err = httpry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpry_cmd_direct := exec.Command("./binary")
	err = httpry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
