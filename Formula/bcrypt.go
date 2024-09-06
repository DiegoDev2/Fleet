package main

// Bcrypt - Cross platform file encryption utility using blowfish
// Homepage: https://bcrypt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	bcrypt_tar_url := "https://bcrypt.sourceforge.net/bcrypt-1.1.tar.gz"
	bcrypt_cmd_tar := exec.Command("curl", "-L", bcrypt_tar_url, "-o", "package.tar.gz")
	err := bcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bcrypt_zip_url := "https://bcrypt.sourceforge.net/bcrypt-1.1.zip"
	bcrypt_cmd_zip := exec.Command("curl", "-L", bcrypt_zip_url, "-o", "package.zip")
	err = bcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bcrypt_bin_url := "https://bcrypt.sourceforge.net/bcrypt-1.1.bin"
	bcrypt_cmd_bin := exec.Command("curl", "-L", bcrypt_bin_url, "-o", "binary.bin")
	err = bcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bcrypt_src_url := "https://bcrypt.sourceforge.net/bcrypt-1.1.src.tar.gz"
	bcrypt_cmd_src := exec.Command("curl", "-L", bcrypt_src_url, "-o", "source.tar.gz")
	err = bcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bcrypt_cmd_direct := exec.Command("./binary")
	err = bcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
