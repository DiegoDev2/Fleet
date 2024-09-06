package main

// Conserver - Allows multiple users to watch a serial console at the same time
// Homepage: https://www.conserver.com/

import (
	"fmt"
	
	"os/exec"
)

func installConserver() {
	// Método 1: Descargar y extraer .tar.gz
	conserver_tar_url := "https://github.com/bstansell/conserver/releases/download/v8.2.7/conserver-8.2.7.tar.gz"
	conserver_cmd_tar := exec.Command("curl", "-L", conserver_tar_url, "-o", "package.tar.gz")
	err := conserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conserver_zip_url := "https://github.com/bstansell/conserver/releases/download/v8.2.7/conserver-8.2.7.zip"
	conserver_cmd_zip := exec.Command("curl", "-L", conserver_zip_url, "-o", "package.zip")
	err = conserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conserver_bin_url := "https://github.com/bstansell/conserver/releases/download/v8.2.7/conserver-8.2.7.bin"
	conserver_cmd_bin := exec.Command("curl", "-L", conserver_bin_url, "-o", "binary.bin")
	err = conserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conserver_src_url := "https://github.com/bstansell/conserver/releases/download/v8.2.7/conserver-8.2.7.src.tar.gz"
	conserver_cmd_src := exec.Command("curl", "-L", conserver_src_url, "-o", "source.tar.gz")
	err = conserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conserver_cmd_direct := exec.Command("./binary")
	err = conserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
