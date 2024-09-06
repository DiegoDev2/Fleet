package main

// Ntopng - Next generation version of the original ntop
// Homepage: https://www.ntop.org/products/traffic-analysis/ntop/

import (
	"fmt"
	
	"os/exec"
)

func installNtopng() {
	// Método 1: Descargar y extraer .tar.gz
	ntopng_tar_url := "https://github.com/ntop/ntopng/archive/refs/tags/5.2.1.tar.gz"
	ntopng_cmd_tar := exec.Command("curl", "-L", ntopng_tar_url, "-o", "package.tar.gz")
	err := ntopng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntopng_zip_url := "https://github.com/ntop/ntopng/archive/refs/tags/5.2.1.zip"
	ntopng_cmd_zip := exec.Command("curl", "-L", ntopng_zip_url, "-o", "package.zip")
	err = ntopng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntopng_bin_url := "https://github.com/ntop/ntopng/archive/refs/tags/5.2.1.bin"
	ntopng_cmd_bin := exec.Command("curl", "-L", ntopng_bin_url, "-o", "binary.bin")
	err = ntopng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntopng_src_url := "https://github.com/ntop/ntopng/archive/refs/tags/5.2.1.src.tar.gz"
	ntopng_cmd_src := exec.Command("curl", "-L", ntopng_src_url, "-o", "source.tar.gz")
	err = ntopng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntopng_cmd_direct := exec.Command("./binary")
	err = ntopng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ndpi")
exec.Command("latte", "install", "ndpi")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libmaxminddb")
exec.Command("latte", "install", "libmaxminddb")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: mysql-client")
exec.Command("latte", "install", "mysql-client")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: redis")
exec.Command("latte", "install", "redis")
	fmt.Println("Instalando dependencia: rrdtool")
exec.Command("latte", "install", "rrdtool")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
}
