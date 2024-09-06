package main

// Tlx - Collection of Sophisticated C++ Data Structures, Algorithms and Helpers
// Homepage: https://tlx.github.io

import (
	"fmt"
	
	"os/exec"
)

func installTlx() {
	// Método 1: Descargar y extraer .tar.gz
	tlx_tar_url := "https://github.com/tlx/tlx/archive/refs/tags/v0.6.1.tar.gz"
	tlx_cmd_tar := exec.Command("curl", "-L", tlx_tar_url, "-o", "package.tar.gz")
	err := tlx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tlx_zip_url := "https://github.com/tlx/tlx/archive/refs/tags/v0.6.1.zip"
	tlx_cmd_zip := exec.Command("curl", "-L", tlx_zip_url, "-o", "package.zip")
	err = tlx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tlx_bin_url := "https://github.com/tlx/tlx/archive/refs/tags/v0.6.1.bin"
	tlx_cmd_bin := exec.Command("curl", "-L", tlx_bin_url, "-o", "binary.bin")
	err = tlx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tlx_src_url := "https://github.com/tlx/tlx/archive/refs/tags/v0.6.1.src.tar.gz"
	tlx_cmd_src := exec.Command("curl", "-L", tlx_src_url, "-o", "source.tar.gz")
	err = tlx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tlx_cmd_direct := exec.Command("./binary")
	err = tlx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
