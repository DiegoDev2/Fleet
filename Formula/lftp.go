package main

// Lftp - Sophisticated file transfer program
// Homepage: https://lftp.yar.ru/

import (
	"fmt"
	
	"os/exec"
)

func installLftp() {
	// Método 1: Descargar y extraer .tar.gz
	lftp_tar_url := "https://lftp.yar.ru/ftp/lftp-4.9.2.tar.xz"
	lftp_cmd_tar := exec.Command("curl", "-L", lftp_tar_url, "-o", "package.tar.gz")
	err := lftp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lftp_zip_url := "https://lftp.yar.ru/ftp/lftp-4.9.2.tar.xz"
	lftp_cmd_zip := exec.Command("curl", "-L", lftp_zip_url, "-o", "package.zip")
	err = lftp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lftp_bin_url := "https://lftp.yar.ru/ftp/lftp-4.9.2.tar.xz"
	lftp_cmd_bin := exec.Command("curl", "-L", lftp_bin_url, "-o", "binary.bin")
	err = lftp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lftp_src_url := "https://lftp.yar.ru/ftp/lftp-4.9.2.tar.xz"
	lftp_cmd_src := exec.Command("curl", "-L", lftp_src_url, "-o", "source.tar.gz")
	err = lftp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lftp_cmd_direct := exec.Command("./binary")
	err = lftp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
