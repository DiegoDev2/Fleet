package main

// Openconnect - Open client for Cisco AnyConnect VPN
// Homepage: https://www.infradead.org/openconnect/

import (
	"fmt"
	
	"os/exec"
)

func installOpenconnect() {
	// Método 1: Descargar y extraer .tar.gz
	openconnect_tar_url := "https://www.infradead.org/openconnect/download/openconnect-9.12.tar.gz"
	openconnect_cmd_tar := exec.Command("curl", "-L", openconnect_tar_url, "-o", "package.tar.gz")
	err := openconnect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openconnect_zip_url := "https://www.infradead.org/openconnect/download/openconnect-9.12.zip"
	openconnect_cmd_zip := exec.Command("curl", "-L", openconnect_zip_url, "-o", "package.zip")
	err = openconnect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openconnect_bin_url := "https://www.infradead.org/openconnect/download/openconnect-9.12.bin"
	openconnect_cmd_bin := exec.Command("curl", "-L", openconnect_bin_url, "-o", "binary.bin")
	err = openconnect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openconnect_src_url := "https://www.infradead.org/openconnect/download/openconnect-9.12.src.tar.gz"
	openconnect_cmd_src := exec.Command("curl", "-L", openconnect_src_url, "-o", "source.tar.gz")
	err = openconnect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openconnect_cmd_direct := exec.Command("./binary")
	err = openconnect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: p11-kit")
exec.Command("latte", "install", "p11-kit")
	fmt.Println("Instalando dependencia: stoken")
exec.Command("latte", "install", "stoken")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
