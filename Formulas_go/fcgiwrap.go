package main

// Fcgiwrap - CGI support for Nginx
// Homepage: https://www.nginx.com/resources/wiki/start/topics/examples/fcgiwrap/

import (
	"fmt"
	
	"os/exec"
)

func installFcgiwrap() {
	// Método 1: Descargar y extraer .tar.gz
	fcgiwrap_tar_url := "https://github.com/gnosek/fcgiwrap/archive/refs/tags/1.1.0.tar.gz"
	fcgiwrap_cmd_tar := exec.Command("curl", "-L", fcgiwrap_tar_url, "-o", "package.tar.gz")
	err := fcgiwrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fcgiwrap_zip_url := "https://github.com/gnosek/fcgiwrap/archive/refs/tags/1.1.0.zip"
	fcgiwrap_cmd_zip := exec.Command("curl", "-L", fcgiwrap_zip_url, "-o", "package.zip")
	err = fcgiwrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fcgiwrap_bin_url := "https://github.com/gnosek/fcgiwrap/archive/refs/tags/1.1.0.bin"
	fcgiwrap_cmd_bin := exec.Command("curl", "-L", fcgiwrap_bin_url, "-o", "binary.bin")
	err = fcgiwrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fcgiwrap_src_url := "https://github.com/gnosek/fcgiwrap/archive/refs/tags/1.1.0.src.tar.gz"
	fcgiwrap_cmd_src := exec.Command("curl", "-L", fcgiwrap_src_url, "-o", "source.tar.gz")
	err = fcgiwrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fcgiwrap_cmd_direct := exec.Command("./binary")
	err = fcgiwrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fcgi")
exec.Command("latte", "install", "fcgi")
}
