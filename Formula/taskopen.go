package main

// Taskopen - Tool for taking notes and open urls with taskwarrior
// Homepage: https://github.com/jschlatow/taskopen

import (
	"fmt"
	
	"os/exec"
)

func installTaskopen() {
	// Método 1: Descargar y extraer .tar.gz
	taskopen_tar_url := "https://github.com/jschlatow/taskopen/archive/refs/tags/v2.0.1.tar.gz"
	taskopen_cmd_tar := exec.Command("curl", "-L", taskopen_tar_url, "-o", "package.tar.gz")
	err := taskopen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskopen_zip_url := "https://github.com/jschlatow/taskopen/archive/refs/tags/v2.0.1.zip"
	taskopen_cmd_zip := exec.Command("curl", "-L", taskopen_zip_url, "-o", "package.zip")
	err = taskopen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskopen_bin_url := "https://github.com/jschlatow/taskopen/archive/refs/tags/v2.0.1.bin"
	taskopen_cmd_bin := exec.Command("curl", "-L", taskopen_bin_url, "-o", "binary.bin")
	err = taskopen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskopen_src_url := "https://github.com/jschlatow/taskopen/archive/refs/tags/v2.0.1.src.tar.gz"
	taskopen_cmd_src := exec.Command("curl", "-L", taskopen_src_url, "-o", "source.tar.gz")
	err = taskopen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskopen_cmd_direct := exec.Command("./binary")
	err = taskopen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nim")
	exec.Command("latte", "install", "nim").Run()
	fmt.Println("Instalando dependencia: task")
	exec.Command("latte", "install", "task").Run()
}
