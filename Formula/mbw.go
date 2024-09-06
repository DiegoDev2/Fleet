package main

// Mbw - Memory Bandwidth Benchmark
// Homepage: https://github.com/raas/mbw/

import (
	"fmt"
	
	"os/exec"
)

func installMbw() {
	// Método 1: Descargar y extraer .tar.gz
	mbw_tar_url := "https://github.com/raas/mbw/archive/refs/tags/v2.0.tar.gz"
	mbw_cmd_tar := exec.Command("curl", "-L", mbw_tar_url, "-o", "package.tar.gz")
	err := mbw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mbw_zip_url := "https://github.com/raas/mbw/archive/refs/tags/v2.0.zip"
	mbw_cmd_zip := exec.Command("curl", "-L", mbw_zip_url, "-o", "package.zip")
	err = mbw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mbw_bin_url := "https://github.com/raas/mbw/archive/refs/tags/v2.0.bin"
	mbw_cmd_bin := exec.Command("curl", "-L", mbw_bin_url, "-o", "binary.bin")
	err = mbw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mbw_src_url := "https://github.com/raas/mbw/archive/refs/tags/v2.0.src.tar.gz"
	mbw_cmd_src := exec.Command("curl", "-L", mbw_src_url, "-o", "source.tar.gz")
	err = mbw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mbw_cmd_direct := exec.Command("./binary")
	err = mbw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
