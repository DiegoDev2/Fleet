package main

// VaultCli - Subversion-like utility to work with Jackrabbit FileVault
// Homepage: https://jackrabbit.apache.org/filevault/index.html

import (
	"fmt"
	
	"os/exec"
)

func installVaultCli() {
	// Método 1: Descargar y extraer .tar.gz
	vaultcli_tar_url := "https://search.maven.org/remotecontent?filepath=org/apache/jackrabbit/vault/vault-cli/3.8.0/vault-cli-3.8.0-bin.tar.gz"
	vaultcli_cmd_tar := exec.Command("curl", "-L", vaultcli_tar_url, "-o", "package.tar.gz")
	err := vaultcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vaultcli_zip_url := "https://search.maven.org/remotecontent?filepath=org/apache/jackrabbit/vault/vault-cli/3.8.0/vault-cli-3.8.0-bin.zip"
	vaultcli_cmd_zip := exec.Command("curl", "-L", vaultcli_zip_url, "-o", "package.zip")
	err = vaultcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vaultcli_bin_url := "https://search.maven.org/remotecontent?filepath=org/apache/jackrabbit/vault/vault-cli/3.8.0/vault-cli-3.8.0-bin.bin"
	vaultcli_cmd_bin := exec.Command("curl", "-L", vaultcli_bin_url, "-o", "binary.bin")
	err = vaultcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vaultcli_src_url := "https://search.maven.org/remotecontent?filepath=org/apache/jackrabbit/vault/vault-cli/3.8.0/vault-cli-3.8.0-bin.src.tar.gz"
	vaultcli_cmd_src := exec.Command("curl", "-L", vaultcli_src_url, "-o", "source.tar.gz")
	err = vaultcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vaultcli_cmd_direct := exec.Command("./binary")
	err = vaultcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
