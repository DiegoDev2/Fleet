package main

// Libpcap - Portable library for network traffic capture
// Homepage: https://www.tcpdump.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibpcap() {
	// Método 1: Descargar y extraer .tar.gz
	libpcap_tar_url := "https://www.tcpdump.org/release/libpcap-1.10.5.tar.gz"
	libpcap_cmd_tar := exec.Command("curl", "-L", libpcap_tar_url, "-o", "package.tar.gz")
	err := libpcap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpcap_zip_url := "https://www.tcpdump.org/release/libpcap-1.10.5.zip"
	libpcap_cmd_zip := exec.Command("curl", "-L", libpcap_zip_url, "-o", "package.zip")
	err = libpcap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpcap_bin_url := "https://www.tcpdump.org/release/libpcap-1.10.5.bin"
	libpcap_cmd_bin := exec.Command("curl", "-L", libpcap_bin_url, "-o", "binary.bin")
	err = libpcap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpcap_src_url := "https://www.tcpdump.org/release/libpcap-1.10.5.src.tar.gz"
	libpcap_cmd_src := exec.Command("curl", "-L", libpcap_src_url, "-o", "source.tar.gz")
	err = libpcap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpcap_cmd_direct := exec.Command("./binary")
	err = libpcap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
