package main

// Rkhunter - Rootkit hunter
// Homepage: https://rkhunter.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installRkhunter() {
	// Método 1: Descargar y extraer .tar.gz
	rkhunter_tar_url := "https://downloads.sourceforge.net/project/rkhunter/rkhunter/1.4.6/rkhunter-1.4.6.tar.gz"
	rkhunter_cmd_tar := exec.Command("curl", "-L", rkhunter_tar_url, "-o", "package.tar.gz")
	err := rkhunter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rkhunter_zip_url := "https://downloads.sourceforge.net/project/rkhunter/rkhunter/1.4.6/rkhunter-1.4.6.zip"
	rkhunter_cmd_zip := exec.Command("curl", "-L", rkhunter_zip_url, "-o", "package.zip")
	err = rkhunter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rkhunter_bin_url := "https://downloads.sourceforge.net/project/rkhunter/rkhunter/1.4.6/rkhunter-1.4.6.bin"
	rkhunter_cmd_bin := exec.Command("curl", "-L", rkhunter_bin_url, "-o", "binary.bin")
	err = rkhunter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rkhunter_src_url := "https://downloads.sourceforge.net/project/rkhunter/rkhunter/1.4.6/rkhunter-1.4.6.src.tar.gz"
	rkhunter_cmd_src := exec.Command("curl", "-L", rkhunter_src_url, "-o", "source.tar.gz")
	err = rkhunter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rkhunter_cmd_direct := exec.Command("./binary")
	err = rkhunter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
