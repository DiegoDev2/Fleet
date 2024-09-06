package main

// PureFtpd - Secure and efficient FTP server
// Homepage: https://www.pureftpd.org/

import (
	"fmt"
	
	"os/exec"
)

func installPureFtpd() {
	// Método 1: Descargar y extraer .tar.gz
	pureftpd_tar_url := "https://download.pureftpd.org/pub/pure-ftpd/releases/pure-ftpd-1.0.51.tar.gz"
	pureftpd_cmd_tar := exec.Command("curl", "-L", pureftpd_tar_url, "-o", "package.tar.gz")
	err := pureftpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pureftpd_zip_url := "https://download.pureftpd.org/pub/pure-ftpd/releases/pure-ftpd-1.0.51.zip"
	pureftpd_cmd_zip := exec.Command("curl", "-L", pureftpd_zip_url, "-o", "package.zip")
	err = pureftpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pureftpd_bin_url := "https://download.pureftpd.org/pub/pure-ftpd/releases/pure-ftpd-1.0.51.bin"
	pureftpd_cmd_bin := exec.Command("curl", "-L", pureftpd_bin_url, "-o", "binary.bin")
	err = pureftpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pureftpd_src_url := "https://download.pureftpd.org/pub/pure-ftpd/releases/pure-ftpd-1.0.51.src.tar.gz"
	pureftpd_cmd_src := exec.Command("curl", "-L", pureftpd_src_url, "-o", "source.tar.gz")
	err = pureftpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pureftpd_cmd_direct := exec.Command("./binary")
	err = pureftpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
}
