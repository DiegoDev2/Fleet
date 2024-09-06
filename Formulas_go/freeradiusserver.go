package main

// FreeradiusServer - High-performance and highly configurable RADIUS server
// Homepage: https://freeradius.org/

import (
	"fmt"
	
	"os/exec"
)

func installFreeradiusServer() {
	// Método 1: Descargar y extraer .tar.gz
	freeradiusserver_tar_url := "https://github.com/FreeRADIUS/freeradius-server/archive/refs/tags/release_3_2_6.tar.gz"
	freeradiusserver_cmd_tar := exec.Command("curl", "-L", freeradiusserver_tar_url, "-o", "package.tar.gz")
	err := freeradiusserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeradiusserver_zip_url := "https://github.com/FreeRADIUS/freeradius-server/archive/refs/tags/release_3_2_6.zip"
	freeradiusserver_cmd_zip := exec.Command("curl", "-L", freeradiusserver_zip_url, "-o", "package.zip")
	err = freeradiusserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeradiusserver_bin_url := "https://github.com/FreeRADIUS/freeradius-server/archive/refs/tags/release_3_2_6.bin"
	freeradiusserver_cmd_bin := exec.Command("curl", "-L", freeradiusserver_bin_url, "-o", "binary.bin")
	err = freeradiusserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeradiusserver_src_url := "https://github.com/FreeRADIUS/freeradius-server/archive/refs/tags/release_3_2_6.src.tar.gz"
	freeradiusserver_cmd_src := exec.Command("curl", "-L", freeradiusserver_src_url, "-o", "source.tar.gz")
	err = freeradiusserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeradiusserver_cmd_direct := exec.Command("./binary")
	err = freeradiusserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: collectd")
exec.Command("latte", "install", "collectd")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: talloc")
exec.Command("latte", "install", "talloc")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
