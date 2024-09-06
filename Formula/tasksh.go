package main

// Tasksh - Shell wrapper for Taskwarrior commands
// Homepage: https://gothenburgbitfactory.org/projects/tasksh.html

import (
	"fmt"
	
	"os/exec"
)

func installTasksh() {
	// Método 1: Descargar y extraer .tar.gz
	tasksh_tar_url := "https://github.com/GothenburgBitFactory/taskshell/releases/download/v1.2.0/tasksh-1.2.0.tar.gz"
	tasksh_cmd_tar := exec.Command("curl", "-L", tasksh_tar_url, "-o", "package.tar.gz")
	err := tasksh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tasksh_zip_url := "https://github.com/GothenburgBitFactory/taskshell/releases/download/v1.2.0/tasksh-1.2.0.zip"
	tasksh_cmd_zip := exec.Command("curl", "-L", tasksh_zip_url, "-o", "package.zip")
	err = tasksh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tasksh_bin_url := "https://github.com/GothenburgBitFactory/taskshell/releases/download/v1.2.0/tasksh-1.2.0.bin"
	tasksh_cmd_bin := exec.Command("curl", "-L", tasksh_bin_url, "-o", "binary.bin")
	err = tasksh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tasksh_src_url := "https://github.com/GothenburgBitFactory/taskshell/releases/download/v1.2.0/tasksh-1.2.0.src.tar.gz"
	tasksh_cmd_src := exec.Command("curl", "-L", tasksh_src_url, "-o", "source.tar.gz")
	err = tasksh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tasksh_cmd_direct := exec.Command("./binary")
	err = tasksh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: task")
	exec.Command("latte", "install", "task").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
