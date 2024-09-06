package main

// TaskAT2 - Feature-rich console based todo list manager
// Homepage: https://taskwarrior.org/

import (
	"fmt"
	
	"os/exec"
)

func installTaskAT2() {
	// Método 1: Descargar y extraer .tar.gz
	taskat2_tar_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v2.6.2/task-2.6.2.tar.gz"
	taskat2_cmd_tar := exec.Command("curl", "-L", taskat2_tar_url, "-o", "package.tar.gz")
	err := taskat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskat2_zip_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v2.6.2/task-2.6.2.zip"
	taskat2_cmd_zip := exec.Command("curl", "-L", taskat2_zip_url, "-o", "package.zip")
	err = taskat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskat2_bin_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v2.6.2/task-2.6.2.bin"
	taskat2_cmd_bin := exec.Command("curl", "-L", taskat2_bin_url, "-o", "binary.bin")
	err = taskat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskat2_src_url := "https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v2.6.2/task-2.6.2.src.tar.gz"
	taskat2_cmd_src := exec.Command("curl", "-L", taskat2_src_url, "-o", "source.tar.gz")
	err = taskat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskat2_cmd_direct := exec.Command("./binary")
	err = taskat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: linux-headers@5.15")
	exec.Command("latte", "install", "linux-headers@5.15").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
