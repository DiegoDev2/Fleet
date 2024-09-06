package main

// BashCompletion - Programmable completion for Bash 3.2
// Homepage: https://salsa.debian.org/debian/bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installBashCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	bashcompletion_tar_url := "https://src.fedoraproject.org/repo/pkgs/bash-completion/bash-completion-1.3.tar.bz2/a1262659b4bbf44dc9e59d034de505ec/bash-completion-1.3.tar.bz2"
	bashcompletion_cmd_tar := exec.Command("curl", "-L", bashcompletion_tar_url, "-o", "package.tar.gz")
	err := bashcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashcompletion_zip_url := "https://src.fedoraproject.org/repo/pkgs/bash-completion/bash-completion-1.3.tar.bz2/a1262659b4bbf44dc9e59d034de505ec/bash-completion-1.3.tar.bz2"
	bashcompletion_cmd_zip := exec.Command("curl", "-L", bashcompletion_zip_url, "-o", "package.zip")
	err = bashcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashcompletion_bin_url := "https://src.fedoraproject.org/repo/pkgs/bash-completion/bash-completion-1.3.tar.bz2/a1262659b4bbf44dc9e59d034de505ec/bash-completion-1.3.tar.bz2"
	bashcompletion_cmd_bin := exec.Command("curl", "-L", bashcompletion_bin_url, "-o", "binary.bin")
	err = bashcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashcompletion_src_url := "https://src.fedoraproject.org/repo/pkgs/bash-completion/bash-completion-1.3.tar.bz2/a1262659b4bbf44dc9e59d034de505ec/bash-completion-1.3.tar.bz2"
	bashcompletion_cmd_src := exec.Command("curl", "-L", bashcompletion_src_url, "-o", "source.tar.gz")
	err = bashcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashcompletion_cmd_direct := exec.Command("./binary")
	err = bashcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
