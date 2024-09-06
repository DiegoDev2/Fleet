package main

// Inadyn - Dynamic DNS client with IPv4, IPv6, and SSL/TLS support
// Homepage: https://troglobit.com/projects/inadyn/

import (
	"fmt"
	
	"os/exec"
)

func installInadyn() {
	// Método 1: Descargar y extraer .tar.gz
	inadyn_tar_url := "https://github.com/troglobit/inadyn/releases/download/v2.12.0/inadyn-2.12.0.tar.xz"
	inadyn_cmd_tar := exec.Command("curl", "-L", inadyn_tar_url, "-o", "package.tar.gz")
	err := inadyn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inadyn_zip_url := "https://github.com/troglobit/inadyn/releases/download/v2.12.0/inadyn-2.12.0.tar.xz"
	inadyn_cmd_zip := exec.Command("curl", "-L", inadyn_zip_url, "-o", "package.zip")
	err = inadyn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inadyn_bin_url := "https://github.com/troglobit/inadyn/releases/download/v2.12.0/inadyn-2.12.0.tar.xz"
	inadyn_cmd_bin := exec.Command("curl", "-L", inadyn_bin_url, "-o", "binary.bin")
	err = inadyn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inadyn_src_url := "https://github.com/troglobit/inadyn/releases/download/v2.12.0/inadyn-2.12.0.tar.xz"
	inadyn_cmd_src := exec.Command("curl", "-L", inadyn_src_url, "-o", "source.tar.gz")
	err = inadyn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inadyn_cmd_direct := exec.Command("./binary")
	err = inadyn_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: confuse")
exec.Command("latte", "install", "confuse")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
}
