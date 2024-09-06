package main

// Task - Feature-rich console based todo list manager
// Homepage: https://taskwarrior.org/

import (
	"fmt"
	
	"os/exec"
)

func installTask() {
	// Método 1: Descargar y extraer .tar.gz
	task_tar_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v3.1.0/task-3.1.0.tar.gz"
	task_cmd_tar := exec.Command("curl", "-L", task_tar_url, "-o", "package.tar.gz")
	err := task_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	task_zip_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v3.1.0/task-3.1.0.zip"
	task_cmd_zip := exec.Command("curl", "-L", task_zip_url, "-o", "package.zip")
	err = task_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	task_bin_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v3.1.0/task-3.1.0.bin"
	task_cmd_bin := exec.Command("curl", "-L", task_bin_url, "-o", "binary.bin")
	err = task_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	task_src_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v3.1.0/task-3.1.0.src.tar.gz"
	task_cmd_src := exec.Command("curl", "-L", task_src_url, "-o", "source.tar.gz")
	err = task_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	task_cmd_direct := exec.Command("./binary")
	err = task_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: corrosion")
	exec.Command("latte", "install", "corrosion").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: linux-headers@5.15")
	exec.Command("latte", "install", "linux-headers@5.15").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
