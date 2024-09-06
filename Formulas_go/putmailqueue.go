package main

// PutmailQueue - Queue package for putmail
// Homepage: https://putmail.sourceforge.net/home.html

import (
	"fmt"
	
	"os/exec"
)

func installPutmailQueue() {
	// Método 1: Descargar y extraer .tar.gz
	putmailqueue_tar_url := "https://downloads.sourceforge.net/project/putmail/putmail-queue/0.2/putmail-queue-0.2.tar.bz2"
	putmailqueue_cmd_tar := exec.Command("curl", "-L", putmailqueue_tar_url, "-o", "package.tar.gz")
	err := putmailqueue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	putmailqueue_zip_url := "https://downloads.sourceforge.net/project/putmail/putmail-queue/0.2/putmail-queue-0.2.tar.bz2"
	putmailqueue_cmd_zip := exec.Command("curl", "-L", putmailqueue_zip_url, "-o", "package.zip")
	err = putmailqueue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	putmailqueue_bin_url := "https://downloads.sourceforge.net/project/putmail/putmail-queue/0.2/putmail-queue-0.2.tar.bz2"
	putmailqueue_cmd_bin := exec.Command("curl", "-L", putmailqueue_bin_url, "-o", "binary.bin")
	err = putmailqueue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	putmailqueue_src_url := "https://downloads.sourceforge.net/project/putmail/putmail-queue/0.2/putmail-queue-0.2.tar.bz2"
	putmailqueue_cmd_src := exec.Command("curl", "-L", putmailqueue_src_url, "-o", "source.tar.gz")
	err = putmailqueue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	putmailqueue_cmd_direct := exec.Command("./binary")
	err = putmailqueue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: putmail")
exec.Command("latte", "install", "putmail")
}
