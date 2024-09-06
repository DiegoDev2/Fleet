package main

// Libtomcrypt - Comprehensive, modular and portable cryptographic toolkit
// Homepage: https://www.libtom.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibtomcrypt() {
	// Método 1: Descargar y extraer .tar.gz
	libtomcrypt_tar_url := "https://github.com/libtom/libtomcrypt/archive/refs/tags/v1.18.2.tar.gz"
	libtomcrypt_cmd_tar := exec.Command("curl", "-L", libtomcrypt_tar_url, "-o", "package.tar.gz")
	err := libtomcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtomcrypt_zip_url := "https://github.com/libtom/libtomcrypt/archive/refs/tags/v1.18.2.zip"
	libtomcrypt_cmd_zip := exec.Command("curl", "-L", libtomcrypt_zip_url, "-o", "package.zip")
	err = libtomcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtomcrypt_bin_url := "https://github.com/libtom/libtomcrypt/archive/refs/tags/v1.18.2.bin"
	libtomcrypt_cmd_bin := exec.Command("curl", "-L", libtomcrypt_bin_url, "-o", "binary.bin")
	err = libtomcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtomcrypt_src_url := "https://github.com/libtom/libtomcrypt/archive/refs/tags/v1.18.2.src.tar.gz"
	libtomcrypt_cmd_src := exec.Command("curl", "-L", libtomcrypt_src_url, "-o", "source.tar.gz")
	err = libtomcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtomcrypt_cmd_direct := exec.Command("./binary")
	err = libtomcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libtommath")
	exec.Command("latte", "install", "libtommath").Run()
}
