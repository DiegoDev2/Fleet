package main

// SshPermitA38 - Central management and deployment for SSH keys
// Homepage: https://github.com/ierror/ssh-permit-a38

import (
	"fmt"
	
	"os/exec"
)

func installSshPermitA38() {
	// Método 1: Descargar y extraer .tar.gz
	sshpermita38_tar_url := "https://github.com/ierror/ssh-permit-a38/archive/refs/tags/v0.2.0.tar.gz"
	sshpermita38_cmd_tar := exec.Command("curl", "-L", sshpermita38_tar_url, "-o", "package.tar.gz")
	err := sshpermita38_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshpermita38_zip_url := "https://github.com/ierror/ssh-permit-a38/archive/refs/tags/v0.2.0.zip"
	sshpermita38_cmd_zip := exec.Command("curl", "-L", sshpermita38_zip_url, "-o", "package.zip")
	err = sshpermita38_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshpermita38_bin_url := "https://github.com/ierror/ssh-permit-a38/archive/refs/tags/v0.2.0.bin"
	sshpermita38_cmd_bin := exec.Command("curl", "-L", sshpermita38_bin_url, "-o", "binary.bin")
	err = sshpermita38_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshpermita38_src_url := "https://github.com/ierror/ssh-permit-a38/archive/refs/tags/v0.2.0.src.tar.gz"
	sshpermita38_cmd_src := exec.Command("curl", "-L", sshpermita38_src_url, "-o", "source.tar.gz")
	err = sshpermita38_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshpermita38_cmd_direct := exec.Command("./binary")
	err = sshpermita38_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
}
