package main

// Vsftpd - Secure FTP server for UNIX
// Homepage: https://security.appspot.com/vsftpd.html

import (
	"fmt"
	
	"os/exec"
)

func installVsftpd() {
	// Método 1: Descargar y extraer .tar.gz
	vsftpd_tar_url := "https://security.appspot.com/downloads/vsftpd-3.0.5.tar.gz"
	vsftpd_cmd_tar := exec.Command("curl", "-L", vsftpd_tar_url, "-o", "package.tar.gz")
	err := vsftpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vsftpd_zip_url := "https://security.appspot.com/downloads/vsftpd-3.0.5.zip"
	vsftpd_cmd_zip := exec.Command("curl", "-L", vsftpd_zip_url, "-o", "package.zip")
	err = vsftpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vsftpd_bin_url := "https://security.appspot.com/downloads/vsftpd-3.0.5.bin"
	vsftpd_cmd_bin := exec.Command("curl", "-L", vsftpd_bin_url, "-o", "binary.bin")
	err = vsftpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vsftpd_src_url := "https://security.appspot.com/downloads/vsftpd-3.0.5.src.tar.gz"
	vsftpd_cmd_src := exec.Command("curl", "-L", vsftpd_src_url, "-o", "source.tar.gz")
	err = vsftpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vsftpd_cmd_direct := exec.Command("./binary")
	err = vsftpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
