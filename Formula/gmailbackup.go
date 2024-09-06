package main

// GmailBackup - Backup and restore the content of your Gmail account
// Homepage: https://code.google.com/archive/p/gmail-backup-com/

import (
	"fmt"
	
	"os/exec"
)

func installGmailBackup() {
	// Método 1: Descargar y extraer .tar.gz
	gmailbackup_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/gmail-backup-com/gmail-backup-20110307.tar.gz"
	gmailbackup_cmd_tar := exec.Command("curl", "-L", gmailbackup_tar_url, "-o", "package.tar.gz")
	err := gmailbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmailbackup_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/gmail-backup-com/gmail-backup-20110307.zip"
	gmailbackup_cmd_zip := exec.Command("curl", "-L", gmailbackup_zip_url, "-o", "package.zip")
	err = gmailbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmailbackup_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/gmail-backup-com/gmail-backup-20110307.bin"
	gmailbackup_cmd_bin := exec.Command("curl", "-L", gmailbackup_bin_url, "-o", "binary.bin")
	err = gmailbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmailbackup_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/gmail-backup-com/gmail-backup-20110307.src.tar.gz"
	gmailbackup_cmd_src := exec.Command("curl", "-L", gmailbackup_src_url, "-o", "source.tar.gz")
	err = gmailbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmailbackup_cmd_direct := exec.Command("./binary")
	err = gmailbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
