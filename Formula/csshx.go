package main

// Csshx - Cluster ssh tool for Terminal.app
// Homepage: https://github.com/brockgr/csshx

import (
	"fmt"
	
	"os/exec"
)

func installCsshx() {
	// Método 1: Descargar y extraer .tar.gz
	csshx_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/csshx/csshX-0.74.tgz"
	csshx_cmd_tar := exec.Command("curl", "-L", csshx_tar_url, "-o", "package.tar.gz")
	err := csshx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csshx_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/csshx/csshX-0.74.tgz"
	csshx_cmd_zip := exec.Command("curl", "-L", csshx_zip_url, "-o", "package.zip")
	err = csshx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csshx_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/csshx/csshX-0.74.tgz"
	csshx_cmd_bin := exec.Command("curl", "-L", csshx_bin_url, "-o", "binary.bin")
	err = csshx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csshx_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/csshx/csshX-0.74.tgz"
	csshx_cmd_src := exec.Command("curl", "-L", csshx_src_url, "-o", "source.tar.gz")
	err = csshx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csshx_cmd_direct := exec.Command("./binary")
	err = csshx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
