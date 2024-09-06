package main

// Dhcping - Perform a dhcp-request to check whether a dhcp-server is running
// Homepage: http://www.mavetju.org/unix/general.php

import (
	"fmt"
	
	"os/exec"
)

func installDhcping() {
	// Método 1: Descargar y extraer .tar.gz
	dhcping_tar_url := "http://www.mavetju.org/download/dhcping-1.2.tar.gz"
	dhcping_cmd_tar := exec.Command("curl", "-L", dhcping_tar_url, "-o", "package.tar.gz")
	err := dhcping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dhcping_zip_url := "http://www.mavetju.org/download/dhcping-1.2.zip"
	dhcping_cmd_zip := exec.Command("curl", "-L", dhcping_zip_url, "-o", "package.zip")
	err = dhcping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dhcping_bin_url := "http://www.mavetju.org/download/dhcping-1.2.bin"
	dhcping_cmd_bin := exec.Command("curl", "-L", dhcping_bin_url, "-o", "binary.bin")
	err = dhcping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dhcping_src_url := "http://www.mavetju.org/download/dhcping-1.2.src.tar.gz"
	dhcping_cmd_src := exec.Command("curl", "-L", dhcping_src_url, "-o", "source.tar.gz")
	err = dhcping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dhcping_cmd_direct := exec.Command("./binary")
	err = dhcping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
