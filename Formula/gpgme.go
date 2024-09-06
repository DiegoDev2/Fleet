package main

// Gpgme - Library access to GnuPG
// Homepage: https://www.gnupg.org/related_software/gpgme/

import (
	"fmt"
	
	"os/exec"
)

func installGpgme() {
	// Método 1: Descargar y extraer .tar.gz
	gpgme_tar_url := "https://www.gnupg.org/ftp/gcrypt/gpgme/gpgme-1.23.2.tar.bz2"
	gpgme_cmd_tar := exec.Command("curl", "-L", gpgme_tar_url, "-o", "package.tar.gz")
	err := gpgme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpgme_zip_url := "https://www.gnupg.org/ftp/gcrypt/gpgme/gpgme-1.23.2.tar.bz2"
	gpgme_cmd_zip := exec.Command("curl", "-L", gpgme_zip_url, "-o", "package.zip")
	err = gpgme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpgme_bin_url := "https://www.gnupg.org/ftp/gcrypt/gpgme/gpgme-1.23.2.tar.bz2"
	gpgme_cmd_bin := exec.Command("curl", "-L", gpgme_bin_url, "-o", "binary.bin")
	err = gpgme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpgme_src_url := "https://www.gnupg.org/ftp/gcrypt/gpgme/gpgme-1.23.2.tar.bz2"
	gpgme_cmd_src := exec.Command("curl", "-L", gpgme_src_url, "-o", "source.tar.gz")
	err = gpgme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpgme_cmd_direct := exec.Command("./binary")
	err = gpgme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: libassuan")
	exec.Command("latte", "install", "libassuan").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
