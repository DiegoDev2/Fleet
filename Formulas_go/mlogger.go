package main

// Mlogger - Log to syslog from the command-line
// Homepage: https://github.com/nbrownus/mlogger

import (
	"fmt"
	
	"os/exec"
)

func installMlogger() {
	// Método 1: Descargar y extraer .tar.gz
	mlogger_tar_url := "https://github.com/nbrownus/mlogger/archive/refs/tags/v1.2.0.tar.gz"
	mlogger_cmd_tar := exec.Command("curl", "-L", mlogger_tar_url, "-o", "package.tar.gz")
	err := mlogger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlogger_zip_url := "https://github.com/nbrownus/mlogger/archive/refs/tags/v1.2.0.zip"
	mlogger_cmd_zip := exec.Command("curl", "-L", mlogger_zip_url, "-o", "package.zip")
	err = mlogger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlogger_bin_url := "https://github.com/nbrownus/mlogger/archive/refs/tags/v1.2.0.bin"
	mlogger_cmd_bin := exec.Command("curl", "-L", mlogger_bin_url, "-o", "binary.bin")
	err = mlogger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlogger_src_url := "https://github.com/nbrownus/mlogger/archive/refs/tags/v1.2.0.src.tar.gz"
	mlogger_cmd_src := exec.Command("curl", "-L", mlogger_src_url, "-o", "source.tar.gz")
	err = mlogger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlogger_cmd_direct := exec.Command("./binary")
	err = mlogger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
