package main

// Rlog - Flexible message logging facility for C++
// Homepage: https://github.com/vgough/rlog

import (
	"fmt"
	
	"os/exec"
)

func installRlog() {
	// Método 1: Descargar y extraer .tar.gz
	rlog_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/rlog/rlog-1.4.tar.gz"
	rlog_cmd_tar := exec.Command("curl", "-L", rlog_tar_url, "-o", "package.tar.gz")
	err := rlog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rlog_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/rlog/rlog-1.4.zip"
	rlog_cmd_zip := exec.Command("curl", "-L", rlog_zip_url, "-o", "package.zip")
	err = rlog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rlog_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/rlog/rlog-1.4.bin"
	rlog_cmd_bin := exec.Command("curl", "-L", rlog_bin_url, "-o", "binary.bin")
	err = rlog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rlog_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/rlog/rlog-1.4.src.tar.gz"
	rlog_cmd_src := exec.Command("curl", "-L", rlog_src_url, "-o", "source.tar.gz")
	err = rlog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rlog_cmd_direct := exec.Command("./binary")
	err = rlog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
