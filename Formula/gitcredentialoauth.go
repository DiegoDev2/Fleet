package main

// GitCredentialOauth - Git credential helper that authenticates in browser using OAuth
// Homepage: https://github.com/hickford/git-credential-oauth

import (
	"fmt"
	
	"os/exec"
)

func installGitCredentialOauth() {
	// Método 1: Descargar y extraer .tar.gz
	gitcredentialoauth_tar_url := "https://github.com/hickford/git-credential-oauth/archive/refs/tags/v0.13.2.tar.gz"
	gitcredentialoauth_cmd_tar := exec.Command("curl", "-L", gitcredentialoauth_tar_url, "-o", "package.tar.gz")
	err := gitcredentialoauth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcredentialoauth_zip_url := "https://github.com/hickford/git-credential-oauth/archive/refs/tags/v0.13.2.zip"
	gitcredentialoauth_cmd_zip := exec.Command("curl", "-L", gitcredentialoauth_zip_url, "-o", "package.zip")
	err = gitcredentialoauth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcredentialoauth_bin_url := "https://github.com/hickford/git-credential-oauth/archive/refs/tags/v0.13.2.bin"
	gitcredentialoauth_cmd_bin := exec.Command("curl", "-L", gitcredentialoauth_bin_url, "-o", "binary.bin")
	err = gitcredentialoauth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcredentialoauth_src_url := "https://github.com/hickford/git-credential-oauth/archive/refs/tags/v0.13.2.src.tar.gz"
	gitcredentialoauth_cmd_src := exec.Command("curl", "-L", gitcredentialoauth_src_url, "-o", "source.tar.gz")
	err = gitcredentialoauth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcredentialoauth_cmd_direct := exec.Command("./binary")
	err = gitcredentialoauth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
