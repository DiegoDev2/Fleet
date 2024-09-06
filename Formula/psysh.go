package main

// Psysh - Runtime developer console, interactive debugger and REPL for PHP
// Homepage: https://psysh.org/

import (
	"fmt"
	
	"os/exec"
)

func installPsysh() {
	// Método 1: Descargar y extraer .tar.gz
	psysh_tar_url := "https://github.com/bobthecow/psysh/releases/download/v0.12.4/psysh-v0.12.4.tar.gz"
	psysh_cmd_tar := exec.Command("curl", "-L", psysh_tar_url, "-o", "package.tar.gz")
	err := psysh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psysh_zip_url := "https://github.com/bobthecow/psysh/releases/download/v0.12.4/psysh-v0.12.4.zip"
	psysh_cmd_zip := exec.Command("curl", "-L", psysh_zip_url, "-o", "package.zip")
	err = psysh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psysh_bin_url := "https://github.com/bobthecow/psysh/releases/download/v0.12.4/psysh-v0.12.4.bin"
	psysh_cmd_bin := exec.Command("curl", "-L", psysh_bin_url, "-o", "binary.bin")
	err = psysh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psysh_src_url := "https://github.com/bobthecow/psysh/releases/download/v0.12.4/psysh-v0.12.4.src.tar.gz"
	psysh_cmd_src := exec.Command("curl", "-L", psysh_src_url, "-o", "source.tar.gz")
	err = psysh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psysh_cmd_direct := exec.Command("./binary")
	err = psysh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
