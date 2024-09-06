package main

// Nextdns - CLI for NextDNS's DNS-over-HTTPS (DoH)
// Homepage: https://nextdns.io

import (
	"fmt"
	
	"os/exec"
)

func installNextdns() {
	// Método 1: Descargar y extraer .tar.gz
	nextdns_tar_url := "https://github.com/nextdns/nextdns/archive/refs/tags/v1.43.5.tar.gz"
	nextdns_cmd_tar := exec.Command("curl", "-L", nextdns_tar_url, "-o", "package.tar.gz")
	err := nextdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nextdns_zip_url := "https://github.com/nextdns/nextdns/archive/refs/tags/v1.43.5.zip"
	nextdns_cmd_zip := exec.Command("curl", "-L", nextdns_zip_url, "-o", "package.zip")
	err = nextdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nextdns_bin_url := "https://github.com/nextdns/nextdns/archive/refs/tags/v1.43.5.bin"
	nextdns_cmd_bin := exec.Command("curl", "-L", nextdns_bin_url, "-o", "binary.bin")
	err = nextdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nextdns_src_url := "https://github.com/nextdns/nextdns/archive/refs/tags/v1.43.5.src.tar.gz"
	nextdns_cmd_src := exec.Command("curl", "-L", nextdns_src_url, "-o", "source.tar.gz")
	err = nextdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nextdns_cmd_direct := exec.Command("./binary")
	err = nextdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
