package main

// Tinc - Virtual Private Network (VPN) tool
// Homepage: https://www.tinc-vpn.org/

import (
	"fmt"
	
	"os/exec"
)

func installTinc() {
	// Método 1: Descargar y extraer .tar.gz
	tinc_tar_url := "https://tinc-vpn.org/packages/tinc-1.0.36.tar.gz"
	tinc_cmd_tar := exec.Command("curl", "-L", tinc_tar_url, "-o", "package.tar.gz")
	err := tinc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinc_zip_url := "https://tinc-vpn.org/packages/tinc-1.0.36.zip"
	tinc_cmd_zip := exec.Command("curl", "-L", tinc_zip_url, "-o", "package.zip")
	err = tinc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinc_bin_url := "https://tinc-vpn.org/packages/tinc-1.0.36.bin"
	tinc_cmd_bin := exec.Command("curl", "-L", tinc_bin_url, "-o", "binary.bin")
	err = tinc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinc_src_url := "https://tinc-vpn.org/packages/tinc-1.0.36.src.tar.gz"
	tinc_cmd_src := exec.Command("curl", "-L", tinc_src_url, "-o", "source.tar.gz")
	err = tinc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinc_cmd_direct := exec.Command("./binary")
	err = tinc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
