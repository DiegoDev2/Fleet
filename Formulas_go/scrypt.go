package main

// Scrypt - Encrypt and decrypt files using memory-hard password function
// Homepage: https://www.tarsnap.com/scrypt.html

import (
	"fmt"
	
	"os/exec"
)

func installScrypt() {
	// Método 1: Descargar y extraer .tar.gz
	scrypt_tar_url := "https://www.tarsnap.com/scrypt/scrypt-1.3.2.tgz"
	scrypt_cmd_tar := exec.Command("curl", "-L", scrypt_tar_url, "-o", "package.tar.gz")
	err := scrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scrypt_zip_url := "https://www.tarsnap.com/scrypt/scrypt-1.3.2.tgz"
	scrypt_cmd_zip := exec.Command("curl", "-L", scrypt_zip_url, "-o", "package.zip")
	err = scrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scrypt_bin_url := "https://www.tarsnap.com/scrypt/scrypt-1.3.2.tgz"
	scrypt_cmd_bin := exec.Command("curl", "-L", scrypt_bin_url, "-o", "binary.bin")
	err = scrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scrypt_src_url := "https://www.tarsnap.com/scrypt/scrypt-1.3.2.tgz"
	scrypt_cmd_src := exec.Command("curl", "-L", scrypt_src_url, "-o", "source.tar.gz")
	err = scrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scrypt_cmd_direct := exec.Command("./binary")
	err = scrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
