package main

// Backupninja - Backup automation tool
// Homepage: https://0xacab.org/liberate/backupninja

import (
	"fmt"
	
	"os/exec"
)

func installBackupninja() {
	// Método 1: Descargar y extraer .tar.gz
	backupninja_tar_url := "https://0xacab.org/liberate/backupninja/-/archive/backupninja_upstream/1.2.2/backupninja-backupninja_upstream-1.2.2.tar.gz"
	backupninja_cmd_tar := exec.Command("curl", "-L", backupninja_tar_url, "-o", "package.tar.gz")
	err := backupninja_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	backupninja_zip_url := "https://0xacab.org/liberate/backupninja/-/archive/backupninja_upstream/1.2.2/backupninja-backupninja_upstream-1.2.2.zip"
	backupninja_cmd_zip := exec.Command("curl", "-L", backupninja_zip_url, "-o", "package.zip")
	err = backupninja_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	backupninja_bin_url := "https://0xacab.org/liberate/backupninja/-/archive/backupninja_upstream/1.2.2/backupninja-backupninja_upstream-1.2.2.bin"
	backupninja_cmd_bin := exec.Command("curl", "-L", backupninja_bin_url, "-o", "binary.bin")
	err = backupninja_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	backupninja_src_url := "https://0xacab.org/liberate/backupninja/-/archive/backupninja_upstream/1.2.2/backupninja-backupninja_upstream-1.2.2.src.tar.gz"
	backupninja_cmd_src := exec.Command("curl", "-L", backupninja_src_url, "-o", "source.tar.gz")
	err = backupninja_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	backupninja_cmd_direct := exec.Command("./binary")
	err = backupninja_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
	fmt.Println("Instalando dependencia: dialog")
exec.Command("latte", "install", "dialog")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
}
