package main

// Mediamtx - Zero-dependency real-time media server and media proxy
// Homepage: https://github.com/bluenviron/mediamtx

import (
	"fmt"
	
	"os/exec"
)

func installMediamtx() {
	// Método 1: Descargar y extraer .tar.gz
	mediamtx_tar_url := "https://github.com/bluenviron/mediamtx/archive/refs/tags/v1.9.0.tar.gz"
	mediamtx_cmd_tar := exec.Command("curl", "-L", mediamtx_tar_url, "-o", "package.tar.gz")
	err := mediamtx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mediamtx_zip_url := "https://github.com/bluenviron/mediamtx/archive/refs/tags/v1.9.0.zip"
	mediamtx_cmd_zip := exec.Command("curl", "-L", mediamtx_zip_url, "-o", "package.zip")
	err = mediamtx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mediamtx_bin_url := "https://github.com/bluenviron/mediamtx/archive/refs/tags/v1.9.0.bin"
	mediamtx_cmd_bin := exec.Command("curl", "-L", mediamtx_bin_url, "-o", "binary.bin")
	err = mediamtx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mediamtx_src_url := "https://github.com/bluenviron/mediamtx/archive/refs/tags/v1.9.0.src.tar.gz"
	mediamtx_cmd_src := exec.Command("curl", "-L", mediamtx_src_url, "-o", "source.tar.gz")
	err = mediamtx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mediamtx_cmd_direct := exec.Command("./binary")
	err = mediamtx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
