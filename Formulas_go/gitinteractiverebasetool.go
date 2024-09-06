package main

// GitInteractiveRebaseTool - Native sequence editor for Git interactive rebase
// Homepage: https://gitrebasetool.mitmaro.ca/

import (
	"fmt"
	
	"os/exec"
)

func installGitInteractiveRebaseTool() {
	// Método 1: Descargar y extraer .tar.gz
	gitinteractiverebasetool_tar_url := "https://github.com/MitMaro/git-interactive-rebase-tool/archive/refs/tags/2.4.1.tar.gz"
	gitinteractiverebasetool_cmd_tar := exec.Command("curl", "-L", gitinteractiverebasetool_tar_url, "-o", "package.tar.gz")
	err := gitinteractiverebasetool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitinteractiverebasetool_zip_url := "https://github.com/MitMaro/git-interactive-rebase-tool/archive/refs/tags/2.4.1.zip"
	gitinteractiverebasetool_cmd_zip := exec.Command("curl", "-L", gitinteractiverebasetool_zip_url, "-o", "package.zip")
	err = gitinteractiverebasetool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitinteractiverebasetool_bin_url := "https://github.com/MitMaro/git-interactive-rebase-tool/archive/refs/tags/2.4.1.bin"
	gitinteractiverebasetool_cmd_bin := exec.Command("curl", "-L", gitinteractiverebasetool_bin_url, "-o", "binary.bin")
	err = gitinteractiverebasetool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitinteractiverebasetool_src_url := "https://github.com/MitMaro/git-interactive-rebase-tool/archive/refs/tags/2.4.1.src.tar.gz"
	gitinteractiverebasetool_cmd_src := exec.Command("curl", "-L", gitinteractiverebasetool_src_url, "-o", "source.tar.gz")
	err = gitinteractiverebasetool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitinteractiverebasetool_cmd_direct := exec.Command("./binary")
	err = gitinteractiverebasetool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
}
