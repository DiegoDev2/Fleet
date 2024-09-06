package main

// GitWorkspace - Sync personal and work git repositories from multiple providers
// Homepage: https://github.com/orf/git-workspace

import (
	"fmt"
	
	"os/exec"
)

func installGitWorkspace() {
	// Método 1: Descargar y extraer .tar.gz
	gitworkspace_tar_url := "https://github.com/orf/git-workspace/archive/refs/tags/v1.5.0.tar.gz"
	gitworkspace_cmd_tar := exec.Command("curl", "-L", gitworkspace_tar_url, "-o", "package.tar.gz")
	err := gitworkspace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitworkspace_zip_url := "https://github.com/orf/git-workspace/archive/refs/tags/v1.5.0.zip"
	gitworkspace_cmd_zip := exec.Command("curl", "-L", gitworkspace_zip_url, "-o", "package.zip")
	err = gitworkspace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitworkspace_bin_url := "https://github.com/orf/git-workspace/archive/refs/tags/v1.5.0.bin"
	gitworkspace_cmd_bin := exec.Command("curl", "-L", gitworkspace_bin_url, "-o", "binary.bin")
	err = gitworkspace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitworkspace_src_url := "https://github.com/orf/git-workspace/archive/refs/tags/v1.5.0.src.tar.gz"
	gitworkspace_cmd_src := exec.Command("curl", "-L", gitworkspace_src_url, "-o", "source.tar.gz")
	err = gitworkspace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitworkspace_cmd_direct := exec.Command("./binary")
	err = gitworkspace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
