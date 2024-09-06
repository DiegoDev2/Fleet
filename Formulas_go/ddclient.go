package main

// Ddclient - Update dynamic DNS entries
// Homepage: https://ddclient.net/

import (
	"fmt"
	
	"os/exec"
)

func installDdclient() {
	// Método 1: Descargar y extraer .tar.gz
	ddclient_tar_url := "https://github.com/ddclient/ddclient/archive/refs/tags/v3.11.2.tar.gz"
	ddclient_cmd_tar := exec.Command("curl", "-L", ddclient_tar_url, "-o", "package.tar.gz")
	err := ddclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddclient_zip_url := "https://github.com/ddclient/ddclient/archive/refs/tags/v3.11.2.zip"
	ddclient_cmd_zip := exec.Command("curl", "-L", ddclient_zip_url, "-o", "package.zip")
	err = ddclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddclient_bin_url := "https://github.com/ddclient/ddclient/archive/refs/tags/v3.11.2.bin"
	ddclient_cmd_bin := exec.Command("curl", "-L", ddclient_bin_url, "-o", "binary.bin")
	err = ddclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddclient_src_url := "https://github.com/ddclient/ddclient/archive/refs/tags/v3.11.2.src.tar.gz"
	ddclient_cmd_src := exec.Command("curl", "-L", ddclient_src_url, "-o", "source.tar.gz")
	err = ddclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddclient_cmd_direct := exec.Command("./binary")
	err = ddclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
