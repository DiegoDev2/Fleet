package main

// Pure - Pretty, minimal and fast ZSH prompt
// Homepage: https://github.com/sindresorhus/pure

import (
	"fmt"
	
	"os/exec"
)

func installPure() {
	// Método 1: Descargar y extraer .tar.gz
	pure_tar_url := "https://github.com/sindresorhus/pure/archive/refs/tags/v1.23.0.tar.gz"
	pure_cmd_tar := exec.Command("curl", "-L", pure_tar_url, "-o", "package.tar.gz")
	err := pure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pure_zip_url := "https://github.com/sindresorhus/pure/archive/refs/tags/v1.23.0.zip"
	pure_cmd_zip := exec.Command("curl", "-L", pure_zip_url, "-o", "package.zip")
	err = pure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pure_bin_url := "https://github.com/sindresorhus/pure/archive/refs/tags/v1.23.0.bin"
	pure_cmd_bin := exec.Command("curl", "-L", pure_bin_url, "-o", "binary.bin")
	err = pure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pure_src_url := "https://github.com/sindresorhus/pure/archive/refs/tags/v1.23.0.src.tar.gz"
	pure_cmd_src := exec.Command("curl", "-L", pure_src_url, "-o", "source.tar.gz")
	err = pure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pure_cmd_direct := exec.Command("./binary")
	err = pure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zsh")
exec.Command("latte", "install", "zsh")
	fmt.Println("Instalando dependencia: zsh-async")
exec.Command("latte", "install", "zsh-async")
}
