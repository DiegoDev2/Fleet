package main

// Cjdns - Advanced mesh routing system with cryptographic addressing
// Homepage: https://github.com/cjdelisle/cjdns/

import (
	"fmt"
	
	"os/exec"
)

func installCjdns() {
	// Método 1: Descargar y extraer .tar.gz
	cjdns_tar_url := "https://github.com/cjdelisle/cjdns/archive/refs/tags/cjdns-v22.tar.gz"
	cjdns_cmd_tar := exec.Command("curl", "-L", cjdns_tar_url, "-o", "package.tar.gz")
	err := cjdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cjdns_zip_url := "https://github.com/cjdelisle/cjdns/archive/refs/tags/cjdns-v22.zip"
	cjdns_cmd_zip := exec.Command("curl", "-L", cjdns_zip_url, "-o", "package.zip")
	err = cjdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cjdns_bin_url := "https://github.com/cjdelisle/cjdns/archive/refs/tags/cjdns-v22.bin"
	cjdns_cmd_bin := exec.Command("curl", "-L", cjdns_bin_url, "-o", "binary.bin")
	err = cjdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cjdns_src_url := "https://github.com/cjdelisle/cjdns/archive/refs/tags/cjdns-v22.src.tar.gz"
	cjdns_cmd_src := exec.Command("curl", "-L", cjdns_src_url, "-o", "source.tar.gz")
	err = cjdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cjdns_cmd_direct := exec.Command("./binary")
	err = cjdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
