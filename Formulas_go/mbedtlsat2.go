package main

// MbedtlsAT2 - Cryptographic & SSL/TLS library
// Homepage: https://tls.mbed.org/

import (
	"fmt"
	
	"os/exec"
)

func installMbedtlsAT2() {
	// Método 1: Descargar y extraer .tar.gz
	mbedtlsat2_tar_url := "https://github.com/Mbed-TLS/mbedtls/archive/refs/tags/mbedtls-2.28.9.tar.gz"
	mbedtlsat2_cmd_tar := exec.Command("curl", "-L", mbedtlsat2_tar_url, "-o", "package.tar.gz")
	err := mbedtlsat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mbedtlsat2_zip_url := "https://github.com/Mbed-TLS/mbedtls/archive/refs/tags/mbedtls-2.28.9.zip"
	mbedtlsat2_cmd_zip := exec.Command("curl", "-L", mbedtlsat2_zip_url, "-o", "package.zip")
	err = mbedtlsat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mbedtlsat2_bin_url := "https://github.com/Mbed-TLS/mbedtls/archive/refs/tags/mbedtls-2.28.9.bin"
	mbedtlsat2_cmd_bin := exec.Command("curl", "-L", mbedtlsat2_bin_url, "-o", "binary.bin")
	err = mbedtlsat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mbedtlsat2_src_url := "https://github.com/Mbed-TLS/mbedtls/archive/refs/tags/mbedtls-2.28.9.src.tar.gz"
	mbedtlsat2_cmd_src := exec.Command("curl", "-L", mbedtlsat2_src_url, "-o", "source.tar.gz")
	err = mbedtlsat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mbedtlsat2_cmd_direct := exec.Command("./binary")
	err = mbedtlsat2_cmd_direct.Run()
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
