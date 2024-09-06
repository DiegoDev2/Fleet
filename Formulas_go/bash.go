package main

// Bash - Bourne-Again SHell, a UNIX command interpreter
// Homepage: https://www.gnu.org/software/bash/

import (
	"fmt"
	
	"os/exec"
)

func installBash() {
	// Método 1: Descargar y extraer .tar.gz
	bash_tar_url := "https://ftp.gnu.org/gnu/bash/bash-5.2.tar.gz"
	bash_cmd_tar := exec.Command("curl", "-L", bash_tar_url, "-o", "package.tar.gz")
	err := bash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bash_zip_url := "https://ftp.gnu.org/gnu/bash/bash-5.2.zip"
	bash_cmd_zip := exec.Command("curl", "-L", bash_zip_url, "-o", "package.zip")
	err = bash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bash_bin_url := "https://ftp.gnu.org/gnu/bash/bash-5.2.bin"
	bash_cmd_bin := exec.Command("curl", "-L", bash_bin_url, "-o", "binary.bin")
	err = bash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bash_src_url := "https://ftp.gnu.org/gnu/bash/bash-5.2.src.tar.gz"
	bash_cmd_src := exec.Command("curl", "-L", bash_src_url, "-o", "source.tar.gz")
	err = bash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bash_cmd_direct := exec.Command("./binary")
	err = bash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
