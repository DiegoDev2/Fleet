package main

// Mailcheck - Check multiple mailboxes/maildirs for mail
// Homepage: https://mailcheck.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMailcheck() {
	// Método 1: Descargar y extraer .tar.gz
	mailcheck_tar_url := "https://downloads.sourceforge.net/project/mailcheck/mailcheck/1.91.2/mailcheck_1.91.2.tar.gz"
	mailcheck_cmd_tar := exec.Command("curl", "-L", mailcheck_tar_url, "-o", "package.tar.gz")
	err := mailcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mailcheck_zip_url := "https://downloads.sourceforge.net/project/mailcheck/mailcheck/1.91.2/mailcheck_1.91.2.zip"
	mailcheck_cmd_zip := exec.Command("curl", "-L", mailcheck_zip_url, "-o", "package.zip")
	err = mailcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mailcheck_bin_url := "https://downloads.sourceforge.net/project/mailcheck/mailcheck/1.91.2/mailcheck_1.91.2.bin"
	mailcheck_cmd_bin := exec.Command("curl", "-L", mailcheck_bin_url, "-o", "binary.bin")
	err = mailcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mailcheck_src_url := "https://downloads.sourceforge.net/project/mailcheck/mailcheck/1.91.2/mailcheck_1.91.2.src.tar.gz"
	mailcheck_cmd_src := exec.Command("curl", "-L", mailcheck_src_url, "-o", "source.tar.gz")
	err = mailcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mailcheck_cmd_direct := exec.Command("./binary")
	err = mailcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
