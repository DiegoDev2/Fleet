package main

// GnupgAT22 - GNU Pretty Good Privacy (PGP) package
// Homepage: https://gnupg.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnupgAT22() {
	// Método 1: Descargar y extraer .tar.gz
	gnupgat22_tar_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.44.tar.bz2"
	gnupgat22_cmd_tar := exec.Command("curl", "-L", gnupgat22_tar_url, "-o", "package.tar.gz")
	err := gnupgat22_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnupgat22_zip_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.44.tar.bz2"
	gnupgat22_cmd_zip := exec.Command("curl", "-L", gnupgat22_zip_url, "-o", "package.zip")
	err = gnupgat22_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnupgat22_bin_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.44.tar.bz2"
	gnupgat22_cmd_bin := exec.Command("curl", "-L", gnupgat22_bin_url, "-o", "binary.bin")
	err = gnupgat22_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnupgat22_src_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.2.44.tar.bz2"
	gnupgat22_cmd_src := exec.Command("curl", "-L", gnupgat22_src_url, "-o", "source.tar.gz")
	err = gnupgat22_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnupgat22_cmd_direct := exec.Command("./binary")
	err = gnupgat22_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: libksba")
exec.Command("latte", "install", "libksba")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: npth")
exec.Command("latte", "install", "npth")
	fmt.Println("Instalando dependencia: pinentry")
exec.Command("latte", "install", "pinentry")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
