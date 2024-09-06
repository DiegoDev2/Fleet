package main

// PamU2f - Provides an easy way to use U2F-compliant authenticators with PAM
// Homepage: https://developers.yubico.com/pam-u2f/

import (
	"fmt"
	
	"os/exec"
)

func installPamU2f() {
	// Método 1: Descargar y extraer .tar.gz
	pamu2f_tar_url := "https://developers.yubico.com/pam-u2f/Releases/pam_u2f-1.3.0.tar.gz"
	pamu2f_cmd_tar := exec.Command("curl", "-L", pamu2f_tar_url, "-o", "package.tar.gz")
	err := pamu2f_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pamu2f_zip_url := "https://developers.yubico.com/pam-u2f/Releases/pam_u2f-1.3.0.zip"
	pamu2f_cmd_zip := exec.Command("curl", "-L", pamu2f_zip_url, "-o", "package.zip")
	err = pamu2f_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pamu2f_bin_url := "https://developers.yubico.com/pam-u2f/Releases/pam_u2f-1.3.0.bin"
	pamu2f_cmd_bin := exec.Command("curl", "-L", pamu2f_bin_url, "-o", "binary.bin")
	err = pamu2f_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pamu2f_src_url := "https://developers.yubico.com/pam-u2f/Releases/pam_u2f-1.3.0.src.tar.gz"
	pamu2f_cmd_src := exec.Command("curl", "-L", pamu2f_src_url, "-o", "source.tar.gz")
	err = pamu2f_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pamu2f_cmd_direct := exec.Command("./binary")
	err = pamu2f_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
