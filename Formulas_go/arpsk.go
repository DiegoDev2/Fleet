package main

// ArpSk - ARP traffic generation tool
// Homepage: https://web.archive.org/web/20180223202629/sid.rstack.org/arp-sk/

import (
	"fmt"
	
	"os/exec"
)

func installArpSk() {
	// Método 1: Descargar y extraer .tar.gz
	arpsk_tar_url := "https://web.archive.org/web/20180223202629/sid.rstack.org/arp-sk/files/arp-sk-0.0.16.tgz"
	arpsk_cmd_tar := exec.Command("curl", "-L", arpsk_tar_url, "-o", "package.tar.gz")
	err := arpsk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arpsk_zip_url := "https://web.archive.org/web/20180223202629/sid.rstack.org/arp-sk/files/arp-sk-0.0.16.tgz"
	arpsk_cmd_zip := exec.Command("curl", "-L", arpsk_zip_url, "-o", "package.zip")
	err = arpsk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arpsk_bin_url := "https://web.archive.org/web/20180223202629/sid.rstack.org/arp-sk/files/arp-sk-0.0.16.tgz"
	arpsk_cmd_bin := exec.Command("curl", "-L", arpsk_bin_url, "-o", "binary.bin")
	err = arpsk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arpsk_src_url := "https://web.archive.org/web/20180223202629/sid.rstack.org/arp-sk/files/arp-sk-0.0.16.tgz"
	arpsk_cmd_src := exec.Command("curl", "-L", arpsk_src_url, "-o", "source.tar.gz")
	err = arpsk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arpsk_cmd_direct := exec.Command("./binary")
	err = arpsk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
}
