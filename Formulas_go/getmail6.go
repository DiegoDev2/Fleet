package main

// Getmail6 - Extensible mail retrieval system with POP3, IMAP4, SSL support
// Homepage: https://getmail6.org/

import (
	"fmt"
	
	"os/exec"
)

func installGetmail6() {
	// Método 1: Descargar y extraer .tar.gz
	getmail6_tar_url := "https://github.com/getmail6/getmail6/archive/refs/tags/v6.19.04.tar.gz"
	getmail6_cmd_tar := exec.Command("curl", "-L", getmail6_tar_url, "-o", "package.tar.gz")
	err := getmail6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	getmail6_zip_url := "https://github.com/getmail6/getmail6/archive/refs/tags/v6.19.04.zip"
	getmail6_cmd_zip := exec.Command("curl", "-L", getmail6_zip_url, "-o", "package.zip")
	err = getmail6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	getmail6_bin_url := "https://github.com/getmail6/getmail6/archive/refs/tags/v6.19.04.bin"
	getmail6_cmd_bin := exec.Command("curl", "-L", getmail6_bin_url, "-o", "binary.bin")
	err = getmail6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	getmail6_src_url := "https://github.com/getmail6/getmail6/archive/refs/tags/v6.19.04.src.tar.gz"
	getmail6_cmd_src := exec.Command("curl", "-L", getmail6_src_url, "-o", "source.tar.gz")
	err = getmail6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	getmail6_cmd_direct := exec.Command("./binary")
	err = getmail6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
