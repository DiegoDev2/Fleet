package main

// Ccrypt - Encrypt and decrypt files and streams
// Homepage: https://ccrypt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	ccrypt_tar_url := "https://downloads.sourceforge.net/project/ccrypt/1.11/ccrypt-1.11.tar.gz"
	ccrypt_cmd_tar := exec.Command("curl", "-L", ccrypt_tar_url, "-o", "package.tar.gz")
	err := ccrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccrypt_zip_url := "https://downloads.sourceforge.net/project/ccrypt/1.11/ccrypt-1.11.zip"
	ccrypt_cmd_zip := exec.Command("curl", "-L", ccrypt_zip_url, "-o", "package.zip")
	err = ccrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccrypt_bin_url := "https://downloads.sourceforge.net/project/ccrypt/1.11/ccrypt-1.11.bin"
	ccrypt_cmd_bin := exec.Command("curl", "-L", ccrypt_bin_url, "-o", "binary.bin")
	err = ccrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccrypt_src_url := "https://downloads.sourceforge.net/project/ccrypt/1.11/ccrypt-1.11.src.tar.gz"
	ccrypt_cmd_src := exec.Command("curl", "-L", ccrypt_src_url, "-o", "source.tar.gz")
	err = ccrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccrypt_cmd_direct := exec.Command("./binary")
	err = ccrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
