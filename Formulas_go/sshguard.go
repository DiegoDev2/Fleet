package main

// Sshguard - Protect from brute force attacks against SSH
// Homepage: https://www.sshguard.net/

import (
	"fmt"
	
	"os/exec"
)

func installSshguard() {
	// Método 1: Descargar y extraer .tar.gz
	sshguard_tar_url := "https://downloads.sourceforge.net/project/sshguard/sshguard/2.4.3/sshguard-2.4.3.tar.gz"
	sshguard_cmd_tar := exec.Command("curl", "-L", sshguard_tar_url, "-o", "package.tar.gz")
	err := sshguard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshguard_zip_url := "https://downloads.sourceforge.net/project/sshguard/sshguard/2.4.3/sshguard-2.4.3.zip"
	sshguard_cmd_zip := exec.Command("curl", "-L", sshguard_zip_url, "-o", "package.zip")
	err = sshguard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshguard_bin_url := "https://downloads.sourceforge.net/project/sshguard/sshguard/2.4.3/sshguard-2.4.3.bin"
	sshguard_cmd_bin := exec.Command("curl", "-L", sshguard_bin_url, "-o", "binary.bin")
	err = sshguard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshguard_src_url := "https://downloads.sourceforge.net/project/sshguard/sshguard/2.4.3/sshguard-2.4.3.src.tar.gz"
	sshguard_cmd_src := exec.Command("curl", "-L", sshguard_src_url, "-o", "source.tar.gz")
	err = sshguard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshguard_cmd_direct := exec.Command("./binary")
	err = sshguard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: docutils")
exec.Command("latte", "install", "docutils")
}
