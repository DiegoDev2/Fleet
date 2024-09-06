package main

// Oksh - Portable OpenBSD ksh, based on the public domain Korn shell (pdksh)
// Homepage: https://github.com/ibara/oksh

import (
	"fmt"
	
	"os/exec"
)

func installOksh() {
	// Método 1: Descargar y extraer .tar.gz
	oksh_tar_url := "https://github.com/ibara/oksh/releases/download/oksh-7.5/oksh-7.5.tar.gz"
	oksh_cmd_tar := exec.Command("curl", "-L", oksh_tar_url, "-o", "package.tar.gz")
	err := oksh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oksh_zip_url := "https://github.com/ibara/oksh/releases/download/oksh-7.5/oksh-7.5.zip"
	oksh_cmd_zip := exec.Command("curl", "-L", oksh_zip_url, "-o", "package.zip")
	err = oksh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oksh_bin_url := "https://github.com/ibara/oksh/releases/download/oksh-7.5/oksh-7.5.bin"
	oksh_cmd_bin := exec.Command("curl", "-L", oksh_bin_url, "-o", "binary.bin")
	err = oksh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oksh_src_url := "https://github.com/ibara/oksh/releases/download/oksh-7.5/oksh-7.5.src.tar.gz"
	oksh_cmd_src := exec.Command("curl", "-L", oksh_src_url, "-o", "source.tar.gz")
	err = oksh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oksh_cmd_direct := exec.Command("./binary")
	err = oksh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
