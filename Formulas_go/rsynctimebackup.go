package main

// RsyncTimeBackup - Time Machine-style backup for the terminal using rsync
// Homepage: https://github.com/laurent22/rsync-time-backup

import (
	"fmt"
	
	"os/exec"
)

func installRsyncTimeBackup() {
	// Método 1: Descargar y extraer .tar.gz
	rsynctimebackup_tar_url := "https://github.com/laurent22/rsync-time-backup/archive/refs/tags/v1.1.5.tar.gz"
	rsynctimebackup_cmd_tar := exec.Command("curl", "-L", rsynctimebackup_tar_url, "-o", "package.tar.gz")
	err := rsynctimebackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsynctimebackup_zip_url := "https://github.com/laurent22/rsync-time-backup/archive/refs/tags/v1.1.5.zip"
	rsynctimebackup_cmd_zip := exec.Command("curl", "-L", rsynctimebackup_zip_url, "-o", "package.zip")
	err = rsynctimebackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsynctimebackup_bin_url := "https://github.com/laurent22/rsync-time-backup/archive/refs/tags/v1.1.5.bin"
	rsynctimebackup_cmd_bin := exec.Command("curl", "-L", rsynctimebackup_bin_url, "-o", "binary.bin")
	err = rsynctimebackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsynctimebackup_src_url := "https://github.com/laurent22/rsync-time-backup/archive/refs/tags/v1.1.5.src.tar.gz"
	rsynctimebackup_cmd_src := exec.Command("curl", "-L", rsynctimebackup_src_url, "-o", "source.tar.gz")
	err = rsynctimebackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsynctimebackup_cmd_direct := exec.Command("./binary")
	err = rsynctimebackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
