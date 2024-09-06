package main

// Webdis - Redis HTTP interface with JSON output
// Homepage: https://webd.is/

import (
	"fmt"
	
	"os/exec"
)

func installWebdis() {
	// Método 1: Descargar y extraer .tar.gz
	webdis_tar_url := "https://github.com/nicolasff/webdis/archive/refs/tags/0.1.22.tar.gz"
	webdis_cmd_tar := exec.Command("curl", "-L", webdis_tar_url, "-o", "package.tar.gz")
	err := webdis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webdis_zip_url := "https://github.com/nicolasff/webdis/archive/refs/tags/0.1.22.zip"
	webdis_cmd_zip := exec.Command("curl", "-L", webdis_zip_url, "-o", "package.zip")
	err = webdis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webdis_bin_url := "https://github.com/nicolasff/webdis/archive/refs/tags/0.1.22.bin"
	webdis_cmd_bin := exec.Command("curl", "-L", webdis_bin_url, "-o", "binary.bin")
	err = webdis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webdis_src_url := "https://github.com/nicolasff/webdis/archive/refs/tags/0.1.22.src.tar.gz"
	webdis_cmd_src := exec.Command("curl", "-L", webdis_src_url, "-o", "source.tar.gz")
	err = webdis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webdis_cmd_direct := exec.Command("./binary")
	err = webdis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
}
