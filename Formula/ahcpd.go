package main

// Ahcpd - Autoconfiguration protocol for IPv6 and IPv6/IPv4 networks
// Homepage: https://www.irif.fr/~jch/software/ahcp/

import (
	"fmt"
	
	"os/exec"
)

func installAhcpd() {
	// Método 1: Descargar y extraer .tar.gz
	ahcpd_tar_url := "https://www.irif.fr/~jch/software/files/ahcpd-0.53.tar.gz"
	ahcpd_cmd_tar := exec.Command("curl", "-L", ahcpd_tar_url, "-o", "package.tar.gz")
	err := ahcpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ahcpd_zip_url := "https://www.irif.fr/~jch/software/files/ahcpd-0.53.zip"
	ahcpd_cmd_zip := exec.Command("curl", "-L", ahcpd_zip_url, "-o", "package.zip")
	err = ahcpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ahcpd_bin_url := "https://www.irif.fr/~jch/software/files/ahcpd-0.53.bin"
	ahcpd_cmd_bin := exec.Command("curl", "-L", ahcpd_bin_url, "-o", "binary.bin")
	err = ahcpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ahcpd_src_url := "https://www.irif.fr/~jch/software/files/ahcpd-0.53.src.tar.gz"
	ahcpd_cmd_src := exec.Command("curl", "-L", ahcpd_src_url, "-o", "source.tar.gz")
	err = ahcpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ahcpd_cmd_direct := exec.Command("./binary")
	err = ahcpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
