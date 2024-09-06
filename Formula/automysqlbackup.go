package main

// Automysqlbackup - Automate MySQL backups
// Homepage: https://sourceforge.net/projects/automysqlbackup/

import (
	"fmt"
	
	"os/exec"
)

func installAutomysqlbackup() {
	// Método 1: Descargar y extraer .tar.gz
	automysqlbackup_tar_url := "https://downloads.sourceforge.net/project/automysqlbackup/AutoMySQLBackup/AutoMySQLBackup%20VER%203.0/automysqlbackup-v3.0_rc6.tar.gz"
	automysqlbackup_cmd_tar := exec.Command("curl", "-L", automysqlbackup_tar_url, "-o", "package.tar.gz")
	err := automysqlbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	automysqlbackup_zip_url := "https://downloads.sourceforge.net/project/automysqlbackup/AutoMySQLBackup/AutoMySQLBackup%20VER%203.0/automysqlbackup-v3.0_rc6.zip"
	automysqlbackup_cmd_zip := exec.Command("curl", "-L", automysqlbackup_zip_url, "-o", "package.zip")
	err = automysqlbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	automysqlbackup_bin_url := "https://downloads.sourceforge.net/project/automysqlbackup/AutoMySQLBackup/AutoMySQLBackup%20VER%203.0/automysqlbackup-v3.0_rc6.bin"
	automysqlbackup_cmd_bin := exec.Command("curl", "-L", automysqlbackup_bin_url, "-o", "binary.bin")
	err = automysqlbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	automysqlbackup_src_url := "https://downloads.sourceforge.net/project/automysqlbackup/AutoMySQLBackup/AutoMySQLBackup%20VER%203.0/automysqlbackup-v3.0_rc6.src.tar.gz"
	automysqlbackup_cmd_src := exec.Command("curl", "-L", automysqlbackup_src_url, "-o", "source.tar.gz")
	err = automysqlbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	automysqlbackup_cmd_direct := exec.Command("./binary")
	err = automysqlbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
