package main

// GitMediate - Utility to help resolve merge conflicts
// Homepage: https://github.com/Peaker/git-mediate

import (
	"fmt"
	
	"os/exec"
)

func installGitMediate() {
	// Método 1: Descargar y extraer .tar.gz
	gitmediate_tar_url := "https://github.com/Peaker/git-mediate/archive/refs/tags/1.0.9.tar.gz"
	gitmediate_cmd_tar := exec.Command("curl", "-L", gitmediate_tar_url, "-o", "package.tar.gz")
	err := gitmediate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitmediate_zip_url := "https://github.com/Peaker/git-mediate/archive/refs/tags/1.0.9.zip"
	gitmediate_cmd_zip := exec.Command("curl", "-L", gitmediate_zip_url, "-o", "package.zip")
	err = gitmediate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitmediate_bin_url := "https://github.com/Peaker/git-mediate/archive/refs/tags/1.0.9.bin"
	gitmediate_cmd_bin := exec.Command("curl", "-L", gitmediate_bin_url, "-o", "binary.bin")
	err = gitmediate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitmediate_src_url := "https://github.com/Peaker/git-mediate/archive/refs/tags/1.0.9.src.tar.gz"
	gitmediate_cmd_src := exec.Command("curl", "-L", gitmediate_src_url, "-o", "source.tar.gz")
	err = gitmediate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitmediate_cmd_direct := exec.Command("./binary")
	err = gitmediate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
	fmt.Println("Instalando dependencia: haskell-stack")
exec.Command("latte", "install", "haskell-stack")
}
