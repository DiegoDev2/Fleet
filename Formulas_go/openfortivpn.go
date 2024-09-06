package main

// Openfortivpn - Open Fortinet client for PPP+TLS VPN tunnel services
// Homepage: https://github.com/adrienverge/openfortivpn

import (
	"fmt"
	
	"os/exec"
)

func installOpenfortivpn() {
	// Método 1: Descargar y extraer .tar.gz
	openfortivpn_tar_url := "https://github.com/adrienverge/openfortivpn/archive/refs/tags/v1.22.1.tar.gz"
	openfortivpn_cmd_tar := exec.Command("curl", "-L", openfortivpn_tar_url, "-o", "package.tar.gz")
	err := openfortivpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openfortivpn_zip_url := "https://github.com/adrienverge/openfortivpn/archive/refs/tags/v1.22.1.zip"
	openfortivpn_cmd_zip := exec.Command("curl", "-L", openfortivpn_zip_url, "-o", "package.zip")
	err = openfortivpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openfortivpn_bin_url := "https://github.com/adrienverge/openfortivpn/archive/refs/tags/v1.22.1.bin"
	openfortivpn_cmd_bin := exec.Command("curl", "-L", openfortivpn_bin_url, "-o", "binary.bin")
	err = openfortivpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openfortivpn_src_url := "https://github.com/adrienverge/openfortivpn/archive/refs/tags/v1.22.1.src.tar.gz"
	openfortivpn_cmd_src := exec.Command("curl", "-L", openfortivpn_src_url, "-o", "source.tar.gz")
	err = openfortivpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openfortivpn_cmd_direct := exec.Command("./binary")
	err = openfortivpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
