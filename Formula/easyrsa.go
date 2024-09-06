package main

// EasyRsa - CLI utility to build and manage a PKI CA
// Homepage: https://github.com/OpenVPN/easy-rsa

import (
	"fmt"
	
	"os/exec"
)

func installEasyRsa() {
	// Método 1: Descargar y extraer .tar.gz
	easyrsa_tar_url := "https://github.com/OpenVPN/easy-rsa/releases/download/v3.2.0/EasyRSA-3.2.0.tgz"
	easyrsa_cmd_tar := exec.Command("curl", "-L", easyrsa_tar_url, "-o", "package.tar.gz")
	err := easyrsa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	easyrsa_zip_url := "https://github.com/OpenVPN/easy-rsa/releases/download/v3.2.0/EasyRSA-3.2.0.tgz"
	easyrsa_cmd_zip := exec.Command("curl", "-L", easyrsa_zip_url, "-o", "package.zip")
	err = easyrsa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	easyrsa_bin_url := "https://github.com/OpenVPN/easy-rsa/releases/download/v3.2.0/EasyRSA-3.2.0.tgz"
	easyrsa_cmd_bin := exec.Command("curl", "-L", easyrsa_bin_url, "-o", "binary.bin")
	err = easyrsa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	easyrsa_src_url := "https://github.com/OpenVPN/easy-rsa/releases/download/v3.2.0/EasyRSA-3.2.0.tgz"
	easyrsa_cmd_src := exec.Command("curl", "-L", easyrsa_src_url, "-o", "source.tar.gz")
	err = easyrsa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	easyrsa_cmd_direct := exec.Command("./binary")
	err = easyrsa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
