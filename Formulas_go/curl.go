package main

// Curl - Get a file from an HTTP, HTTPS or FTP server
// Homepage: https://curl.se

import (
	"fmt"
	
	"os/exec"
)

func installCurl() {
	// Método 1: Descargar y extraer .tar.gz
	curl_tar_url := "https://curl.se/download/curl-8.9.1.tar.bz2"
	curl_cmd_tar := exec.Command("curl", "-L", curl_tar_url, "-o", "package.tar.gz")
	err := curl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curl_zip_url := "https://curl.se/download/curl-8.9.1.tar.bz2"
	curl_cmd_zip := exec.Command("curl", "-L", curl_zip_url, "-o", "package.zip")
	err = curl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curl_bin_url := "https://curl.se/download/curl-8.9.1.tar.bz2"
	curl_cmd_bin := exec.Command("curl", "-L", curl_bin_url, "-o", "binary.bin")
	err = curl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curl_src_url := "https://curl.se/download/curl-8.9.1.tar.bz2"
	curl_cmd_src := exec.Command("curl", "-L", curl_src_url, "-o", "source.tar.gz")
	err = curl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curl_cmd_direct := exec.Command("./binary")
	err = curl_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: rtmpdump")
exec.Command("latte", "install", "rtmpdump")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
