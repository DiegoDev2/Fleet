package main

// Ragel - State machine compiler
// Homepage: https://www.colm.net/open-source/ragel/

import (
	"fmt"
	
	"os/exec"
)

func installRagel() {
	// Método 1: Descargar y extraer .tar.gz
	ragel_tar_url := "https://www.colm.net/files/ragel/ragel-6.10.tar.gz"
	ragel_cmd_tar := exec.Command("curl", "-L", ragel_tar_url, "-o", "package.tar.gz")
	err := ragel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ragel_zip_url := "https://www.colm.net/files/ragel/ragel-6.10.zip"
	ragel_cmd_zip := exec.Command("curl", "-L", ragel_zip_url, "-o", "package.zip")
	err = ragel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ragel_bin_url := "https://www.colm.net/files/ragel/ragel-6.10.bin"
	ragel_cmd_bin := exec.Command("curl", "-L", ragel_bin_url, "-o", "binary.bin")
	err = ragel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ragel_src_url := "https://www.colm.net/files/ragel/ragel-6.10.src.tar.gz"
	ragel_cmd_src := exec.Command("curl", "-L", ragel_src_url, "-o", "source.tar.gz")
	err = ragel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ragel_cmd_direct := exec.Command("./binary")
	err = ragel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
