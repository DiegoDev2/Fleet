package main

// Getdns - Modern asynchronous DNS API
// Homepage: https://getdnsapi.net

import (
	"fmt"
	
	"os/exec"
)

func installGetdns() {
	// Método 1: Descargar y extraer .tar.gz
	getdns_tar_url := "https://getdnsapi.net/releases/getdns-1-7-3/getdns-1.7.3.tar.gz"
	getdns_cmd_tar := exec.Command("curl", "-L", getdns_tar_url, "-o", "package.tar.gz")
	err := getdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	getdns_zip_url := "https://getdnsapi.net/releases/getdns-1-7-3/getdns-1.7.3.zip"
	getdns_cmd_zip := exec.Command("curl", "-L", getdns_zip_url, "-o", "package.zip")
	err = getdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	getdns_bin_url := "https://getdnsapi.net/releases/getdns-1-7-3/getdns-1.7.3.bin"
	getdns_cmd_bin := exec.Command("curl", "-L", getdns_bin_url, "-o", "binary.bin")
	err = getdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	getdns_src_url := "https://getdnsapi.net/releases/getdns-1-7-3/getdns-1.7.3.src.tar.gz"
	getdns_cmd_src := exec.Command("curl", "-L", getdns_src_url, "-o", "source.tar.gz")
	err = getdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	getdns_cmd_direct := exec.Command("./binary")
	err = getdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libev")
	exec.Command("latte", "install", "libev").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: unbound")
	exec.Command("latte", "install", "unbound").Run()
}
