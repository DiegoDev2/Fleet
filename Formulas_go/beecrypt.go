package main

// Beecrypt - C/C++ cryptography library
// Homepage: https://beecrypt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBeecrypt() {
	// Método 1: Descargar y extraer .tar.gz
	beecrypt_tar_url := "https://downloads.sourceforge.net/project/beecrypt/beecrypt/4.2.1/beecrypt-4.2.1.tar.gz"
	beecrypt_cmd_tar := exec.Command("curl", "-L", beecrypt_tar_url, "-o", "package.tar.gz")
	err := beecrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beecrypt_zip_url := "https://downloads.sourceforge.net/project/beecrypt/beecrypt/4.2.1/beecrypt-4.2.1.zip"
	beecrypt_cmd_zip := exec.Command("curl", "-L", beecrypt_zip_url, "-o", "package.zip")
	err = beecrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beecrypt_bin_url := "https://downloads.sourceforge.net/project/beecrypt/beecrypt/4.2.1/beecrypt-4.2.1.bin"
	beecrypt_cmd_bin := exec.Command("curl", "-L", beecrypt_bin_url, "-o", "binary.bin")
	err = beecrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beecrypt_src_url := "https://downloads.sourceforge.net/project/beecrypt/beecrypt/4.2.1/beecrypt-4.2.1.src.tar.gz"
	beecrypt_cmd_src := exec.Command("curl", "-L", beecrypt_src_url, "-o", "source.tar.gz")
	err = beecrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beecrypt_cmd_direct := exec.Command("./binary")
	err = beecrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
