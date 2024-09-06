package main

// Sdedit - Tool for generating sequence diagrams very quickly
// Homepage: https://sdedit.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSdedit() {
	// Método 1: Descargar y extraer .tar.gz
	sdedit_tar_url := "https://downloads.sourceforge.net/project/sdedit/sdedit/4.2/sdedit-4.2.1.jar"
	sdedit_cmd_tar := exec.Command("curl", "-L", sdedit_tar_url, "-o", "package.tar.gz")
	err := sdedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdedit_zip_url := "https://downloads.sourceforge.net/project/sdedit/sdedit/4.2/sdedit-4.2.1.jar"
	sdedit_cmd_zip := exec.Command("curl", "-L", sdedit_zip_url, "-o", "package.zip")
	err = sdedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdedit_bin_url := "https://downloads.sourceforge.net/project/sdedit/sdedit/4.2/sdedit-4.2.1.jar"
	sdedit_cmd_bin := exec.Command("curl", "-L", sdedit_bin_url, "-o", "binary.bin")
	err = sdedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdedit_src_url := "https://downloads.sourceforge.net/project/sdedit/sdedit/4.2/sdedit-4.2.1.jar"
	sdedit_cmd_src := exec.Command("curl", "-L", sdedit_src_url, "-o", "source.tar.gz")
	err = sdedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdedit_cmd_direct := exec.Command("./binary")
	err = sdedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
