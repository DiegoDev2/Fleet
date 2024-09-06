package main

// PipesSh - Animated pipes terminal screensaver
// Homepage: https://github.com/pipeseroni/pipes.sh

import (
	"fmt"
	
	"os/exec"
)

func installPipesSh() {
	// Método 1: Descargar y extraer .tar.gz
	pipessh_tar_url := "https://github.com/pipeseroni/pipes.sh/archive/refs/tags/v1.3.0.tar.gz"
	pipessh_cmd_tar := exec.Command("curl", "-L", pipessh_tar_url, "-o", "package.tar.gz")
	err := pipessh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipessh_zip_url := "https://github.com/pipeseroni/pipes.sh/archive/refs/tags/v1.3.0.zip"
	pipessh_cmd_zip := exec.Command("curl", "-L", pipessh_zip_url, "-o", "package.zip")
	err = pipessh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipessh_bin_url := "https://github.com/pipeseroni/pipes.sh/archive/refs/tags/v1.3.0.bin"
	pipessh_cmd_bin := exec.Command("curl", "-L", pipessh_bin_url, "-o", "binary.bin")
	err = pipessh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipessh_src_url := "https://github.com/pipeseroni/pipes.sh/archive/refs/tags/v1.3.0.src.tar.gz"
	pipessh_cmd_src := exec.Command("curl", "-L", pipessh_src_url, "-o", "source.tar.gz")
	err = pipessh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipessh_cmd_direct := exec.Command("./binary")
	err = pipessh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
}
