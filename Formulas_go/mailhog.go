package main

// Mailhog - Web and API based SMTP testing tool
// Homepage: https://github.com/mailhog/MailHog

import (
	"fmt"
	
	"os/exec"
)

func installMailhog() {
	// Método 1: Descargar y extraer .tar.gz
	mailhog_tar_url := "https://github.com/mailhog/MailHog/archive/refs/tags/v1.0.1.tar.gz"
	mailhog_cmd_tar := exec.Command("curl", "-L", mailhog_tar_url, "-o", "package.tar.gz")
	err := mailhog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mailhog_zip_url := "https://github.com/mailhog/MailHog/archive/refs/tags/v1.0.1.zip"
	mailhog_cmd_zip := exec.Command("curl", "-L", mailhog_zip_url, "-o", "package.zip")
	err = mailhog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mailhog_bin_url := "https://github.com/mailhog/MailHog/archive/refs/tags/v1.0.1.bin"
	mailhog_cmd_bin := exec.Command("curl", "-L", mailhog_bin_url, "-o", "binary.bin")
	err = mailhog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mailhog_src_url := "https://github.com/mailhog/MailHog/archive/refs/tags/v1.0.1.src.tar.gz"
	mailhog_cmd_src := exec.Command("curl", "-L", mailhog_src_url, "-o", "source.tar.gz")
	err = mailhog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mailhog_cmd_direct := exec.Command("./binary")
	err = mailhog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
