package main

// Gitwatch - Watch a file or folder and automatically commit changes to a git repo easily
// Homepage: https://github.com/gitwatch/gitwatch

import (
	"fmt"
	
	"os/exec"
)

func installGitwatch() {
	// Método 1: Descargar y extraer .tar.gz
	gitwatch_tar_url := "https://github.com/gitwatch/gitwatch/archive/refs/tags/v0.2.tar.gz"
	gitwatch_cmd_tar := exec.Command("curl", "-L", gitwatch_tar_url, "-o", "package.tar.gz")
	err := gitwatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitwatch_zip_url := "https://github.com/gitwatch/gitwatch/archive/refs/tags/v0.2.zip"
	gitwatch_cmd_zip := exec.Command("curl", "-L", gitwatch_zip_url, "-o", "package.zip")
	err = gitwatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitwatch_bin_url := "https://github.com/gitwatch/gitwatch/archive/refs/tags/v0.2.bin"
	gitwatch_cmd_bin := exec.Command("curl", "-L", gitwatch_bin_url, "-o", "binary.bin")
	err = gitwatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitwatch_src_url := "https://github.com/gitwatch/gitwatch/archive/refs/tags/v0.2.src.tar.gz"
	gitwatch_cmd_src := exec.Command("curl", "-L", gitwatch_src_url, "-o", "source.tar.gz")
	err = gitwatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitwatch_cmd_direct := exec.Command("./binary")
	err = gitwatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: fswatch")
exec.Command("latte", "install", "fswatch")
	fmt.Println("Instalando dependencia: inotify-tools")
exec.Command("latte", "install", "inotify-tools")
}
