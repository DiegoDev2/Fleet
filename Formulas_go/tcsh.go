package main

// Tcsh - Enhanced, fully compatible version of the Berkeley C shell
// Homepage: https://www.tcsh.org/

import (
	"fmt"
	
	"os/exec"
)

func installTcsh() {
	// Método 1: Descargar y extraer .tar.gz
	tcsh_tar_url := "https://astron.com/pub/tcsh/tcsh-6.24.13.tar.gz"
	tcsh_cmd_tar := exec.Command("curl", "-L", tcsh_tar_url, "-o", "package.tar.gz")
	err := tcsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcsh_zip_url := "https://astron.com/pub/tcsh/tcsh-6.24.13.zip"
	tcsh_cmd_zip := exec.Command("curl", "-L", tcsh_zip_url, "-o", "package.zip")
	err = tcsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcsh_bin_url := "https://astron.com/pub/tcsh/tcsh-6.24.13.bin"
	tcsh_cmd_bin := exec.Command("curl", "-L", tcsh_bin_url, "-o", "binary.bin")
	err = tcsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcsh_src_url := "https://astron.com/pub/tcsh/tcsh-6.24.13.src.tar.gz"
	tcsh_cmd_src := exec.Command("curl", "-L", tcsh_src_url, "-o", "source.tar.gz")
	err = tcsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcsh_cmd_direct := exec.Command("./binary")
	err = tcsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
