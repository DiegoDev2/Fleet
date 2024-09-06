package main

// GnupgPkcs11Scd - Enable the use of PKCS#11 tokens with GnuPG
// Homepage: https://gnupg-pkcs11.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGnupgPkcs11Scd() {
	// Método 1: Descargar y extraer .tar.gz
	gnupgpkcs11scd_tar_url := "https://github.com/alonbl/gnupg-pkcs11-scd/releases/download/gnupg-pkcs11-scd-0.10.0/gnupg-pkcs11-scd-0.10.0.tar.bz2"
	gnupgpkcs11scd_cmd_tar := exec.Command("curl", "-L", gnupgpkcs11scd_tar_url, "-o", "package.tar.gz")
	err := gnupgpkcs11scd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnupgpkcs11scd_zip_url := "https://github.com/alonbl/gnupg-pkcs11-scd/releases/download/gnupg-pkcs11-scd-0.10.0/gnupg-pkcs11-scd-0.10.0.tar.bz2"
	gnupgpkcs11scd_cmd_zip := exec.Command("curl", "-L", gnupgpkcs11scd_zip_url, "-o", "package.zip")
	err = gnupgpkcs11scd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnupgpkcs11scd_bin_url := "https://github.com/alonbl/gnupg-pkcs11-scd/releases/download/gnupg-pkcs11-scd-0.10.0/gnupg-pkcs11-scd-0.10.0.tar.bz2"
	gnupgpkcs11scd_cmd_bin := exec.Command("curl", "-L", gnupgpkcs11scd_bin_url, "-o", "binary.bin")
	err = gnupgpkcs11scd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnupgpkcs11scd_src_url := "https://github.com/alonbl/gnupg-pkcs11-scd/releases/download/gnupg-pkcs11-scd-0.10.0/gnupg-pkcs11-scd-0.10.0.tar.bz2"
	gnupgpkcs11scd_cmd_src := exec.Command("curl", "-L", gnupgpkcs11scd_src_url, "-o", "source.tar.gz")
	err = gnupgpkcs11scd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnupgpkcs11scd_cmd_direct := exec.Command("./binary")
	err = gnupgpkcs11scd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libassuan@2")
exec.Command("latte", "install", "libassuan@2")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkcs11-helper")
exec.Command("latte", "install", "pkcs11-helper")
}
