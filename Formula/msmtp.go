package main

// Msmtp - SMTP client that can be used as an SMTP plugin for Mutt
// Homepage: https://marlam.de/msmtp/

import (
	"fmt"
	
	"os/exec"
)

func installMsmtp() {
	// Método 1: Descargar y extraer .tar.gz
	msmtp_tar_url := "https://marlam.de/msmtp/releases/msmtp-1.8.26.tar.xz"
	msmtp_cmd_tar := exec.Command("curl", "-L", msmtp_tar_url, "-o", "package.tar.gz")
	err := msmtp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msmtp_zip_url := "https://marlam.de/msmtp/releases/msmtp-1.8.26.tar.xz"
	msmtp_cmd_zip := exec.Command("curl", "-L", msmtp_zip_url, "-o", "package.zip")
	err = msmtp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msmtp_bin_url := "https://marlam.de/msmtp/releases/msmtp-1.8.26.tar.xz"
	msmtp_cmd_bin := exec.Command("curl", "-L", msmtp_bin_url, "-o", "binary.bin")
	err = msmtp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msmtp_src_url := "https://marlam.de/msmtp/releases/msmtp-1.8.26.tar.xz"
	msmtp_cmd_src := exec.Command("curl", "-L", msmtp_src_url, "-o", "source.tar.gz")
	err = msmtp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msmtp_cmd_direct := exec.Command("./binary")
	err = msmtp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
}
