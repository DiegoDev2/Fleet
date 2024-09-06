package main

// GitRevise - Rebase alternative for easy & efficient in-memory rebases and fixups
// Homepage: https://github.com/mystor/git-revise

import (
	"fmt"
	
	"os/exec"
)

func installGitRevise() {
	// Método 1: Descargar y extraer .tar.gz
	gitrevise_tar_url := "https://files.pythonhosted.org/packages/99/fe/03e0afc973c19af8ebf9c7a4a090a974c0c39578b1d4082d201d126b7f9a/git-revise-0.7.0.tar.gz"
	gitrevise_cmd_tar := exec.Command("curl", "-L", gitrevise_tar_url, "-o", "package.tar.gz")
	err := gitrevise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitrevise_zip_url := "https://files.pythonhosted.org/packages/99/fe/03e0afc973c19af8ebf9c7a4a090a974c0c39578b1d4082d201d126b7f9a/git-revise-0.7.0.zip"
	gitrevise_cmd_zip := exec.Command("curl", "-L", gitrevise_zip_url, "-o", "package.zip")
	err = gitrevise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitrevise_bin_url := "https://files.pythonhosted.org/packages/99/fe/03e0afc973c19af8ebf9c7a4a090a974c0c39578b1d4082d201d126b7f9a/git-revise-0.7.0.bin"
	gitrevise_cmd_bin := exec.Command("curl", "-L", gitrevise_bin_url, "-o", "binary.bin")
	err = gitrevise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitrevise_src_url := "https://files.pythonhosted.org/packages/99/fe/03e0afc973c19af8ebf9c7a4a090a974c0c39578b1d4082d201d126b7f9a/git-revise-0.7.0.src.tar.gz"
	gitrevise_cmd_src := exec.Command("curl", "-L", gitrevise_src_url, "-o", "source.tar.gz")
	err = gitrevise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitrevise_cmd_direct := exec.Command("./binary")
	err = gitrevise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
