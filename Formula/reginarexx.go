package main

// ReginaRexx - Interpreter for Rexx
// Homepage: https://regina-rexx.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installReginaRexx() {
	// Método 1: Descargar y extraer .tar.gz
	reginarexx_tar_url := "https://downloads.sourceforge.net/project/regina-rexx/regina-rexx/3.9.6/regina-rexx-3.9.6.tar.gz"
	reginarexx_cmd_tar := exec.Command("curl", "-L", reginarexx_tar_url, "-o", "package.tar.gz")
	err := reginarexx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reginarexx_zip_url := "https://downloads.sourceforge.net/project/regina-rexx/regina-rexx/3.9.6/regina-rexx-3.9.6.zip"
	reginarexx_cmd_zip := exec.Command("curl", "-L", reginarexx_zip_url, "-o", "package.zip")
	err = reginarexx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reginarexx_bin_url := "https://downloads.sourceforge.net/project/regina-rexx/regina-rexx/3.9.6/regina-rexx-3.9.6.bin"
	reginarexx_cmd_bin := exec.Command("curl", "-L", reginarexx_bin_url, "-o", "binary.bin")
	err = reginarexx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reginarexx_src_url := "https://downloads.sourceforge.net/project/regina-rexx/regina-rexx/3.9.6/regina-rexx-3.9.6.src.tar.gz"
	reginarexx_cmd_src := exec.Command("curl", "-L", reginarexx_src_url, "-o", "source.tar.gz")
	err = reginarexx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reginarexx_cmd_direct := exec.Command("./binary")
	err = reginarexx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
