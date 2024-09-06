package main

// ClawsMail - User-friendly, lightweight, and fast email client
// Homepage: https://www.claws-mail.org/

import (
	"fmt"
	
	"os/exec"
)

func installClawsMail() {
	// Método 1: Descargar y extraer .tar.gz
	clawsmail_tar_url := "https://www.claws-mail.org/releases/claws-mail-4.3.0.tar.gz"
	clawsmail_cmd_tar := exec.Command("curl", "-L", clawsmail_tar_url, "-o", "package.tar.gz")
	err := clawsmail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clawsmail_zip_url := "https://www.claws-mail.org/releases/claws-mail-4.3.0.zip"
	clawsmail_cmd_zip := exec.Command("curl", "-L", clawsmail_zip_url, "-o", "package.zip")
	err = clawsmail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clawsmail_bin_url := "https://www.claws-mail.org/releases/claws-mail-4.3.0.bin"
	clawsmail_cmd_bin := exec.Command("curl", "-L", clawsmail_bin_url, "-o", "binary.bin")
	err = clawsmail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clawsmail_src_url := "https://www.claws-mail.org/releases/claws-mail-4.3.0.src.tar.gz"
	clawsmail_cmd_src := exec.Command("curl", "-L", clawsmail_src_url, "-o", "source.tar.gz")
	err = clawsmail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clawsmail_cmd_direct := exec.Command("./binary")
	err = clawsmail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libetpan")
exec.Command("latte", "install", "libetpan")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
}
