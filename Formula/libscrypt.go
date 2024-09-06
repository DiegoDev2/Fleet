package main

// Libscrypt - Library for scrypt
// Homepage: https://github.com/technion/libscrypt

import (
	"fmt"
	
	"os/exec"
)

func installLibscrypt() {
	// Método 1: Descargar y extraer .tar.gz
	libscrypt_tar_url := "https://github.com/technion/libscrypt/archive/refs/tags/v1.22.tar.gz"
	libscrypt_cmd_tar := exec.Command("curl", "-L", libscrypt_tar_url, "-o", "package.tar.gz")
	err := libscrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libscrypt_zip_url := "https://github.com/technion/libscrypt/archive/refs/tags/v1.22.zip"
	libscrypt_cmd_zip := exec.Command("curl", "-L", libscrypt_zip_url, "-o", "package.zip")
	err = libscrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libscrypt_bin_url := "https://github.com/technion/libscrypt/archive/refs/tags/v1.22.bin"
	libscrypt_cmd_bin := exec.Command("curl", "-L", libscrypt_bin_url, "-o", "binary.bin")
	err = libscrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libscrypt_src_url := "https://github.com/technion/libscrypt/archive/refs/tags/v1.22.src.tar.gz"
	libscrypt_cmd_src := exec.Command("curl", "-L", libscrypt_src_url, "-o", "source.tar.gz")
	err = libscrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libscrypt_cmd_direct := exec.Command("./binary")
	err = libscrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
