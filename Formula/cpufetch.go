package main

// Cpufetch - CPU architecture fetching tool
// Homepage: https://github.com/Dr-Noob/cpufetch

import (
	"fmt"
	
	"os/exec"
)

func installCpufetch() {
	// Método 1: Descargar y extraer .tar.gz
	cpufetch_tar_url := "https://github.com/Dr-Noob/cpufetch/archive/refs/tags/v1.06.tar.gz"
	cpufetch_cmd_tar := exec.Command("curl", "-L", cpufetch_tar_url, "-o", "package.tar.gz")
	err := cpufetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpufetch_zip_url := "https://github.com/Dr-Noob/cpufetch/archive/refs/tags/v1.06.zip"
	cpufetch_cmd_zip := exec.Command("curl", "-L", cpufetch_zip_url, "-o", "package.zip")
	err = cpufetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpufetch_bin_url := "https://github.com/Dr-Noob/cpufetch/archive/refs/tags/v1.06.bin"
	cpufetch_cmd_bin := exec.Command("curl", "-L", cpufetch_bin_url, "-o", "binary.bin")
	err = cpufetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpufetch_src_url := "https://github.com/Dr-Noob/cpufetch/archive/refs/tags/v1.06.src.tar.gz"
	cpufetch_cmd_src := exec.Command("curl", "-L", cpufetch_src_url, "-o", "source.tar.gz")
	err = cpufetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpufetch_cmd_direct := exec.Command("./binary")
	err = cpufetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
