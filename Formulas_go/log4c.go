package main

// Log4c - Logging Framework for C
// Homepage: https://log4c.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLog4c() {
	// Método 1: Descargar y extraer .tar.gz
	log4c_tar_url := "https://downloads.sourceforge.net/project/log4c/log4c/1.2.4/log4c-1.2.4.tar.gz"
	log4c_cmd_tar := exec.Command("curl", "-L", log4c_tar_url, "-o", "package.tar.gz")
	err := log4c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	log4c_zip_url := "https://downloads.sourceforge.net/project/log4c/log4c/1.2.4/log4c-1.2.4.zip"
	log4c_cmd_zip := exec.Command("curl", "-L", log4c_zip_url, "-o", "package.zip")
	err = log4c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	log4c_bin_url := "https://downloads.sourceforge.net/project/log4c/log4c/1.2.4/log4c-1.2.4.bin"
	log4c_cmd_bin := exec.Command("curl", "-L", log4c_bin_url, "-o", "binary.bin")
	err = log4c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	log4c_src_url := "https://downloads.sourceforge.net/project/log4c/log4c/1.2.4/log4c-1.2.4.src.tar.gz"
	log4c_cmd_src := exec.Command("curl", "-L", log4c_src_url, "-o", "source.tar.gz")
	err = log4c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	log4c_cmd_direct := exec.Command("./binary")
	err = log4c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
