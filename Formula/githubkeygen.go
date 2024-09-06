package main

// GithubKeygen - Bootstrap GitHub SSH configuration
// Homepage: https://github.com/dolmen/github-keygen

import (
	"fmt"
	
	"os/exec"
)

func installGithubKeygen() {
	// Método 1: Descargar y extraer .tar.gz
	githubkeygen_tar_url := "https://github.com/dolmen/github-keygen/archive/refs/tags/v1.306.tar.gz"
	githubkeygen_cmd_tar := exec.Command("curl", "-L", githubkeygen_tar_url, "-o", "package.tar.gz")
	err := githubkeygen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	githubkeygen_zip_url := "https://github.com/dolmen/github-keygen/archive/refs/tags/v1.306.zip"
	githubkeygen_cmd_zip := exec.Command("curl", "-L", githubkeygen_zip_url, "-o", "package.zip")
	err = githubkeygen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	githubkeygen_bin_url := "https://github.com/dolmen/github-keygen/archive/refs/tags/v1.306.bin"
	githubkeygen_cmd_bin := exec.Command("curl", "-L", githubkeygen_bin_url, "-o", "binary.bin")
	err = githubkeygen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	githubkeygen_src_url := "https://github.com/dolmen/github-keygen/archive/refs/tags/v1.306.src.tar.gz"
	githubkeygen_cmd_src := exec.Command("curl", "-L", githubkeygen_src_url, "-o", "source.tar.gz")
	err = githubkeygen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	githubkeygen_cmd_direct := exec.Command("./binary")
	err = githubkeygen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
