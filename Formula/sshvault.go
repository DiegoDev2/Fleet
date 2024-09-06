package main

// SshVault - Encrypt/decrypt using SSH keys
// Homepage: https://ssh-vault.com/

import (
	"fmt"
	
	"os/exec"
)

func installSshVault() {
	// Método 1: Descargar y extraer .tar.gz
	sshvault_tar_url := "https://github.com/ssh-vault/ssh-vault/archive/refs/tags/1.0.13.tar.gz"
	sshvault_cmd_tar := exec.Command("curl", "-L", sshvault_tar_url, "-o", "package.tar.gz")
	err := sshvault_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshvault_zip_url := "https://github.com/ssh-vault/ssh-vault/archive/refs/tags/1.0.13.zip"
	sshvault_cmd_zip := exec.Command("curl", "-L", sshvault_zip_url, "-o", "package.zip")
	err = sshvault_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshvault_bin_url := "https://github.com/ssh-vault/ssh-vault/archive/refs/tags/1.0.13.bin"
	sshvault_cmd_bin := exec.Command("curl", "-L", sshvault_bin_url, "-o", "binary.bin")
	err = sshvault_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshvault_src_url := "https://github.com/ssh-vault/ssh-vault/archive/refs/tags/1.0.13.src.tar.gz"
	sshvault_cmd_src := exec.Command("curl", "-L", sshvault_src_url, "-o", "source.tar.gz")
	err = sshvault_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshvault_cmd_direct := exec.Command("./binary")
	err = sshvault_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
