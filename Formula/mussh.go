package main

// Mussh - Multi-host SSH wrapper
// Homepage: https://mussh.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMussh() {
	// Método 1: Descargar y extraer .tar.gz
	mussh_tar_url := "https://downloads.sourceforge.net/project/mussh/mussh/1.0/mussh-1.0.tgz"
	mussh_cmd_tar := exec.Command("curl", "-L", mussh_tar_url, "-o", "package.tar.gz")
	err := mussh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mussh_zip_url := "https://downloads.sourceforge.net/project/mussh/mussh/1.0/mussh-1.0.tgz"
	mussh_cmd_zip := exec.Command("curl", "-L", mussh_zip_url, "-o", "package.zip")
	err = mussh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mussh_bin_url := "https://downloads.sourceforge.net/project/mussh/mussh/1.0/mussh-1.0.tgz"
	mussh_cmd_bin := exec.Command("curl", "-L", mussh_bin_url, "-o", "binary.bin")
	err = mussh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mussh_src_url := "https://downloads.sourceforge.net/project/mussh/mussh/1.0/mussh-1.0.tgz"
	mussh_cmd_src := exec.Command("curl", "-L", mussh_src_url, "-o", "source.tar.gz")
	err = mussh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mussh_cmd_direct := exec.Command("./binary")
	err = mussh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
