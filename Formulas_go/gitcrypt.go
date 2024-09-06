package main

// GitCrypt - Enable transparent encryption/decryption of files in a git repo
// Homepage: https://www.agwa.name/projects/git-crypt/

import (
	"fmt"
	
	"os/exec"
)

func installGitCrypt() {
	// Método 1: Descargar y extraer .tar.gz
	gitcrypt_tar_url := "https://www.agwa.name/projects/git-crypt/downloads/git-crypt-0.7.0.tar.gz"
	gitcrypt_cmd_tar := exec.Command("curl", "-L", gitcrypt_tar_url, "-o", "package.tar.gz")
	err := gitcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcrypt_zip_url := "https://www.agwa.name/projects/git-crypt/downloads/git-crypt-0.7.0.zip"
	gitcrypt_cmd_zip := exec.Command("curl", "-L", gitcrypt_zip_url, "-o", "package.zip")
	err = gitcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcrypt_bin_url := "https://www.agwa.name/projects/git-crypt/downloads/git-crypt-0.7.0.bin"
	gitcrypt_cmd_bin := exec.Command("curl", "-L", gitcrypt_bin_url, "-o", "binary.bin")
	err = gitcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcrypt_src_url := "https://www.agwa.name/projects/git-crypt/downloads/git-crypt-0.7.0.src.tar.gz"
	gitcrypt_cmd_src := exec.Command("curl", "-L", gitcrypt_src_url, "-o", "source.tar.gz")
	err = gitcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcrypt_cmd_direct := exec.Command("./binary")
	err = gitcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
