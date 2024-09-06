package main

// DuoUnix - Two-factor authentication for SSH
// Homepage: https://www.duosecurity.com/docs/duounix

import (
	"fmt"
	
	"os/exec"
)

func installDuoUnix() {
	// Método 1: Descargar y extraer .tar.gz
	duounix_tar_url := "https://github.com/duosecurity/duo_unix/archive/refs/tags/duo_unix-2.0.3.tar.gz"
	duounix_cmd_tar := exec.Command("curl", "-L", duounix_tar_url, "-o", "package.tar.gz")
	err := duounix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duounix_zip_url := "https://github.com/duosecurity/duo_unix/archive/refs/tags/duo_unix-2.0.3.zip"
	duounix_cmd_zip := exec.Command("curl", "-L", duounix_zip_url, "-o", "package.zip")
	err = duounix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duounix_bin_url := "https://github.com/duosecurity/duo_unix/archive/refs/tags/duo_unix-2.0.3.bin"
	duounix_cmd_bin := exec.Command("curl", "-L", duounix_bin_url, "-o", "binary.bin")
	err = duounix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duounix_src_url := "https://github.com/duosecurity/duo_unix/archive/refs/tags/duo_unix-2.0.3.src.tar.gz"
	duounix_cmd_src := exec.Command("curl", "-L", duounix_src_url, "-o", "source.tar.gz")
	err = duounix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duounix_cmd_direct := exec.Command("./binary")
	err = duounix_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
