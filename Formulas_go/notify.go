package main

// Notify - Stream the output of any CLI and publish it to a variety of supported platforms
// Homepage: https://github.com/projectdiscovery/notify

import (
	"fmt"
	
	"os/exec"
)

func installNotify() {
	// Método 1: Descargar y extraer .tar.gz
	notify_tar_url := "https://github.com/projectdiscovery/notify/archive/refs/tags/v1.0.6.tar.gz"
	notify_cmd_tar := exec.Command("curl", "-L", notify_tar_url, "-o", "package.tar.gz")
	err := notify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	notify_zip_url := "https://github.com/projectdiscovery/notify/archive/refs/tags/v1.0.6.zip"
	notify_cmd_zip := exec.Command("curl", "-L", notify_zip_url, "-o", "package.zip")
	err = notify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	notify_bin_url := "https://github.com/projectdiscovery/notify/archive/refs/tags/v1.0.6.bin"
	notify_cmd_bin := exec.Command("curl", "-L", notify_bin_url, "-o", "binary.bin")
	err = notify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	notify_src_url := "https://github.com/projectdiscovery/notify/archive/refs/tags/v1.0.6.src.tar.gz"
	notify_cmd_src := exec.Command("curl", "-L", notify_src_url, "-o", "source.tar.gz")
	err = notify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	notify_cmd_direct := exec.Command("./binary")
	err = notify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
