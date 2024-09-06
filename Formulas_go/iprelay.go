package main

// IpRelay - TCP traffic shaping relay application
// Homepage: https://stewart.com.au/ip_relay/

import (
	"fmt"
	
	"os/exec"
)

func installIpRelay() {
	// Método 1: Descargar y extraer .tar.gz
	iprelay_tar_url := "https://stewart.com.au/ip_relay/ip_relay-0.71.tgz"
	iprelay_cmd_tar := exec.Command("curl", "-L", iprelay_tar_url, "-o", "package.tar.gz")
	err := iprelay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iprelay_zip_url := "https://stewart.com.au/ip_relay/ip_relay-0.71.tgz"
	iprelay_cmd_zip := exec.Command("curl", "-L", iprelay_zip_url, "-o", "package.zip")
	err = iprelay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iprelay_bin_url := "https://stewart.com.au/ip_relay/ip_relay-0.71.tgz"
	iprelay_cmd_bin := exec.Command("curl", "-L", iprelay_bin_url, "-o", "binary.bin")
	err = iprelay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iprelay_src_url := "https://stewart.com.au/ip_relay/ip_relay-0.71.tgz"
	iprelay_cmd_src := exec.Command("curl", "-L", iprelay_src_url, "-o", "source.tar.gz")
	err = iprelay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iprelay_cmd_direct := exec.Command("./binary")
	err = iprelay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
