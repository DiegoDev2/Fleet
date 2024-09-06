package main

// Threemux - Terminal multiplexer inspired by i3
// Homepage: https://github.com/aaronjanse/3mux

import (
	"fmt"
	
	"os/exec"
)

func installThreemux() {
	// Método 1: Descargar y extraer .tar.gz
	threemux_tar_url := "https://github.com/aaronjanse/3mux/archive/refs/tags/v1.1.0.tar.gz"
	threemux_cmd_tar := exec.Command("curl", "-L", threemux_tar_url, "-o", "package.tar.gz")
	err := threemux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	threemux_zip_url := "https://github.com/aaronjanse/3mux/archive/refs/tags/v1.1.0.zip"
	threemux_cmd_zip := exec.Command("curl", "-L", threemux_zip_url, "-o", "package.zip")
	err = threemux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	threemux_bin_url := "https://github.com/aaronjanse/3mux/archive/refs/tags/v1.1.0.bin"
	threemux_cmd_bin := exec.Command("curl", "-L", threemux_bin_url, "-o", "binary.bin")
	err = threemux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	threemux_src_url := "https://github.com/aaronjanse/3mux/archive/refs/tags/v1.1.0.src.tar.gz"
	threemux_cmd_src := exec.Command("curl", "-L", threemux_src_url, "-o", "source.tar.gz")
	err = threemux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	threemux_cmd_direct := exec.Command("./binary")
	err = threemux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
