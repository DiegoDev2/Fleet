package main

// Srtp - Implementation of the Secure Real-time Transport Protocol
// Homepage: https://github.com/cisco/libsrtp

import (
	"fmt"
	
	"os/exec"
)

func installSrtp() {
	// Método 1: Descargar y extraer .tar.gz
	srtp_tar_url := "https://github.com/cisco/libsrtp/archive/refs/tags/v2.6.0.tar.gz"
	srtp_cmd_tar := exec.Command("curl", "-L", srtp_tar_url, "-o", "package.tar.gz")
	err := srtp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	srtp_zip_url := "https://github.com/cisco/libsrtp/archive/refs/tags/v2.6.0.zip"
	srtp_cmd_zip := exec.Command("curl", "-L", srtp_zip_url, "-o", "package.zip")
	err = srtp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	srtp_bin_url := "https://github.com/cisco/libsrtp/archive/refs/tags/v2.6.0.bin"
	srtp_cmd_bin := exec.Command("curl", "-L", srtp_bin_url, "-o", "binary.bin")
	err = srtp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	srtp_src_url := "https://github.com/cisco/libsrtp/archive/refs/tags/v2.6.0.src.tar.gz"
	srtp_cmd_src := exec.Command("curl", "-L", srtp_src_url, "-o", "source.tar.gz")
	err = srtp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	srtp_cmd_direct := exec.Command("./binary")
	err = srtp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
