package main

// Dsvpn - Dead Simple VPN
// Homepage: https://github.com/jedisct1/dsvpn

import (
	"fmt"
	
	"os/exec"
)

func installDsvpn() {
	// Método 1: Descargar y extraer .tar.gz
	dsvpn_tar_url := "https://github.com/jedisct1/dsvpn/archive/refs/tags/0.1.4.tar.gz"
	dsvpn_cmd_tar := exec.Command("curl", "-L", dsvpn_tar_url, "-o", "package.tar.gz")
	err := dsvpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dsvpn_zip_url := "https://github.com/jedisct1/dsvpn/archive/refs/tags/0.1.4.zip"
	dsvpn_cmd_zip := exec.Command("curl", "-L", dsvpn_zip_url, "-o", "package.zip")
	err = dsvpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dsvpn_bin_url := "https://github.com/jedisct1/dsvpn/archive/refs/tags/0.1.4.bin"
	dsvpn_cmd_bin := exec.Command("curl", "-L", dsvpn_bin_url, "-o", "binary.bin")
	err = dsvpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dsvpn_src_url := "https://github.com/jedisct1/dsvpn/archive/refs/tags/0.1.4.src.tar.gz"
	dsvpn_cmd_src := exec.Command("curl", "-L", dsvpn_src_url, "-o", "source.tar.gz")
	err = dsvpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dsvpn_cmd_direct := exec.Command("./binary")
	err = dsvpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
