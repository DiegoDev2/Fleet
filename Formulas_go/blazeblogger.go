package main

// Blazeblogger - CMS for the command-line
// Homepage: http://blaze.blackened.cz/

import (
	"fmt"
	
	"os/exec"
)

func installBlazeblogger() {
	// Método 1: Descargar y extraer .tar.gz
	blazeblogger_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/blazeblogger/blazeblogger-1.2.0.tar.gz"
	blazeblogger_cmd_tar := exec.Command("curl", "-L", blazeblogger_tar_url, "-o", "package.tar.gz")
	err := blazeblogger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blazeblogger_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/blazeblogger/blazeblogger-1.2.0.zip"
	blazeblogger_cmd_zip := exec.Command("curl", "-L", blazeblogger_zip_url, "-o", "package.zip")
	err = blazeblogger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blazeblogger_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/blazeblogger/blazeblogger-1.2.0.bin"
	blazeblogger_cmd_bin := exec.Command("curl", "-L", blazeblogger_bin_url, "-o", "binary.bin")
	err = blazeblogger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blazeblogger_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/blazeblogger/blazeblogger-1.2.0.src.tar.gz"
	blazeblogger_cmd_src := exec.Command("curl", "-L", blazeblogger_src_url, "-o", "source.tar.gz")
	err = blazeblogger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blazeblogger_cmd_direct := exec.Command("./binary")
	err = blazeblogger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
