package main

// RpgCli - Your filesystem as a dungeon!
// Homepage: https://github.com/facundoolano/rpg-cli

import (
	"fmt"
	
	"os/exec"
)

func installRpgCli() {
	// Método 1: Descargar y extraer .tar.gz
	rpgcli_tar_url := "https://github.com/facundoolano/rpg-cli/archive/refs/tags/1.0.1.tar.gz"
	rpgcli_cmd_tar := exec.Command("curl", "-L", rpgcli_tar_url, "-o", "package.tar.gz")
	err := rpgcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpgcli_zip_url := "https://github.com/facundoolano/rpg-cli/archive/refs/tags/1.0.1.zip"
	rpgcli_cmd_zip := exec.Command("curl", "-L", rpgcli_zip_url, "-o", "package.zip")
	err = rpgcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpgcli_bin_url := "https://github.com/facundoolano/rpg-cli/archive/refs/tags/1.0.1.bin"
	rpgcli_cmd_bin := exec.Command("curl", "-L", rpgcli_bin_url, "-o", "binary.bin")
	err = rpgcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpgcli_src_url := "https://github.com/facundoolano/rpg-cli/archive/refs/tags/1.0.1.src.tar.gz"
	rpgcli_cmd_src := exec.Command("curl", "-L", rpgcli_src_url, "-o", "source.tar.gz")
	err = rpgcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpgcli_cmd_direct := exec.Command("./binary")
	err = rpgcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
