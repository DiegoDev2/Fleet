package main

// Gitmoji - Interactive command-line tool for using emoji in commit messages
// Homepage: https://gitmoji.dev

import (
	"fmt"
	
	"os/exec"
)

func installGitmoji() {
	// Método 1: Descargar y extraer .tar.gz
	gitmoji_tar_url := "https://registry.npmjs.org/gitmoji-cli/-/gitmoji-cli-9.4.0.tgz"
	gitmoji_cmd_tar := exec.Command("curl", "-L", gitmoji_tar_url, "-o", "package.tar.gz")
	err := gitmoji_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitmoji_zip_url := "https://registry.npmjs.org/gitmoji-cli/-/gitmoji-cli-9.4.0.tgz"
	gitmoji_cmd_zip := exec.Command("curl", "-L", gitmoji_zip_url, "-o", "package.zip")
	err = gitmoji_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitmoji_bin_url := "https://registry.npmjs.org/gitmoji-cli/-/gitmoji-cli-9.4.0.tgz"
	gitmoji_cmd_bin := exec.Command("curl", "-L", gitmoji_bin_url, "-o", "binary.bin")
	err = gitmoji_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitmoji_src_url := "https://registry.npmjs.org/gitmoji-cli/-/gitmoji-cli-9.4.0.tgz"
	gitmoji_cmd_src := exec.Command("curl", "-L", gitmoji_src_url, "-o", "source.tar.gz")
	err = gitmoji_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitmoji_cmd_direct := exec.Command("./binary")
	err = gitmoji_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
