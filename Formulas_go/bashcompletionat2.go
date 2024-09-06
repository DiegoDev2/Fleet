package main

// BashCompletionAT2 - Programmable completion for Bash 4.2+
// Homepage: https://github.com/scop/bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installBashCompletionAT2() {
	// Método 1: Descargar y extraer .tar.gz
	bashcompletionat2_tar_url := "https://github.com/scop/bash-completion/releases/download/2.14.0/bash-completion-2.14.0.tar.xz"
	bashcompletionat2_cmd_tar := exec.Command("curl", "-L", bashcompletionat2_tar_url, "-o", "package.tar.gz")
	err := bashcompletionat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashcompletionat2_zip_url := "https://github.com/scop/bash-completion/releases/download/2.14.0/bash-completion-2.14.0.tar.xz"
	bashcompletionat2_cmd_zip := exec.Command("curl", "-L", bashcompletionat2_zip_url, "-o", "package.zip")
	err = bashcompletionat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashcompletionat2_bin_url := "https://github.com/scop/bash-completion/releases/download/2.14.0/bash-completion-2.14.0.tar.xz"
	bashcompletionat2_cmd_bin := exec.Command("curl", "-L", bashcompletionat2_bin_url, "-o", "binary.bin")
	err = bashcompletionat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashcompletionat2_src_url := "https://github.com/scop/bash-completion/releases/download/2.14.0/bash-completion-2.14.0.tar.xz"
	bashcompletionat2_cmd_src := exec.Command("curl", "-L", bashcompletionat2_src_url, "-o", "source.tar.gz")
	err = bashcompletionat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashcompletionat2_cmd_direct := exec.Command("./binary")
	err = bashcompletionat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
}
