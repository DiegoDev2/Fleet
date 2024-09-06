package main

// GitSecret - Bash-tool to store the private data inside a git repo
// Homepage: https://sobolevn.me/git-secret

import (
	"fmt"
	
	"os/exec"
)

func installGitSecret() {
	// Método 1: Descargar y extraer .tar.gz
	gitsecret_tar_url := "https://github.com/sobolevn/git-secret/archive/refs/tags/v0.5.0.tar.gz"
	gitsecret_cmd_tar := exec.Command("curl", "-L", gitsecret_tar_url, "-o", "package.tar.gz")
	err := gitsecret_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsecret_zip_url := "https://github.com/sobolevn/git-secret/archive/refs/tags/v0.5.0.zip"
	gitsecret_cmd_zip := exec.Command("curl", "-L", gitsecret_zip_url, "-o", "package.zip")
	err = gitsecret_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsecret_bin_url := "https://github.com/sobolevn/git-secret/archive/refs/tags/v0.5.0.bin"
	gitsecret_cmd_bin := exec.Command("curl", "-L", gitsecret_bin_url, "-o", "binary.bin")
	err = gitsecret_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsecret_src_url := "https://github.com/sobolevn/git-secret/archive/refs/tags/v0.5.0.src.tar.gz"
	gitsecret_cmd_src := exec.Command("curl", "-L", gitsecret_src_url, "-o", "source.tar.gz")
	err = gitsecret_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsecret_cmd_direct := exec.Command("./binary")
	err = gitsecret_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
}
