package main

// EvernoteBackup - Backup & export all Evernote notes and notebooks
// Homepage: https://github.com/vzhd1701/evernote-backup

import (
	"fmt"
	
	"os/exec"
)

func installEvernoteBackup() {
	// Método 1: Descargar y extraer .tar.gz
	evernotebackup_tar_url := "https://files.pythonhosted.org/packages/81/04/6da56e51723acf47fa7c9148b80caef1be590ce82e5c1394e3faaff9d345/evernote_backup-1.9.3.tar.gz"
	evernotebackup_cmd_tar := exec.Command("curl", "-L", evernotebackup_tar_url, "-o", "package.tar.gz")
	err := evernotebackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	evernotebackup_zip_url := "https://files.pythonhosted.org/packages/81/04/6da56e51723acf47fa7c9148b80caef1be590ce82e5c1394e3faaff9d345/evernote_backup-1.9.3.zip"
	evernotebackup_cmd_zip := exec.Command("curl", "-L", evernotebackup_zip_url, "-o", "package.zip")
	err = evernotebackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	evernotebackup_bin_url := "https://files.pythonhosted.org/packages/81/04/6da56e51723acf47fa7c9148b80caef1be590ce82e5c1394e3faaff9d345/evernote_backup-1.9.3.bin"
	evernotebackup_cmd_bin := exec.Command("curl", "-L", evernotebackup_bin_url, "-o", "binary.bin")
	err = evernotebackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	evernotebackup_src_url := "https://files.pythonhosted.org/packages/81/04/6da56e51723acf47fa7c9148b80caef1be590ce82e5c1394e3faaff9d345/evernote_backup-1.9.3.src.tar.gz"
	evernotebackup_cmd_src := exec.Command("curl", "-L", evernotebackup_src_url, "-o", "source.tar.gz")
	err = evernotebackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	evernotebackup_cmd_direct := exec.Command("./binary")
	err = evernotebackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
