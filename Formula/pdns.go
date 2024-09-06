package main

// Pdns - Authoritative nameserver
// Homepage: https://www.powerdns.com

import (
	"fmt"
	
	"os/exec"
)

func installPdns() {
	// Método 1: Descargar y extraer .tar.gz
	pdns_tar_url := "https://downloads.powerdns.com/releases/pdns-4.9.1.tar.bz2"
	pdns_cmd_tar := exec.Command("curl", "-L", pdns_tar_url, "-o", "package.tar.gz")
	err := pdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdns_zip_url := "https://downloads.powerdns.com/releases/pdns-4.9.1.tar.bz2"
	pdns_cmd_zip := exec.Command("curl", "-L", pdns_zip_url, "-o", "package.zip")
	err = pdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdns_bin_url := "https://downloads.powerdns.com/releases/pdns-4.9.1.tar.bz2"
	pdns_cmd_bin := exec.Command("curl", "-L", pdns_bin_url, "-o", "binary.bin")
	err = pdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdns_src_url := "https://downloads.powerdns.com/releases/pdns-4.9.1.tar.bz2"
	pdns_cmd_src := exec.Command("curl", "-L", pdns_src_url, "-o", "source.tar.gz")
	err = pdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdns_cmd_direct := exec.Command("./binary")
	err = pdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: ragel")
	exec.Command("latte", "install", "ragel").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
