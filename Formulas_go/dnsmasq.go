package main

// Dnsmasq - Lightweight DNS forwarder and DHCP server
// Homepage: https://thekelleys.org.uk/dnsmasq/doc.html

import (
	"fmt"
	
	"os/exec"
)

func installDnsmasq() {
	// Método 1: Descargar y extraer .tar.gz
	dnsmasq_tar_url := "https://thekelleys.org.uk/dnsmasq/dnsmasq-2.90.tar.gz"
	dnsmasq_cmd_tar := exec.Command("curl", "-L", dnsmasq_tar_url, "-o", "package.tar.gz")
	err := dnsmasq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsmasq_zip_url := "https://thekelleys.org.uk/dnsmasq/dnsmasq-2.90.zip"
	dnsmasq_cmd_zip := exec.Command("curl", "-L", dnsmasq_zip_url, "-o", "package.zip")
	err = dnsmasq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsmasq_bin_url := "https://thekelleys.org.uk/dnsmasq/dnsmasq-2.90.bin"
	dnsmasq_cmd_bin := exec.Command("curl", "-L", dnsmasq_bin_url, "-o", "binary.bin")
	err = dnsmasq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsmasq_src_url := "https://thekelleys.org.uk/dnsmasq/dnsmasq-2.90.src.tar.gz"
	dnsmasq_cmd_src := exec.Command("curl", "-L", dnsmasq_src_url, "-o", "source.tar.gz")
	err = dnsmasq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsmasq_cmd_direct := exec.Command("./binary")
	err = dnsmasq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
