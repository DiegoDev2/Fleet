package main

// ImapBackup - Backup GMail (or other IMAP) accounts to disk
// Homepage: https://github.com/joeyates/imap-backup

import (
	"fmt"
	
	"os/exec"
)

func installImapBackup() {
	// Método 1: Descargar y extraer .tar.gz
	imapbackup_tar_url := "https://github.com/joeyates/imap-backup/archive/refs/tags/v15.0.2.tar.gz"
	imapbackup_cmd_tar := exec.Command("curl", "-L", imapbackup_tar_url, "-o", "package.tar.gz")
	err := imapbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imapbackup_zip_url := "https://github.com/joeyates/imap-backup/archive/refs/tags/v15.0.2.zip"
	imapbackup_cmd_zip := exec.Command("curl", "-L", imapbackup_zip_url, "-o", "package.zip")
	err = imapbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imapbackup_bin_url := "https://github.com/joeyates/imap-backup/archive/refs/tags/v15.0.2.bin"
	imapbackup_cmd_bin := exec.Command("curl", "-L", imapbackup_bin_url, "-o", "binary.bin")
	err = imapbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imapbackup_src_url := "https://github.com/joeyates/imap-backup/archive/refs/tags/v15.0.2.src.tar.gz"
	imapbackup_cmd_src := exec.Command("curl", "-L", imapbackup_src_url, "-o", "source.tar.gz")
	err = imapbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imapbackup_cmd_direct := exec.Command("./binary")
	err = imapbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
