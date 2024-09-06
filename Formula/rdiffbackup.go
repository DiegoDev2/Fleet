package main

// RdiffBackup - Reverse differential backup tool, over a network or locally
// Homepage: https://rdiff-backup.net/

import (
	"fmt"
	
	"os/exec"
)

func installRdiffBackup() {
	// Método 1: Descargar y extraer .tar.gz
	rdiffbackup_tar_url := "https://files.pythonhosted.org/packages/e9/9b/487229306904a54c33f485161105bb3f0a6c87951c90a54efdc0fc04a1c9/rdiff-backup-2.2.6.tar.gz"
	rdiffbackup_cmd_tar := exec.Command("curl", "-L", rdiffbackup_tar_url, "-o", "package.tar.gz")
	err := rdiffbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rdiffbackup_zip_url := "https://files.pythonhosted.org/packages/e9/9b/487229306904a54c33f485161105bb3f0a6c87951c90a54efdc0fc04a1c9/rdiff-backup-2.2.6.zip"
	rdiffbackup_cmd_zip := exec.Command("curl", "-L", rdiffbackup_zip_url, "-o", "package.zip")
	err = rdiffbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rdiffbackup_bin_url := "https://files.pythonhosted.org/packages/e9/9b/487229306904a54c33f485161105bb3f0a6c87951c90a54efdc0fc04a1c9/rdiff-backup-2.2.6.bin"
	rdiffbackup_cmd_bin := exec.Command("curl", "-L", rdiffbackup_bin_url, "-o", "binary.bin")
	err = rdiffbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rdiffbackup_src_url := "https://files.pythonhosted.org/packages/e9/9b/487229306904a54c33f485161105bb3f0a6c87951c90a54efdc0fc04a1c9/rdiff-backup-2.2.6.src.tar.gz"
	rdiffbackup_cmd_src := exec.Command("curl", "-L", rdiffbackup_src_url, "-o", "source.tar.gz")
	err = rdiffbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rdiffbackup_cmd_direct := exec.Command("./binary")
	err = rdiffbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: librsync")
	exec.Command("latte", "install", "librsync").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
