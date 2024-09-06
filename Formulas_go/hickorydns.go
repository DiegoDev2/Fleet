package main

// HickoryDns - Rust based DNS client, server, and resolver
// Homepage: https://github.com/hickory-dns/hickory-dns

import (
	"fmt"
	
	"os/exec"
)

func installHickoryDns() {
	// Método 1: Descargar y extraer .tar.gz
	hickorydns_tar_url := "https://github.com/hickory-dns/hickory-dns/archive/refs/tags/v0.24.1.tar.gz"
	hickorydns_cmd_tar := exec.Command("curl", "-L", hickorydns_tar_url, "-o", "package.tar.gz")
	err := hickorydns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hickorydns_zip_url := "https://github.com/hickory-dns/hickory-dns/archive/refs/tags/v0.24.1.zip"
	hickorydns_cmd_zip := exec.Command("curl", "-L", hickorydns_zip_url, "-o", "package.zip")
	err = hickorydns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hickorydns_bin_url := "https://github.com/hickory-dns/hickory-dns/archive/refs/tags/v0.24.1.bin"
	hickorydns_cmd_bin := exec.Command("curl", "-L", hickorydns_bin_url, "-o", "binary.bin")
	err = hickorydns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hickorydns_src_url := "https://github.com/hickory-dns/hickory-dns/archive/refs/tags/v0.24.1.src.tar.gz"
	hickorydns_cmd_src := exec.Command("curl", "-L", hickorydns_src_url, "-o", "source.tar.gz")
	err = hickorydns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hickorydns_cmd_direct := exec.Command("./binary")
	err = hickorydns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: bind")
exec.Command("latte", "install", "bind")
}
