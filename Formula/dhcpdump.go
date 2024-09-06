package main

// Dhcpdump - Monitor DHCP traffic for debugging purposes
// Homepage: https://github.com/bbonev/dhcpdump

import (
	"fmt"
	
	"os/exec"
)

func installDhcpdump() {
	// Método 1: Descargar y extraer .tar.gz
	dhcpdump_tar_url := "https://github.com/bbonev/dhcpdump/releases/download/v1.9/dhcpdump-1.9.tar.xz"
	dhcpdump_cmd_tar := exec.Command("curl", "-L", dhcpdump_tar_url, "-o", "package.tar.gz")
	err := dhcpdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhcpdump_zip_url := "https://github.com/bbonev/dhcpdump/releases/download/v1.9/dhcpdump-1.9.tar.xz"
	dhcpdump_cmd_zip := exec.Command("curl", "-L", dhcpdump_zip_url, "-o", "package.zip")
	err = dhcpdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhcpdump_bin_url := "https://github.com/bbonev/dhcpdump/releases/download/v1.9/dhcpdump-1.9.tar.xz"
	dhcpdump_cmd_bin := exec.Command("curl", "-L", dhcpdump_bin_url, "-o", "binary.bin")
	err = dhcpdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhcpdump_src_url := "https://github.com/bbonev/dhcpdump/releases/download/v1.9/dhcpdump-1.9.tar.xz"
	dhcpdump_cmd_src := exec.Command("curl", "-L", dhcpdump_src_url, "-o", "source.tar.gz")
	err = dhcpdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhcpdump_cmd_direct := exec.Command("./binary")
	err = dhcpdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
