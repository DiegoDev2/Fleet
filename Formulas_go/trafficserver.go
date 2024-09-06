package main

// Trafficserver - HTTP/1.1 and HTTP/2 compliant caching proxy server
// Homepage: https://trafficserver.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTrafficserver() {
	// Método 1: Descargar y extraer .tar.gz
	trafficserver_tar_url := "https://downloads.apache.org/trafficserver/trafficserver-9.2.5.tar.bz2"
	trafficserver_cmd_tar := exec.Command("curl", "-L", trafficserver_tar_url, "-o", "package.tar.gz")
	err := trafficserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trafficserver_zip_url := "https://downloads.apache.org/trafficserver/trafficserver-9.2.5.tar.bz2"
	trafficserver_cmd_zip := exec.Command("curl", "-L", trafficserver_zip_url, "-o", "package.zip")
	err = trafficserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trafficserver_bin_url := "https://downloads.apache.org/trafficserver/trafficserver-9.2.5.tar.bz2"
	trafficserver_cmd_bin := exec.Command("curl", "-L", trafficserver_bin_url, "-o", "binary.bin")
	err = trafficserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trafficserver_src_url := "https://downloads.apache.org/trafficserver/trafficserver-9.2.5.tar.bz2"
	trafficserver_cmd_src := exec.Command("curl", "-L", trafficserver_src_url, "-o", "source.tar.gz")
	err = trafficserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trafficserver_cmd_direct := exec.Command("./binary")
	err = trafficserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: yaml-cpp")
exec.Command("latte", "install", "yaml-cpp")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
