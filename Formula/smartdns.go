package main

// Smartdns - Rule-based DNS server for fast IP resolution, DoT/DoQ/DoH/DoH3 supported
// Homepage: https://github.com/mokeyish/smartdns-rs

import (
	"fmt"
	
	"os/exec"
)

func installSmartdns() {
	// Método 1: Descargar y extraer .tar.gz
	smartdns_tar_url := "https://github.com/mokeyish/smartdns-rs/archive/refs/tags/v0.8.6.tar.gz"
	smartdns_cmd_tar := exec.Command("curl", "-L", smartdns_tar_url, "-o", "package.tar.gz")
	err := smartdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smartdns_zip_url := "https://github.com/mokeyish/smartdns-rs/archive/refs/tags/v0.8.6.zip"
	smartdns_cmd_zip := exec.Command("curl", "-L", smartdns_zip_url, "-o", "package.zip")
	err = smartdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smartdns_bin_url := "https://github.com/mokeyish/smartdns-rs/archive/refs/tags/v0.8.6.bin"
	smartdns_cmd_bin := exec.Command("curl", "-L", smartdns_bin_url, "-o", "binary.bin")
	err = smartdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smartdns_src_url := "https://github.com/mokeyish/smartdns-rs/archive/refs/tags/v0.8.6.src.tar.gz"
	smartdns_cmd_src := exec.Command("curl", "-L", smartdns_src_url, "-o", "source.tar.gz")
	err = smartdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smartdns_cmd_direct := exec.Command("./binary")
	err = smartdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: just")
	exec.Command("latte", "install", "just").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bind")
	exec.Command("latte", "install", "bind").Run()
}
