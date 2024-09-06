package main

// Stoken - Tokencode generator compatible with RSA SecurID 128-bit (AES)
// Homepage: https://github.com/stoken-dev/stoken

import (
	"fmt"
	
	"os/exec"
)

func installStoken() {
	// Método 1: Descargar y extraer .tar.gz
	stoken_tar_url := "https://github.com/stoken-dev/stoken/archive/refs/tags/v0.93.tar.gz"
	stoken_cmd_tar := exec.Command("curl", "-L", stoken_tar_url, "-o", "package.tar.gz")
	err := stoken_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stoken_zip_url := "https://github.com/stoken-dev/stoken/archive/refs/tags/v0.93.zip"
	stoken_cmd_zip := exec.Command("curl", "-L", stoken_zip_url, "-o", "package.zip")
	err = stoken_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stoken_bin_url := "https://github.com/stoken-dev/stoken/archive/refs/tags/v0.93.bin"
	stoken_cmd_bin := exec.Command("curl", "-L", stoken_bin_url, "-o", "binary.bin")
	err = stoken_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stoken_src_url := "https://github.com/stoken-dev/stoken/archive/refs/tags/v0.93.src.tar.gz"
	stoken_cmd_src := exec.Command("curl", "-L", stoken_src_url, "-o", "source.tar.gz")
	err = stoken_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stoken_cmd_direct := exec.Command("./binary")
	err = stoken_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
}
