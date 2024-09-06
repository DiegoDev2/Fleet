package main

// Davmail - POP/IMAP/SMTP/Caldav/Carddav/LDAP exchange gateway
// Homepage: https://davmail.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDavmail() {
	// Método 1: Descargar y extraer .tar.gz
	davmail_tar_url := "https://downloads.sourceforge.net/project/davmail/davmail/6.2.2/davmail-6.2.2-3546.zip"
	davmail_cmd_tar := exec.Command("curl", "-L", davmail_tar_url, "-o", "package.tar.gz")
	err := davmail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	davmail_zip_url := "https://downloads.sourceforge.net/project/davmail/davmail/6.2.2/davmail-6.2.2-3546.zip"
	davmail_cmd_zip := exec.Command("curl", "-L", davmail_zip_url, "-o", "package.zip")
	err = davmail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	davmail_bin_url := "https://downloads.sourceforge.net/project/davmail/davmail/6.2.2/davmail-6.2.2-3546.zip"
	davmail_cmd_bin := exec.Command("curl", "-L", davmail_bin_url, "-o", "binary.bin")
	err = davmail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	davmail_src_url := "https://downloads.sourceforge.net/project/davmail/davmail/6.2.2/davmail-6.2.2-3546.zip"
	davmail_cmd_src := exec.Command("curl", "-L", davmail_src_url, "-o", "source.tar.gz")
	err = davmail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	davmail_cmd_direct := exec.Command("./binary")
	err = davmail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
