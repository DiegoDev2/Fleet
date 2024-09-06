package main

// GitPlus - Git utilities: git multi, git relation, git old-branches, git recent
// Homepage: https://github.com/tkrajina/git-plus

import (
	"fmt"
	
	"os/exec"
)

func installGitPlus() {
	// Método 1: Descargar y extraer .tar.gz
	gitplus_tar_url := "https://files.pythonhosted.org/packages/79/27/765447b46d6fa578892b2bdd59be3f7f3c545d68ab65ba6d3d89994ec7fc/git-plus-0.4.10.tar.gz"
	gitplus_cmd_tar := exec.Command("curl", "-L", gitplus_tar_url, "-o", "package.tar.gz")
	err := gitplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitplus_zip_url := "https://files.pythonhosted.org/packages/79/27/765447b46d6fa578892b2bdd59be3f7f3c545d68ab65ba6d3d89994ec7fc/git-plus-0.4.10.zip"
	gitplus_cmd_zip := exec.Command("curl", "-L", gitplus_zip_url, "-o", "package.zip")
	err = gitplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitplus_bin_url := "https://files.pythonhosted.org/packages/79/27/765447b46d6fa578892b2bdd59be3f7f3c545d68ab65ba6d3d89994ec7fc/git-plus-0.4.10.bin"
	gitplus_cmd_bin := exec.Command("curl", "-L", gitplus_bin_url, "-o", "binary.bin")
	err = gitplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitplus_src_url := "https://files.pythonhosted.org/packages/79/27/765447b46d6fa578892b2bdd59be3f7f3c545d68ab65ba6d3d89994ec7fc/git-plus-0.4.10.src.tar.gz"
	gitplus_cmd_src := exec.Command("curl", "-L", gitplus_src_url, "-o", "source.tar.gz")
	err = gitplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitplus_cmd_direct := exec.Command("./binary")
	err = gitplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
