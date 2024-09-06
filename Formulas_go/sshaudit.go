package main

// SshAudit - SSH server & client auditing
// Homepage: https://github.com/jtesta/ssh-audit

import (
	"fmt"
	
	"os/exec"
)

func installSshAudit() {
	// Método 1: Descargar y extraer .tar.gz
	sshaudit_tar_url := "https://files.pythonhosted.org/packages/f1/26/5b724f1ade0a40aeea41cf39e7db497209a97b947b48acf378bf7630fa87/ssh_audit-3.2.0.tar.gz"
	sshaudit_cmd_tar := exec.Command("curl", "-L", sshaudit_tar_url, "-o", "package.tar.gz")
	err := sshaudit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshaudit_zip_url := "https://files.pythonhosted.org/packages/f1/26/5b724f1ade0a40aeea41cf39e7db497209a97b947b48acf378bf7630fa87/ssh_audit-3.2.0.zip"
	sshaudit_cmd_zip := exec.Command("curl", "-L", sshaudit_zip_url, "-o", "package.zip")
	err = sshaudit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshaudit_bin_url := "https://files.pythonhosted.org/packages/f1/26/5b724f1ade0a40aeea41cf39e7db497209a97b947b48acf378bf7630fa87/ssh_audit-3.2.0.bin"
	sshaudit_cmd_bin := exec.Command("curl", "-L", sshaudit_bin_url, "-o", "binary.bin")
	err = sshaudit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshaudit_src_url := "https://files.pythonhosted.org/packages/f1/26/5b724f1ade0a40aeea41cf39e7db497209a97b947b48acf378bf7630fa87/ssh_audit-3.2.0.src.tar.gz"
	sshaudit_cmd_src := exec.Command("curl", "-L", sshaudit_src_url, "-o", "source.tar.gz")
	err = sshaudit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshaudit_cmd_direct := exec.Command("./binary")
	err = sshaudit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
