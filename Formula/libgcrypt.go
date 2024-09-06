package main

// Libgcrypt - Cryptographic library based on the code from GnuPG
// Homepage: https://gnupg.org/related_software/libgcrypt/

import (
	"fmt"
	
	"os/exec"
)

func installLibgcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	libgcrypt_tar_url := "https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.10.3.tar.bz2"
	libgcrypt_cmd_tar := exec.Command("curl", "-L", libgcrypt_tar_url, "-o", "package.tar.gz")
	err := libgcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgcrypt_zip_url := "https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.10.3.tar.bz2"
	libgcrypt_cmd_zip := exec.Command("curl", "-L", libgcrypt_zip_url, "-o", "package.zip")
	err = libgcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgcrypt_bin_url := "https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.10.3.tar.bz2"
	libgcrypt_cmd_bin := exec.Command("curl", "-L", libgcrypt_bin_url, "-o", "binary.bin")
	err = libgcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgcrypt_src_url := "https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.10.3.tar.bz2"
	libgcrypt_cmd_src := exec.Command("curl", "-L", libgcrypt_src_url, "-o", "source.tar.gz")
	err = libgcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgcrypt_cmd_direct := exec.Command("./binary")
	err = libgcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
