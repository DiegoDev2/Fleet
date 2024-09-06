package main

// OpensslAT11 - Cryptography and SSL/TLS Toolkit
// Homepage: https://openssl.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpensslAT11() {
	// Método 1: Descargar y extraer .tar.gz
	opensslat11_tar_url := "https://www.openssl.org/source/openssl-1.1.1w.tar.gz"
	opensslat11_cmd_tar := exec.Command("curl", "-L", opensslat11_tar_url, "-o", "package.tar.gz")
	err := opensslat11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensslat11_zip_url := "https://www.openssl.org/source/openssl-1.1.1w.zip"
	opensslat11_cmd_zip := exec.Command("curl", "-L", opensslat11_zip_url, "-o", "package.zip")
	err = opensslat11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensslat11_bin_url := "https://www.openssl.org/source/openssl-1.1.1w.bin"
	opensslat11_cmd_bin := exec.Command("curl", "-L", opensslat11_bin_url, "-o", "binary.bin")
	err = opensslat11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensslat11_src_url := "https://www.openssl.org/source/openssl-1.1.1w.src.tar.gz"
	opensslat11_cmd_src := exec.Command("curl", "-L", opensslat11_src_url, "-o", "source.tar.gz")
	err = opensslat11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensslat11_cmd_direct := exec.Command("./binary")
	err = opensslat11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
}
