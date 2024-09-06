package main

// IscDhcp - Production-grade DHCP solution
// Homepage: https://www.isc.org/dhcp

import (
	"fmt"
	
	"os/exec"
)

func installIscDhcp() {
	// Método 1: Descargar y extraer .tar.gz
	iscdhcp_tar_url := "https://ftp.isc.org/isc/dhcp/4.4.3-P1/dhcp-4.4.3-P1.tar.gz"
	iscdhcp_cmd_tar := exec.Command("curl", "-L", iscdhcp_tar_url, "-o", "package.tar.gz")
	err := iscdhcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iscdhcp_zip_url := "https://ftp.isc.org/isc/dhcp/4.4.3-P1/dhcp-4.4.3-P1.zip"
	iscdhcp_cmd_zip := exec.Command("curl", "-L", iscdhcp_zip_url, "-o", "package.zip")
	err = iscdhcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iscdhcp_bin_url := "https://ftp.isc.org/isc/dhcp/4.4.3-P1/dhcp-4.4.3-P1.bin"
	iscdhcp_cmd_bin := exec.Command("curl", "-L", iscdhcp_bin_url, "-o", "binary.bin")
	err = iscdhcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iscdhcp_src_url := "https://ftp.isc.org/isc/dhcp/4.4.3-P1/dhcp-4.4.3-P1.src.tar.gz"
	iscdhcp_cmd_src := exec.Command("curl", "-L", iscdhcp_src_url, "-o", "source.tar.gz")
	err = iscdhcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iscdhcp_cmd_direct := exec.Command("./binary")
	err = iscdhcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
