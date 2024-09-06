package main

// Rinetd - Internet TCP redirection server
// Homepage: https://github.com/samhocevar/rinetd

import (
	"fmt"
	
	"os/exec"
)

func installRinetd() {
	// Método 1: Descargar y extraer .tar.gz
	rinetd_tar_url := "https://github.com/samhocevar/rinetd/releases/download/v0.73/rinetd-0.73.tar.bz2"
	rinetd_cmd_tar := exec.Command("curl", "-L", rinetd_tar_url, "-o", "package.tar.gz")
	err := rinetd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rinetd_zip_url := "https://github.com/samhocevar/rinetd/releases/download/v0.73/rinetd-0.73.tar.bz2"
	rinetd_cmd_zip := exec.Command("curl", "-L", rinetd_zip_url, "-o", "package.zip")
	err = rinetd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rinetd_bin_url := "https://github.com/samhocevar/rinetd/releases/download/v0.73/rinetd-0.73.tar.bz2"
	rinetd_cmd_bin := exec.Command("curl", "-L", rinetd_bin_url, "-o", "binary.bin")
	err = rinetd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rinetd_src_url := "https://github.com/samhocevar/rinetd/releases/download/v0.73/rinetd-0.73.tar.bz2"
	rinetd_cmd_src := exec.Command("curl", "-L", rinetd_src_url, "-o", "source.tar.gz")
	err = rinetd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rinetd_cmd_direct := exec.Command("./binary")
	err = rinetd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
