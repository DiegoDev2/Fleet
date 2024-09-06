package main

// Cheat - Create and view interactive cheat sheets for *nix commands
// Homepage: https://github.com/cheat/cheat

import (
	"fmt"
	
	"os/exec"
)

func installCheat() {
	// Método 1: Descargar y extraer .tar.gz
	cheat_tar_url := "https://github.com/cheat/cheat/archive/refs/tags/4.4.2.tar.gz"
	cheat_cmd_tar := exec.Command("curl", "-L", cheat_tar_url, "-o", "package.tar.gz")
	err := cheat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cheat_zip_url := "https://github.com/cheat/cheat/archive/refs/tags/4.4.2.zip"
	cheat_cmd_zip := exec.Command("curl", "-L", cheat_zip_url, "-o", "package.zip")
	err = cheat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cheat_bin_url := "https://github.com/cheat/cheat/archive/refs/tags/4.4.2.bin"
	cheat_cmd_bin := exec.Command("curl", "-L", cheat_bin_url, "-o", "binary.bin")
	err = cheat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cheat_src_url := "https://github.com/cheat/cheat/archive/refs/tags/4.4.2.src.tar.gz"
	cheat_cmd_src := exec.Command("curl", "-L", cheat_src_url, "-o", "source.tar.gz")
	err = cheat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cheat_cmd_direct := exec.Command("./binary")
	err = cheat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
