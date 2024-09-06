package main

// GitGrab - Clone a git repository into a standard location organised by domain and path
// Homepage: https://github.com/wezm/git-grab

import (
	"fmt"
	
	"os/exec"
)

func installGitGrab() {
	// Método 1: Descargar y extraer .tar.gz
	gitgrab_tar_url := "https://github.com/wezm/git-grab/archive/refs/tags/3.0.0.tar.gz"
	gitgrab_cmd_tar := exec.Command("curl", "-L", gitgrab_tar_url, "-o", "package.tar.gz")
	err := gitgrab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitgrab_zip_url := "https://github.com/wezm/git-grab/archive/refs/tags/3.0.0.zip"
	gitgrab_cmd_zip := exec.Command("curl", "-L", gitgrab_zip_url, "-o", "package.zip")
	err = gitgrab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitgrab_bin_url := "https://github.com/wezm/git-grab/archive/refs/tags/3.0.0.bin"
	gitgrab_cmd_bin := exec.Command("curl", "-L", gitgrab_bin_url, "-o", "binary.bin")
	err = gitgrab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitgrab_src_url := "https://github.com/wezm/git-grab/archive/refs/tags/3.0.0.src.tar.gz"
	gitgrab_cmd_src := exec.Command("curl", "-L", gitgrab_src_url, "-o", "source.tar.gz")
	err = gitgrab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitgrab_cmd_direct := exec.Command("./binary")
	err = gitgrab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
