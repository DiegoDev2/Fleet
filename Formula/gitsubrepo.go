package main

// GitSubrepo - Git Submodule Alternative
// Homepage: https://github.com/ingydotnet/git-subrepo

import (
	"fmt"
	
	"os/exec"
)

func installGitSubrepo() {
	// Método 1: Descargar y extraer .tar.gz
	gitsubrepo_tar_url := "https://github.com/ingydotnet/git-subrepo/archive/refs/tags/0.4.9.tar.gz"
	gitsubrepo_cmd_tar := exec.Command("curl", "-L", gitsubrepo_tar_url, "-o", "package.tar.gz")
	err := gitsubrepo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsubrepo_zip_url := "https://github.com/ingydotnet/git-subrepo/archive/refs/tags/0.4.9.zip"
	gitsubrepo_cmd_zip := exec.Command("curl", "-L", gitsubrepo_zip_url, "-o", "package.zip")
	err = gitsubrepo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsubrepo_bin_url := "https://github.com/ingydotnet/git-subrepo/archive/refs/tags/0.4.9.bin"
	gitsubrepo_cmd_bin := exec.Command("curl", "-L", gitsubrepo_bin_url, "-o", "binary.bin")
	err = gitsubrepo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsubrepo_src_url := "https://github.com/ingydotnet/git-subrepo/archive/refs/tags/0.4.9.src.tar.gz"
	gitsubrepo_cmd_src := exec.Command("curl", "-L", gitsubrepo_src_url, "-o", "source.tar.gz")
	err = gitsubrepo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsubrepo_cmd_direct := exec.Command("./binary")
	err = gitsubrepo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
}
