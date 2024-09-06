package main

// Mbedtls - Cryptographic & SSL/TLS library
// Homepage: https://tls.mbed.org/

import (
	"fmt"
	
	"os/exec"
)

func installMbedtls() {
	// Método 1: Descargar y extraer .tar.gz
	mbedtls_tar_url := "https://github.com/Mbed-TLS/mbedtls/releases/download/mbedtls-3.6.1/mbedtls-3.6.1.tar.bz2"
	mbedtls_cmd_tar := exec.Command("curl", "-L", mbedtls_tar_url, "-o", "package.tar.gz")
	err := mbedtls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mbedtls_zip_url := "https://github.com/Mbed-TLS/mbedtls/releases/download/mbedtls-3.6.1/mbedtls-3.6.1.tar.bz2"
	mbedtls_cmd_zip := exec.Command("curl", "-L", mbedtls_zip_url, "-o", "package.zip")
	err = mbedtls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mbedtls_bin_url := "https://github.com/Mbed-TLS/mbedtls/releases/download/mbedtls-3.6.1/mbedtls-3.6.1.tar.bz2"
	mbedtls_cmd_bin := exec.Command("curl", "-L", mbedtls_bin_url, "-o", "binary.bin")
	err = mbedtls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mbedtls_src_url := "https://github.com/Mbed-TLS/mbedtls/releases/download/mbedtls-3.6.1/mbedtls-3.6.1.tar.bz2"
	mbedtls_cmd_src := exec.Command("curl", "-L", mbedtls_src_url, "-o", "source.tar.gz")
	err = mbedtls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mbedtls_cmd_direct := exec.Command("./binary")
	err = mbedtls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
