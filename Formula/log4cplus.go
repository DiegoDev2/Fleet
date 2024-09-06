package main

// Log4cplus - Logging Framework for C++
// Homepage: https://sourceforge.net/p/log4cplus/wiki/Home/

import (
	"fmt"
	
	"os/exec"
)

func installLog4cplus() {
	// Método 1: Descargar y extraer .tar.gz
	log4cplus_tar_url := "https://downloads.sourceforge.net/project/log4cplus/log4cplus-stable/2.1.1/log4cplus-2.1.1.tar.xz"
	log4cplus_cmd_tar := exec.Command("curl", "-L", log4cplus_tar_url, "-o", "package.tar.gz")
	err := log4cplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	log4cplus_zip_url := "https://downloads.sourceforge.net/project/log4cplus/log4cplus-stable/2.1.1/log4cplus-2.1.1.tar.xz"
	log4cplus_cmd_zip := exec.Command("curl", "-L", log4cplus_zip_url, "-o", "package.zip")
	err = log4cplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	log4cplus_bin_url := "https://downloads.sourceforge.net/project/log4cplus/log4cplus-stable/2.1.1/log4cplus-2.1.1.tar.xz"
	log4cplus_cmd_bin := exec.Command("curl", "-L", log4cplus_bin_url, "-o", "binary.bin")
	err = log4cplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	log4cplus_src_url := "https://downloads.sourceforge.net/project/log4cplus/log4cplus-stable/2.1.1/log4cplus-2.1.1.tar.xz"
	log4cplus_cmd_src := exec.Command("curl", "-L", log4cplus_src_url, "-o", "source.tar.gz")
	err = log4cplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	log4cplus_cmd_direct := exec.Command("./binary")
	err = log4cplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
