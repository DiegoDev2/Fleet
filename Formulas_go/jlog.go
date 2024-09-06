package main

// Jlog - Pure C message queue with subscribers and publishers for logs
// Homepage: https://labs.omniti.com/labs/jlog

import (
	"fmt"
	
	"os/exec"
)

func installJlog() {
	// Método 1: Descargar y extraer .tar.gz
	jlog_tar_url := "https://github.com/omniti-labs/jlog/archive/refs/tags/2.6.0.tar.gz"
	jlog_cmd_tar := exec.Command("curl", "-L", jlog_tar_url, "-o", "package.tar.gz")
	err := jlog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jlog_zip_url := "https://github.com/omniti-labs/jlog/archive/refs/tags/2.6.0.zip"
	jlog_cmd_zip := exec.Command("curl", "-L", jlog_zip_url, "-o", "package.zip")
	err = jlog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jlog_bin_url := "https://github.com/omniti-labs/jlog/archive/refs/tags/2.6.0.bin"
	jlog_cmd_bin := exec.Command("curl", "-L", jlog_bin_url, "-o", "binary.bin")
	err = jlog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jlog_src_url := "https://github.com/omniti-labs/jlog/archive/refs/tags/2.6.0.src.tar.gz"
	jlog_cmd_src := exec.Command("curl", "-L", jlog_src_url, "-o", "source.tar.gz")
	err = jlog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jlog_cmd_direct := exec.Command("./binary")
	err = jlog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
