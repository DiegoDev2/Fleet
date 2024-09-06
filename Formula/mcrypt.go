package main

// Mcrypt - Replacement for the old crypt package and crypt(1) command
// Homepage: https://mcrypt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	mcrypt_tar_url := "https://downloads.sourceforge.net/project/mcrypt/MCrypt/2.6.8/mcrypt-2.6.8.tar.gz"
	mcrypt_cmd_tar := exec.Command("curl", "-L", mcrypt_tar_url, "-o", "package.tar.gz")
	err := mcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcrypt_zip_url := "https://downloads.sourceforge.net/project/mcrypt/MCrypt/2.6.8/mcrypt-2.6.8.zip"
	mcrypt_cmd_zip := exec.Command("curl", "-L", mcrypt_zip_url, "-o", "package.zip")
	err = mcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcrypt_bin_url := "https://downloads.sourceforge.net/project/mcrypt/MCrypt/2.6.8/mcrypt-2.6.8.bin"
	mcrypt_cmd_bin := exec.Command("curl", "-L", mcrypt_bin_url, "-o", "binary.bin")
	err = mcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcrypt_src_url := "https://downloads.sourceforge.net/project/mcrypt/MCrypt/2.6.8/mcrypt-2.6.8.src.tar.gz"
	mcrypt_cmd_src := exec.Command("curl", "-L", mcrypt_src_url, "-o", "source.tar.gz")
	err = mcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcrypt_cmd_direct := exec.Command("./binary")
	err = mcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: mhash")
	exec.Command("latte", "install", "mhash").Run()
}
