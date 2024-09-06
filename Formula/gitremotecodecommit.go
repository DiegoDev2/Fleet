package main

// GitRemoteCodecommit - Git Remote Helper to interact with AWS CodeCommit
// Homepage: https://github.com/aws/git-remote-codecommit

import (
	"fmt"
	
	"os/exec"
)

func installGitRemoteCodecommit() {
	// Método 1: Descargar y extraer .tar.gz
	gitremotecodecommit_tar_url := "https://files.pythonhosted.org/packages/6c/a0/feb4dfa42e8cb1a0bd91667233254e49696cf6618f51ad5629f6efd89dae/git-remote-codecommit-1.17.tar.gz"
	gitremotecodecommit_cmd_tar := exec.Command("curl", "-L", gitremotecodecommit_tar_url, "-o", "package.tar.gz")
	err := gitremotecodecommit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitremotecodecommit_zip_url := "https://files.pythonhosted.org/packages/6c/a0/feb4dfa42e8cb1a0bd91667233254e49696cf6618f51ad5629f6efd89dae/git-remote-codecommit-1.17.zip"
	gitremotecodecommit_cmd_zip := exec.Command("curl", "-L", gitremotecodecommit_zip_url, "-o", "package.zip")
	err = gitremotecodecommit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitremotecodecommit_bin_url := "https://files.pythonhosted.org/packages/6c/a0/feb4dfa42e8cb1a0bd91667233254e49696cf6618f51ad5629f6efd89dae/git-remote-codecommit-1.17.bin"
	gitremotecodecommit_cmd_bin := exec.Command("curl", "-L", gitremotecodecommit_bin_url, "-o", "binary.bin")
	err = gitremotecodecommit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitremotecodecommit_src_url := "https://files.pythonhosted.org/packages/6c/a0/feb4dfa42e8cb1a0bd91667233254e49696cf6618f51ad5629f6efd89dae/git-remote-codecommit-1.17.src.tar.gz"
	gitremotecodecommit_cmd_src := exec.Command("curl", "-L", gitremotecodecommit_src_url, "-o", "source.tar.gz")
	err = gitremotecodecommit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitremotecodecommit_cmd_direct := exec.Command("./binary")
	err = gitremotecodecommit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
