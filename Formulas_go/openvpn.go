package main

// Openvpn - SSL/TLS VPN implementing OSI layer 2 or 3 secure network extension
// Homepage: https://openvpn.net/community/

import (
	"fmt"
	
	"os/exec"
)

func installOpenvpn() {
	// Método 1: Descargar y extraer .tar.gz
	openvpn_tar_url := "https://swupdate.openvpn.org/community/releases/openvpn-2.6.12.tar.gz"
	openvpn_cmd_tar := exec.Command("curl", "-L", openvpn_tar_url, "-o", "package.tar.gz")
	err := openvpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openvpn_zip_url := "https://swupdate.openvpn.org/community/releases/openvpn-2.6.12.zip"
	openvpn_cmd_zip := exec.Command("curl", "-L", openvpn_zip_url, "-o", "package.zip")
	err = openvpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openvpn_bin_url := "https://swupdate.openvpn.org/community/releases/openvpn-2.6.12.bin"
	openvpn_cmd_bin := exec.Command("curl", "-L", openvpn_bin_url, "-o", "binary.bin")
	err = openvpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openvpn_src_url := "https://swupdate.openvpn.org/community/releases/openvpn-2.6.12.src.tar.gz"
	openvpn_cmd_src := exec.Command("curl", "-L", openvpn_src_url, "-o", "source.tar.gz")
	err = openvpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openvpn_cmd_direct := exec.Command("./binary")
	err = openvpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkcs11-helper")
exec.Command("latte", "install", "pkcs11-helper")
	fmt.Println("Instalando dependencia: libcap-ng")
exec.Command("latte", "install", "libcap-ng")
	fmt.Println("Instalando dependencia: libnl")
exec.Command("latte", "install", "libnl")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
	fmt.Println("Instalando dependencia: net-tools")
exec.Command("latte", "install", "net-tools")
}
