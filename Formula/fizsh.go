package main

// Fizsh - Fish-like front end for ZSH
// Homepage: https://github.com/zsh-users/fizsh

import (
	"fmt"
	
	"os/exec"
)

func installFizsh() {
	// Método 1: Descargar y extraer .tar.gz
	fizsh_tar_url := "https://downloads.sourceforge.net/project/fizsh/fizsh-1.0.9.tar.gz"
	fizsh_cmd_tar := exec.Command("curl", "-L", fizsh_tar_url, "-o", "package.tar.gz")
	err := fizsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fizsh_zip_url := "https://downloads.sourceforge.net/project/fizsh/fizsh-1.0.9.zip"
	fizsh_cmd_zip := exec.Command("curl", "-L", fizsh_zip_url, "-o", "package.zip")
	err = fizsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fizsh_bin_url := "https://downloads.sourceforge.net/project/fizsh/fizsh-1.0.9.bin"
	fizsh_cmd_bin := exec.Command("curl", "-L", fizsh_bin_url, "-o", "binary.bin")
	err = fizsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fizsh_src_url := "https://downloads.sourceforge.net/project/fizsh/fizsh-1.0.9.src.tar.gz"
	fizsh_cmd_src := exec.Command("curl", "-L", fizsh_src_url, "-o", "source.tar.gz")
	err = fizsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fizsh_cmd_direct := exec.Command("./binary")
	err = fizsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zsh")
	exec.Command("latte", "install", "zsh").Run()
}
