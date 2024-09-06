package main

// PassOtp - Pass extension for managing one-time-password tokens
// Homepage: https://github.com/tadfisher/pass-otp

import (
	"fmt"
	
	"os/exec"
)

func installPassOtp() {
	// Método 1: Descargar y extraer .tar.gz
	passotp_tar_url := "https://github.com/tadfisher/pass-otp/releases/download/v1.2.0/pass-otp-1.2.0.tar.gz"
	passotp_cmd_tar := exec.Command("curl", "-L", passotp_tar_url, "-o", "package.tar.gz")
	err := passotp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	passotp_zip_url := "https://github.com/tadfisher/pass-otp/releases/download/v1.2.0/pass-otp-1.2.0.zip"
	passotp_cmd_zip := exec.Command("curl", "-L", passotp_zip_url, "-o", "package.zip")
	err = passotp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	passotp_bin_url := "https://github.com/tadfisher/pass-otp/releases/download/v1.2.0/pass-otp-1.2.0.bin"
	passotp_cmd_bin := exec.Command("curl", "-L", passotp_bin_url, "-o", "binary.bin")
	err = passotp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	passotp_src_url := "https://github.com/tadfisher/pass-otp/releases/download/v1.2.0/pass-otp-1.2.0.src.tar.gz"
	passotp_cmd_src := exec.Command("curl", "-L", passotp_src_url, "-o", "source.tar.gz")
	err = passotp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	passotp_cmd_direct := exec.Command("./binary")
	err = passotp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: oath-toolkit")
	exec.Command("latte", "install", "oath-toolkit").Run()
	fmt.Println("Instalando dependencia: pass")
	exec.Command("latte", "install", "pass").Run()
}
