package main

// Pinentry - Passphrase entry dialog utilizing the Assuan protocol
// Homepage: https://www.gnupg.org/related_software/pinentry/

import (
	"fmt"
	
	"os/exec"
)

func installPinentry() {
	// Método 1: Descargar y extraer .tar.gz
	pinentry_tar_url := "https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.1.tar.bz2"
	pinentry_cmd_tar := exec.Command("curl", "-L", pinentry_tar_url, "-o", "package.tar.gz")
	err := pinentry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinentry_zip_url := "https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.1.tar.bz2"
	pinentry_cmd_zip := exec.Command("curl", "-L", pinentry_zip_url, "-o", "package.zip")
	err = pinentry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinentry_bin_url := "https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.1.tar.bz2"
	pinentry_cmd_bin := exec.Command("curl", "-L", pinentry_bin_url, "-o", "binary.bin")
	err = pinentry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinentry_src_url := "https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.1.tar.bz2"
	pinentry_cmd_src := exec.Command("curl", "-L", pinentry_src_url, "-o", "source.tar.gz")
	err = pinentry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinentry_cmd_direct := exec.Command("./binary")
	err = pinentry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libassuan")
	exec.Command("latte", "install", "libassuan").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libsecret")
	exec.Command("latte", "install", "libsecret").Run()
}
