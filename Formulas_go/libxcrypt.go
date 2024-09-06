package main

// Libxcrypt - Extended crypt library for descrypt, md5crypt, bcrypt, and others
// Homepage: https://github.com/besser82/libxcrypt

import (
	"fmt"
	
	"os/exec"
)

func installLibxcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	libxcrypt_tar_url := "https://github.com/besser82/libxcrypt/releases/download/v4.4.36/libxcrypt-4.4.36.tar.xz"
	libxcrypt_cmd_tar := exec.Command("curl", "-L", libxcrypt_tar_url, "-o", "package.tar.gz")
	err := libxcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxcrypt_zip_url := "https://github.com/besser82/libxcrypt/releases/download/v4.4.36/libxcrypt-4.4.36.tar.xz"
	libxcrypt_cmd_zip := exec.Command("curl", "-L", libxcrypt_zip_url, "-o", "package.zip")
	err = libxcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxcrypt_bin_url := "https://github.com/besser82/libxcrypt/releases/download/v4.4.36/libxcrypt-4.4.36.tar.xz"
	libxcrypt_cmd_bin := exec.Command("curl", "-L", libxcrypt_bin_url, "-o", "binary.bin")
	err = libxcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxcrypt_src_url := "https://github.com/besser82/libxcrypt/releases/download/v4.4.36/libxcrypt-4.4.36.tar.xz"
	libxcrypt_cmd_src := exec.Command("curl", "-L", libxcrypt_src_url, "-o", "source.tar.gz")
	err = libxcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxcrypt_cmd_direct := exec.Command("./binary")
	err = libxcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
