package main

// Ipv6toolkit - Security assessment and troubleshooting tool for IPv6
// Homepage: https://www.si6networks.com/research/tools/ipv6toolkit/

import (
	"fmt"
	
	"os/exec"
)

func installIpv6toolkit() {
	// Método 1: Descargar y extraer .tar.gz
	ipv6toolkit_tar_url := "https://pages.cs.wisc.edu/~plonka/ipv6toolkit/ipv6toolkit-v2.0.tar.gz"
	ipv6toolkit_cmd_tar := exec.Command("curl", "-L", ipv6toolkit_tar_url, "-o", "package.tar.gz")
	err := ipv6toolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipv6toolkit_zip_url := "https://pages.cs.wisc.edu/~plonka/ipv6toolkit/ipv6toolkit-v2.0.zip"
	ipv6toolkit_cmd_zip := exec.Command("curl", "-L", ipv6toolkit_zip_url, "-o", "package.zip")
	err = ipv6toolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipv6toolkit_bin_url := "https://pages.cs.wisc.edu/~plonka/ipv6toolkit/ipv6toolkit-v2.0.bin"
	ipv6toolkit_cmd_bin := exec.Command("curl", "-L", ipv6toolkit_bin_url, "-o", "binary.bin")
	err = ipv6toolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipv6toolkit_src_url := "https://pages.cs.wisc.edu/~plonka/ipv6toolkit/ipv6toolkit-v2.0.src.tar.gz"
	ipv6toolkit_cmd_src := exec.Command("curl", "-L", ipv6toolkit_src_url, "-o", "source.tar.gz")
	err = ipv6toolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipv6toolkit_cmd_direct := exec.Command("./binary")
	err = ipv6toolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
