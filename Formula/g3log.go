package main

// G3log - Asynchronous, 'crash safe', logger that is easy to use
// Homepage: https://github.com/KjellKod/g3log

import (
	"fmt"
	
	"os/exec"
)

func installG3log() {
	// Método 1: Descargar y extraer .tar.gz
	g3log_tar_url := "https://github.com/KjellKod/g3log/archive/refs/tags/2.4.tar.gz"
	g3log_cmd_tar := exec.Command("curl", "-L", g3log_tar_url, "-o", "package.tar.gz")
	err := g3log_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	g3log_zip_url := "https://github.com/KjellKod/g3log/archive/refs/tags/2.4.zip"
	g3log_cmd_zip := exec.Command("curl", "-L", g3log_zip_url, "-o", "package.zip")
	err = g3log_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	g3log_bin_url := "https://github.com/KjellKod/g3log/archive/refs/tags/2.4.bin"
	g3log_cmd_bin := exec.Command("curl", "-L", g3log_bin_url, "-o", "binary.bin")
	err = g3log_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	g3log_src_url := "https://github.com/KjellKod/g3log/archive/refs/tags/2.4.src.tar.gz"
	g3log_cmd_src := exec.Command("curl", "-L", g3log_src_url, "-o", "source.tar.gz")
	err = g3log_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	g3log_cmd_direct := exec.Command("./binary")
	err = g3log_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
