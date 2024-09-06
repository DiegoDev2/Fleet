package main

// Cotp - TOTP/HOTP authenticator app with import functionality
// Homepage: https://github.com/replydev/cotp

import (
	"fmt"
	
	"os/exec"
)

func installCotp() {
	// Método 1: Descargar y extraer .tar.gz
	cotp_tar_url := "https://github.com/replydev/cotp/archive/refs/tags/v1.9.0.tar.gz"
	cotp_cmd_tar := exec.Command("curl", "-L", cotp_tar_url, "-o", "package.tar.gz")
	err := cotp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cotp_zip_url := "https://github.com/replydev/cotp/archive/refs/tags/v1.9.0.zip"
	cotp_cmd_zip := exec.Command("curl", "-L", cotp_zip_url, "-o", "package.zip")
	err = cotp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cotp_bin_url := "https://github.com/replydev/cotp/archive/refs/tags/v1.9.0.bin"
	cotp_cmd_bin := exec.Command("curl", "-L", cotp_bin_url, "-o", "binary.bin")
	err = cotp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cotp_src_url := "https://github.com/replydev/cotp/archive/refs/tags/v1.9.0.src.tar.gz"
	cotp_cmd_src := exec.Command("curl", "-L", cotp_src_url, "-o", "source.tar.gz")
	err = cotp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cotp_cmd_direct := exec.Command("./binary")
	err = cotp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
