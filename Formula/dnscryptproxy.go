package main

// DnscryptProxy - Secure communications between a client and a DNS resolver
// Homepage: https://dnscrypt.info

import (
	"fmt"
	
	"os/exec"
)

func installDnscryptProxy() {
	// Método 1: Descargar y extraer .tar.gz
	dnscryptproxy_tar_url := "https://github.com/DNSCrypt/dnscrypt-proxy/archive/refs/tags/2.1.5.tar.gz"
	dnscryptproxy_cmd_tar := exec.Command("curl", "-L", dnscryptproxy_tar_url, "-o", "package.tar.gz")
	err := dnscryptproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnscryptproxy_zip_url := "https://github.com/DNSCrypt/dnscrypt-proxy/archive/refs/tags/2.1.5.zip"
	dnscryptproxy_cmd_zip := exec.Command("curl", "-L", dnscryptproxy_zip_url, "-o", "package.zip")
	err = dnscryptproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnscryptproxy_bin_url := "https://github.com/DNSCrypt/dnscrypt-proxy/archive/refs/tags/2.1.5.bin"
	dnscryptproxy_cmd_bin := exec.Command("curl", "-L", dnscryptproxy_bin_url, "-o", "binary.bin")
	err = dnscryptproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnscryptproxy_src_url := "https://github.com/DNSCrypt/dnscrypt-proxy/archive/refs/tags/2.1.5.src.tar.gz"
	dnscryptproxy_cmd_src := exec.Command("curl", "-L", dnscryptproxy_src_url, "-o", "source.tar.gz")
	err = dnscryptproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnscryptproxy_cmd_direct := exec.Command("./binary")
	err = dnscryptproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
