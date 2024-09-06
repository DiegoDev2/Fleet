package main

// TrzszSsh - Simple ssh client with trzsz ( trz / tsz ) support
// Homepage: https://trzsz.github.io/ssh

import (
	"fmt"
	
	"os/exec"
)

func installTrzszSsh() {
	// Método 1: Descargar y extraer .tar.gz
	trzszssh_tar_url := "https://github.com/trzsz/trzsz-ssh/archive/refs/tags/v0.1.22.tar.gz"
	trzszssh_cmd_tar := exec.Command("curl", "-L", trzszssh_tar_url, "-o", "package.tar.gz")
	err := trzszssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trzszssh_zip_url := "https://github.com/trzsz/trzsz-ssh/archive/refs/tags/v0.1.22.zip"
	trzszssh_cmd_zip := exec.Command("curl", "-L", trzszssh_zip_url, "-o", "package.zip")
	err = trzszssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trzszssh_bin_url := "https://github.com/trzsz/trzsz-ssh/archive/refs/tags/v0.1.22.bin"
	trzszssh_cmd_bin := exec.Command("curl", "-L", trzszssh_bin_url, "-o", "binary.bin")
	err = trzszssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trzszssh_src_url := "https://github.com/trzsz/trzsz-ssh/archive/refs/tags/v0.1.22.src.tar.gz"
	trzszssh_cmd_src := exec.Command("curl", "-L", trzszssh_src_url, "-o", "source.tar.gz")
	err = trzszssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trzszssh_cmd_direct := exec.Command("./binary")
	err = trzszssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
