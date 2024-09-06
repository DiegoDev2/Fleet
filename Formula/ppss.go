package main

// Ppss - Shell script to execute commands in parallel
// Homepage: https://github.com/louwrentius/PPSS

import (
	"fmt"
	
	"os/exec"
)

func installPpss() {
	// Método 1: Descargar y extraer .tar.gz
	ppss_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ppss/ppss-2.97.tgz"
	ppss_cmd_tar := exec.Command("curl", "-L", ppss_tar_url, "-o", "package.tar.gz")
	err := ppss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ppss_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ppss/ppss-2.97.tgz"
	ppss_cmd_zip := exec.Command("curl", "-L", ppss_zip_url, "-o", "package.zip")
	err = ppss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ppss_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ppss/ppss-2.97.tgz"
	ppss_cmd_bin := exec.Command("curl", "-L", ppss_bin_url, "-o", "binary.bin")
	err = ppss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ppss_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ppss/ppss-2.97.tgz"
	ppss_cmd_src := exec.Command("curl", "-L", ppss_src_url, "-o", "source.tar.gz")
	err = ppss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ppss_cmd_direct := exec.Command("./binary")
	err = ppss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
