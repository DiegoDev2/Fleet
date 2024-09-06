package main

// GitHound - Git plugin that prevents sensitive data from being committed
// Homepage: https://github.com/ezekg/git-hound

import (
	"fmt"
	
	"os/exec"
)

func installGitHound() {
	// Método 1: Descargar y extraer .tar.gz
	githound_tar_url := "https://github.com/ezekg/git-hound/archive/refs/tags/1.0.0.tar.gz"
	githound_cmd_tar := exec.Command("curl", "-L", githound_tar_url, "-o", "package.tar.gz")
	err := githound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	githound_zip_url := "https://github.com/ezekg/git-hound/archive/refs/tags/1.0.0.zip"
	githound_cmd_zip := exec.Command("curl", "-L", githound_zip_url, "-o", "package.zip")
	err = githound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	githound_bin_url := "https://github.com/ezekg/git-hound/archive/refs/tags/1.0.0.bin"
	githound_cmd_bin := exec.Command("curl", "-L", githound_bin_url, "-o", "binary.bin")
	err = githound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	githound_src_url := "https://github.com/ezekg/git-hound/archive/refs/tags/1.0.0.src.tar.gz"
	githound_cmd_src := exec.Command("curl", "-L", githound_src_url, "-o", "source.tar.gz")
	err = githound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	githound_cmd_direct := exec.Command("./binary")
	err = githound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
