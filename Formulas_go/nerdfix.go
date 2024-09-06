package main

// Nerdfix - Find/fix obsolete Nerd Font icons
// Homepage: https://github.com/loichyan/nerdfix

import (
	"fmt"
	
	"os/exec"
)

func installNerdfix() {
	// Método 1: Descargar y extraer .tar.gz
	nerdfix_tar_url := "https://github.com/loichyan/nerdfix/archive/refs/tags/v0.4.1.tar.gz"
	nerdfix_cmd_tar := exec.Command("curl", "-L", nerdfix_tar_url, "-o", "package.tar.gz")
	err := nerdfix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nerdfix_zip_url := "https://github.com/loichyan/nerdfix/archive/refs/tags/v0.4.1.zip"
	nerdfix_cmd_zip := exec.Command("curl", "-L", nerdfix_zip_url, "-o", "package.zip")
	err = nerdfix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nerdfix_bin_url := "https://github.com/loichyan/nerdfix/archive/refs/tags/v0.4.1.bin"
	nerdfix_cmd_bin := exec.Command("curl", "-L", nerdfix_bin_url, "-o", "binary.bin")
	err = nerdfix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nerdfix_src_url := "https://github.com/loichyan/nerdfix/archive/refs/tags/v0.4.1.src.tar.gz"
	nerdfix_cmd_src := exec.Command("curl", "-L", nerdfix_src_url, "-o", "source.tar.gz")
	err = nerdfix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nerdfix_cmd_direct := exec.Command("./binary")
	err = nerdfix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
