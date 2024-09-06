package main

// Sendemail - Email program for sending SMTP mail
// Homepage: http://caspian.dotconf.net/menu/Software/SendEmail/

import (
	"fmt"
	
	"os/exec"
)

func installSendemail() {
	// Método 1: Descargar y extraer .tar.gz
	sendemail_tar_url := "http://caspian.dotconf.net/menu/Software/SendEmail/sendEmail-v1.56.tar.gz"
	sendemail_cmd_tar := exec.Command("curl", "-L", sendemail_tar_url, "-o", "package.tar.gz")
	err := sendemail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sendemail_zip_url := "http://caspian.dotconf.net/menu/Software/SendEmail/sendEmail-v1.56.zip"
	sendemail_cmd_zip := exec.Command("curl", "-L", sendemail_zip_url, "-o", "package.zip")
	err = sendemail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sendemail_bin_url := "http://caspian.dotconf.net/menu/Software/SendEmail/sendEmail-v1.56.bin"
	sendemail_cmd_bin := exec.Command("curl", "-L", sendemail_bin_url, "-o", "binary.bin")
	err = sendemail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sendemail_src_url := "http://caspian.dotconf.net/menu/Software/SendEmail/sendEmail-v1.56.src.tar.gz"
	sendemail_cmd_src := exec.Command("curl", "-L", sendemail_src_url, "-o", "source.tar.gz")
	err = sendemail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sendemail_cmd_direct := exec.Command("./binary")
	err = sendemail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
