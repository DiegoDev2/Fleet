package main

// EmacsDracula - Dark color theme available for a number of editors
// Homepage: https://github.com/dracula/emacs

import (
	"fmt"
	
	"os/exec"
)

func installEmacsDracula() {
	// Método 1: Descargar y extraer .tar.gz
	emacsdracula_tar_url := "https://github.com/dracula/emacs/archive/refs/tags/v1.8.2.tar.gz"
	emacsdracula_cmd_tar := exec.Command("curl", "-L", emacsdracula_tar_url, "-o", "package.tar.gz")
	err := emacsdracula_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emacsdracula_zip_url := "https://github.com/dracula/emacs/archive/refs/tags/v1.8.2.zip"
	emacsdracula_cmd_zip := exec.Command("curl", "-L", emacsdracula_zip_url, "-o", "package.zip")
	err = emacsdracula_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emacsdracula_bin_url := "https://github.com/dracula/emacs/archive/refs/tags/v1.8.2.bin"
	emacsdracula_cmd_bin := exec.Command("curl", "-L", emacsdracula_bin_url, "-o", "binary.bin")
	err = emacsdracula_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emacsdracula_src_url := "https://github.com/dracula/emacs/archive/refs/tags/v1.8.2.src.tar.gz"
	emacsdracula_cmd_src := exec.Command("curl", "-L", emacsdracula_src_url, "-o", "source.tar.gz")
	err = emacsdracula_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emacsdracula_cmd_direct := exec.Command("./binary")
	err = emacsdracula_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
}
