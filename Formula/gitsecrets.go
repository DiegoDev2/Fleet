package main

// GitSecrets - Prevents you from committing sensitive information to a git repo
// Homepage: https://github.com/awslabs/git-secrets

import (
	"fmt"
	
	"os/exec"
)

func installGitSecrets() {
	// Método 1: Descargar y extraer .tar.gz
	gitsecrets_tar_url := "https://github.com/awslabs/git-secrets/archive/refs/tags/1.3.0.tar.gz"
	gitsecrets_cmd_tar := exec.Command("curl", "-L", gitsecrets_tar_url, "-o", "package.tar.gz")
	err := gitsecrets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsecrets_zip_url := "https://github.com/awslabs/git-secrets/archive/refs/tags/1.3.0.zip"
	gitsecrets_cmd_zip := exec.Command("curl", "-L", gitsecrets_zip_url, "-o", "package.zip")
	err = gitsecrets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsecrets_bin_url := "https://github.com/awslabs/git-secrets/archive/refs/tags/1.3.0.bin"
	gitsecrets_cmd_bin := exec.Command("curl", "-L", gitsecrets_bin_url, "-o", "binary.bin")
	err = gitsecrets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsecrets_src_url := "https://github.com/awslabs/git-secrets/archive/refs/tags/1.3.0.src.tar.gz"
	gitsecrets_cmd_src := exec.Command("curl", "-L", gitsecrets_src_url, "-o", "source.tar.gz")
	err = gitsecrets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsecrets_cmd_direct := exec.Command("./binary")
	err = gitsecrets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
