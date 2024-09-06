package main

// Squid - Advanced proxy caching server for HTTP, HTTPS, FTP, and Gopher
// Homepage: https://www.squid-cache.org/

import (
	"fmt"
	
	"os/exec"
)

func installSquid() {
	// Método 1: Descargar y extraer .tar.gz
	squid_tar_url := "http://www.squid-cache.org/Versions/v6/squid-6.10.tar.xz"
	squid_cmd_tar := exec.Command("curl", "-L", squid_tar_url, "-o", "package.tar.gz")
	err := squid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squid_zip_url := "http://www.squid-cache.org/Versions/v6/squid-6.10.tar.xz"
	squid_cmd_zip := exec.Command("curl", "-L", squid_zip_url, "-o", "package.zip")
	err = squid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squid_bin_url := "http://www.squid-cache.org/Versions/v6/squid-6.10.tar.xz"
	squid_cmd_bin := exec.Command("curl", "-L", squid_bin_url, "-o", "binary.bin")
	err = squid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squid_src_url := "http://www.squid-cache.org/Versions/v6/squid-6.10.tar.xz"
	squid_cmd_src := exec.Command("curl", "-L", squid_src_url, "-o", "source.tar.gz")
	err = squid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squid_cmd_direct := exec.Command("./binary")
	err = squid_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
