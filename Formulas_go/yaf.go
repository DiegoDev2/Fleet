package main

// Yaf - Yet another flowmeter: processes packet data from pcap(3)
// Homepage: https://tools.netsa.cert.org/yaf/

import (
	"fmt"
	
	"os/exec"
)

func installYaf() {
	// Método 1: Descargar y extraer .tar.gz
	yaf_tar_url := "https://tools.netsa.cert.org/releases/yaf-2.16.0.tar.gz"
	yaf_cmd_tar := exec.Command("curl", "-L", yaf_tar_url, "-o", "package.tar.gz")
	err := yaf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yaf_zip_url := "https://tools.netsa.cert.org/releases/yaf-2.16.0.zip"
	yaf_cmd_zip := exec.Command("curl", "-L", yaf_zip_url, "-o", "package.zip")
	err = yaf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yaf_bin_url := "https://tools.netsa.cert.org/releases/yaf-2.16.0.bin"
	yaf_cmd_bin := exec.Command("curl", "-L", yaf_bin_url, "-o", "binary.bin")
	err = yaf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yaf_src_url := "https://tools.netsa.cert.org/releases/yaf-2.16.0.src.tar.gz"
	yaf_cmd_src := exec.Command("curl", "-L", yaf_src_url, "-o", "source.tar.gz")
	err = yaf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yaf_cmd_direct := exec.Command("./binary")
	err = yaf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libfixbuf")
exec.Command("latte", "install", "libfixbuf")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
