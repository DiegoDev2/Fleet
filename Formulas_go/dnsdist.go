package main

// Dnsdist - Highly DNS-, DoS- and abuse-aware loadbalancer
// Homepage: https://www.dnsdist.org/

import (
	"fmt"
	
	"os/exec"
)

func installDnsdist() {
	// Método 1: Descargar y extraer .tar.gz
	dnsdist_tar_url := "https://downloads.powerdns.com/releases/dnsdist-1.9.6.tar.bz2"
	dnsdist_cmd_tar := exec.Command("curl", "-L", dnsdist_tar_url, "-o", "package.tar.gz")
	err := dnsdist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsdist_zip_url := "https://downloads.powerdns.com/releases/dnsdist-1.9.6.tar.bz2"
	dnsdist_cmd_zip := exec.Command("curl", "-L", dnsdist_zip_url, "-o", "package.zip")
	err = dnsdist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsdist_bin_url := "https://downloads.powerdns.com/releases/dnsdist-1.9.6.tar.bz2"
	dnsdist_cmd_bin := exec.Command("curl", "-L", dnsdist_bin_url, "-o", "binary.bin")
	err = dnsdist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsdist_src_url := "https://downloads.powerdns.com/releases/dnsdist-1.9.6.tar.bz2"
	dnsdist_cmd_src := exec.Command("curl", "-L", dnsdist_src_url, "-o", "source.tar.gz")
	err = dnsdist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsdist_cmd_direct := exec.Command("./binary")
	err = dnsdist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: tinycdb")
exec.Command("latte", "install", "tinycdb")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: fstrm")
exec.Command("latte", "install", "fstrm")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: re2")
exec.Command("latte", "install", "re2")
}
