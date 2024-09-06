package main

// Ipv6calc - Small utility for manipulating IPv6 addresses
// Homepage: https://www.deepspace6.net/projects/ipv6calc.html

import (
	"fmt"
	
	"os/exec"
)

func installIpv6calc() {
	// Método 1: Descargar y extraer .tar.gz
	ipv6calc_tar_url := "https://github.com/pbiering/ipv6calc/archive/refs/tags/4.2.1.tar.gz"
	ipv6calc_cmd_tar := exec.Command("curl", "-L", ipv6calc_tar_url, "-o", "package.tar.gz")
	err := ipv6calc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipv6calc_zip_url := "https://github.com/pbiering/ipv6calc/archive/refs/tags/4.2.1.zip"
	ipv6calc_cmd_zip := exec.Command("curl", "-L", ipv6calc_zip_url, "-o", "package.zip")
	err = ipv6calc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipv6calc_bin_url := "https://github.com/pbiering/ipv6calc/archive/refs/tags/4.2.1.bin"
	ipv6calc_cmd_bin := exec.Command("curl", "-L", ipv6calc_bin_url, "-o", "binary.bin")
	err = ipv6calc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipv6calc_src_url := "https://github.com/pbiering/ipv6calc/archive/refs/tags/4.2.1.src.tar.gz"
	ipv6calc_cmd_src := exec.Command("curl", "-L", ipv6calc_src_url, "-o", "source.tar.gz")
	err = ipv6calc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipv6calc_cmd_direct := exec.Command("./binary")
	err = ipv6calc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
