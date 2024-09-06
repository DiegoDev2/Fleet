package main

// Passwdqc - Password/passphrase strength checking and enforcement toolset
// Homepage: https://www.openwall.com/passwdqc/

import (
	"fmt"
	
	"os/exec"
)

func installPasswdqc() {
	// Método 1: Descargar y extraer .tar.gz
	passwdqc_tar_url := "https://www.openwall.com/passwdqc/passwdqc-2.0.3.tar.gz"
	passwdqc_cmd_tar := exec.Command("curl", "-L", passwdqc_tar_url, "-o", "package.tar.gz")
	err := passwdqc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	passwdqc_zip_url := "https://www.openwall.com/passwdqc/passwdqc-2.0.3.zip"
	passwdqc_cmd_zip := exec.Command("curl", "-L", passwdqc_zip_url, "-o", "package.zip")
	err = passwdqc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	passwdqc_bin_url := "https://www.openwall.com/passwdqc/passwdqc-2.0.3.bin"
	passwdqc_cmd_bin := exec.Command("curl", "-L", passwdqc_bin_url, "-o", "binary.bin")
	err = passwdqc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	passwdqc_src_url := "https://www.openwall.com/passwdqc/passwdqc-2.0.3.src.tar.gz"
	passwdqc_cmd_src := exec.Command("curl", "-L", passwdqc_src_url, "-o", "source.tar.gz")
	err = passwdqc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	passwdqc_cmd_direct := exec.Command("./binary")
	err = passwdqc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
}
