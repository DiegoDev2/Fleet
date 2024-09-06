package main

// OpensslAT30 - Cryptography and SSL/TLS Toolkit
// Homepage: https://openssl-library.org

import (
	"fmt"
	
	"os/exec"
)

func installOpensslAT30() {
	// Método 1: Descargar y extraer .tar.gz
	opensslat30_tar_url := "https://github.com/openssl/openssl/releases/download/openssl-3.0.15/openssl-3.0.15.tar.gz"
	opensslat30_cmd_tar := exec.Command("curl", "-L", opensslat30_tar_url, "-o", "package.tar.gz")
	err := opensslat30_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensslat30_zip_url := "https://github.com/openssl/openssl/releases/download/openssl-3.0.15/openssl-3.0.15.zip"
	opensslat30_cmd_zip := exec.Command("curl", "-L", opensslat30_zip_url, "-o", "package.zip")
	err = opensslat30_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensslat30_bin_url := "https://github.com/openssl/openssl/releases/download/openssl-3.0.15/openssl-3.0.15.bin"
	opensslat30_cmd_bin := exec.Command("curl", "-L", opensslat30_bin_url, "-o", "binary.bin")
	err = opensslat30_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensslat30_src_url := "https://github.com/openssl/openssl/releases/download/openssl-3.0.15/openssl-3.0.15.src.tar.gz"
	opensslat30_cmd_src := exec.Command("curl", "-L", opensslat30_src_url, "-o", "source.tar.gz")
	err = opensslat30_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensslat30_cmd_direct := exec.Command("./binary")
	err = opensslat30_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
}
