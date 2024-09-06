package main

// Silk - Collection of traffic analysis tools
// Homepage: https://tools.netsa.cert.org/silk/

import (
	"fmt"
	
	"os/exec"
)

func installSilk() {
	// Método 1: Descargar y extraer .tar.gz
	silk_tar_url := "https://tools.netsa.cert.org/releases/silk-3.23.0.tar.gz"
	silk_cmd_tar := exec.Command("curl", "-L", silk_tar_url, "-o", "package.tar.gz")
	err := silk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	silk_zip_url := "https://tools.netsa.cert.org/releases/silk-3.23.0.zip"
	silk_cmd_zip := exec.Command("curl", "-L", silk_zip_url, "-o", "package.zip")
	err = silk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	silk_bin_url := "https://tools.netsa.cert.org/releases/silk-3.23.0.bin"
	silk_cmd_bin := exec.Command("curl", "-L", silk_bin_url, "-o", "binary.bin")
	err = silk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	silk_src_url := "https://tools.netsa.cert.org/releases/silk-3.23.0.src.tar.gz"
	silk_cmd_src := exec.Command("curl", "-L", silk_src_url, "-o", "source.tar.gz")
	err = silk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	silk_cmd_direct := exec.Command("./binary")
	err = silk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libfixbuf")
	exec.Command("latte", "install", "libfixbuf").Run()
	fmt.Println("Instalando dependencia: yaf")
	exec.Command("latte", "install", "yaf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
