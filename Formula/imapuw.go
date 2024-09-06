package main

// ImapUw - University of Washington IMAP toolkit
// Homepage: https://salsa.debian.org/holmgren/uw-imap

import (
	"fmt"
	
	"os/exec"
)

func installImapUw() {
	// Método 1: Descargar y extraer .tar.gz
	imapuw_tar_url := "https://mirrorservice.org/sites/ftp.cac.washington.edu/imap/imap-2007f.tar.gz"
	imapuw_cmd_tar := exec.Command("curl", "-L", imapuw_tar_url, "-o", "package.tar.gz")
	err := imapuw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imapuw_zip_url := "https://mirrorservice.org/sites/ftp.cac.washington.edu/imap/imap-2007f.zip"
	imapuw_cmd_zip := exec.Command("curl", "-L", imapuw_zip_url, "-o", "package.zip")
	err = imapuw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imapuw_bin_url := "https://mirrorservice.org/sites/ftp.cac.washington.edu/imap/imap-2007f.bin"
	imapuw_cmd_bin := exec.Command("curl", "-L", imapuw_bin_url, "-o", "binary.bin")
	err = imapuw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imapuw_src_url := "https://mirrorservice.org/sites/ftp.cac.washington.edu/imap/imap-2007f.src.tar.gz"
	imapuw_cmd_src := exec.Command("curl", "-L", imapuw_src_url, "-o", "source.tar.gz")
	err = imapuw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imapuw_cmd_direct := exec.Command("./binary")
	err = imapuw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
}
