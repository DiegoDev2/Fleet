package main

// Aescrypt - Program for encryption/decryption
// Homepage: https://aescrypt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAescrypt() {
	// Método 1: Descargar y extraer .tar.gz
	aescrypt_tar_url := "https://aescrypt.sourceforge.net/aescrypt-0.7.tar.gz"
	aescrypt_cmd_tar := exec.Command("curl", "-L", aescrypt_tar_url, "-o", "package.tar.gz")
	err := aescrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aescrypt_zip_url := "https://aescrypt.sourceforge.net/aescrypt-0.7.zip"
	aescrypt_cmd_zip := exec.Command("curl", "-L", aescrypt_zip_url, "-o", "package.zip")
	err = aescrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aescrypt_bin_url := "https://aescrypt.sourceforge.net/aescrypt-0.7.bin"
	aescrypt_cmd_bin := exec.Command("curl", "-L", aescrypt_bin_url, "-o", "binary.bin")
	err = aescrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aescrypt_src_url := "https://aescrypt.sourceforge.net/aescrypt-0.7.src.tar.gz"
	aescrypt_cmd_src := exec.Command("curl", "-L", aescrypt_src_url, "-o", "source.tar.gz")
	err = aescrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aescrypt_cmd_direct := exec.Command("./binary")
	err = aescrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
