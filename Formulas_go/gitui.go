package main

// Gitui - Blazing fast terminal-ui for git written in rust
// Homepage: https://github.com/extrawurst/gitui

import (
	"fmt"
	
	"os/exec"
)

func installGitui() {
	// Método 1: Descargar y extraer .tar.gz
	gitui_tar_url := "https://github.com/extrawurst/gitui/archive/refs/tags/v0.26.3.tar.gz"
	gitui_cmd_tar := exec.Command("curl", "-L", gitui_tar_url, "-o", "package.tar.gz")
	err := gitui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitui_zip_url := "https://github.com/extrawurst/gitui/archive/refs/tags/v0.26.3.zip"
	gitui_cmd_zip := exec.Command("curl", "-L", gitui_zip_url, "-o", "package.zip")
	err = gitui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitui_bin_url := "https://github.com/extrawurst/gitui/archive/refs/tags/v0.26.3.bin"
	gitui_cmd_bin := exec.Command("curl", "-L", gitui_bin_url, "-o", "binary.bin")
	err = gitui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitui_src_url := "https://github.com/extrawurst/gitui/archive/refs/tags/v0.26.3.src.tar.gz"
	gitui_cmd_src := exec.Command("curl", "-L", gitui_src_url, "-o", "source.tar.gz")
	err = gitui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitui_cmd_direct := exec.Command("./binary")
	err = gitui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
