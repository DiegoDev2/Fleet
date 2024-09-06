package main

// ShallowBackup - Git-integrated backup tool for macOS and Linux devs
// Homepage: https://github.com/alichtman/shallow-backup

import (
	"fmt"
	
	"os/exec"
)

func installShallowBackup() {
	// Método 1: Descargar y extraer .tar.gz
	shallowbackup_tar_url := "https://files.pythonhosted.org/packages/d0/56/427960ea933c35b43b561d8f1379d4f7794b67f785ec3137adaf6ce5073e/shallow_backup-6.4.tar.gz"
	shallowbackup_cmd_tar := exec.Command("curl", "-L", shallowbackup_tar_url, "-o", "package.tar.gz")
	err := shallowbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shallowbackup_zip_url := "https://files.pythonhosted.org/packages/d0/56/427960ea933c35b43b561d8f1379d4f7794b67f785ec3137adaf6ce5073e/shallow_backup-6.4.zip"
	shallowbackup_cmd_zip := exec.Command("curl", "-L", shallowbackup_zip_url, "-o", "package.zip")
	err = shallowbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shallowbackup_bin_url := "https://files.pythonhosted.org/packages/d0/56/427960ea933c35b43b561d8f1379d4f7794b67f785ec3137adaf6ce5073e/shallow_backup-6.4.bin"
	shallowbackup_cmd_bin := exec.Command("curl", "-L", shallowbackup_bin_url, "-o", "binary.bin")
	err = shallowbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shallowbackup_src_url := "https://files.pythonhosted.org/packages/d0/56/427960ea933c35b43b561d8f1379d4f7794b67f785ec3137adaf6ce5073e/shallow_backup-6.4.src.tar.gz"
	shallowbackup_cmd_src := exec.Command("curl", "-L", shallowbackup_src_url, "-o", "source.tar.gz")
	err = shallowbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shallowbackup_cmd_direct := exec.Command("./binary")
	err = shallowbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
