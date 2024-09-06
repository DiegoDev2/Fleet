package main

// GithubRelease - Create and edit releases on Github (and upload artifacts)
// Homepage: https://github.com/github-release/github-release

import (
	"fmt"
	
	"os/exec"
)

func installGithubRelease() {
	// Método 1: Descargar y extraer .tar.gz
	githubrelease_tar_url := "https://github.com/github-release/github-release/archive/refs/tags/v0.10.0.tar.gz"
	githubrelease_cmd_tar := exec.Command("curl", "-L", githubrelease_tar_url, "-o", "package.tar.gz")
	err := githubrelease_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	githubrelease_zip_url := "https://github.com/github-release/github-release/archive/refs/tags/v0.10.0.zip"
	githubrelease_cmd_zip := exec.Command("curl", "-L", githubrelease_zip_url, "-o", "package.zip")
	err = githubrelease_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	githubrelease_bin_url := "https://github.com/github-release/github-release/archive/refs/tags/v0.10.0.bin"
	githubrelease_cmd_bin := exec.Command("curl", "-L", githubrelease_bin_url, "-o", "binary.bin")
	err = githubrelease_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	githubrelease_src_url := "https://github.com/github-release/github-release/archive/refs/tags/v0.10.0.src.tar.gz"
	githubrelease_cmd_src := exec.Command("curl", "-L", githubrelease_src_url, "-o", "source.tar.gz")
	err = githubrelease_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	githubrelease_cmd_direct := exec.Command("./binary")
	err = githubrelease_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
