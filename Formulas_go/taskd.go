package main

// Taskd - Client-server synchronization for todo lists
// Homepage: https://taskwarrior.org/docs/taskserver/setup.html

import (
	"fmt"
	
	"os/exec"
)

func installTaskd() {
	// Método 1: Descargar y extraer .tar.gz
	taskd_tar_url := "https://github.com/GothenburgBitFactory/taskserver/releases/download/v1.1.0/taskd-1.1.0.tar.gz"
	taskd_cmd_tar := exec.Command("curl", "-L", taskd_tar_url, "-o", "package.tar.gz")
	err := taskd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskd_zip_url := "https://github.com/GothenburgBitFactory/taskserver/releases/download/v1.1.0/taskd-1.1.0.zip"
	taskd_cmd_zip := exec.Command("curl", "-L", taskd_zip_url, "-o", "package.zip")
	err = taskd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskd_bin_url := "https://github.com/GothenburgBitFactory/taskserver/releases/download/v1.1.0/taskd-1.1.0.bin"
	taskd_cmd_bin := exec.Command("curl", "-L", taskd_bin_url, "-o", "binary.bin")
	err = taskd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskd_src_url := "https://github.com/GothenburgBitFactory/taskserver/releases/download/v1.1.0/taskd-1.1.0.src.tar.gz"
	taskd_cmd_src := exec.Command("curl", "-L", taskd_src_url, "-o", "source.tar.gz")
	err = taskd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskd_cmd_direct := exec.Command("./binary")
	err = taskd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
