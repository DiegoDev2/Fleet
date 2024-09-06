package main

// Atuin - Improved shell history for zsh, bash, fish and nushell
// Homepage: https://github.com/atuinsh/atuin

import (
	"fmt"
	
	"os/exec"
)

func installAtuin() {
	// Método 1: Descargar y extraer .tar.gz
	atuin_tar_url := "https://github.com/atuinsh/atuin/archive/refs/tags/v18.3.0.tar.gz"
	atuin_cmd_tar := exec.Command("curl", "-L", atuin_tar_url, "-o", "package.tar.gz")
	err := atuin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atuin_zip_url := "https://github.com/atuinsh/atuin/archive/refs/tags/v18.3.0.zip"
	atuin_cmd_zip := exec.Command("curl", "-L", atuin_zip_url, "-o", "package.zip")
	err = atuin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atuin_bin_url := "https://github.com/atuinsh/atuin/archive/refs/tags/v18.3.0.bin"
	atuin_cmd_bin := exec.Command("curl", "-L", atuin_bin_url, "-o", "binary.bin")
	err = atuin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atuin_src_url := "https://github.com/atuinsh/atuin/archive/refs/tags/v18.3.0.src.tar.gz"
	atuin_cmd_src := exec.Command("curl", "-L", atuin_src_url, "-o", "source.tar.gz")
	err = atuin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atuin_cmd_direct := exec.Command("./binary")
	err = atuin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
