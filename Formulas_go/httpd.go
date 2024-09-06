package main

// Httpd - Apache HTTP server
// Homepage: https://httpd.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installHttpd() {
	// Método 1: Descargar y extraer .tar.gz
	httpd_tar_url := "https://dlcdn.apache.org/httpd/httpd-2.4.62.tar.bz2"
	httpd_cmd_tar := exec.Command("curl", "-L", httpd_tar_url, "-o", "package.tar.gz")
	err := httpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpd_zip_url := "https://dlcdn.apache.org/httpd/httpd-2.4.62.tar.bz2"
	httpd_cmd_zip := exec.Command("curl", "-L", httpd_zip_url, "-o", "package.zip")
	err = httpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpd_bin_url := "https://dlcdn.apache.org/httpd/httpd-2.4.62.tar.bz2"
	httpd_cmd_bin := exec.Command("curl", "-L", httpd_bin_url, "-o", "binary.bin")
	err = httpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpd_src_url := "https://dlcdn.apache.org/httpd/httpd-2.4.62.tar.bz2"
	httpd_cmd_src := exec.Command("curl", "-L", httpd_src_url, "-o", "source.tar.gz")
	err = httpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpd_cmd_direct := exec.Command("./binary")
	err = httpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
