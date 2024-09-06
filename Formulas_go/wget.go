package main

// Wget - Internet file retriever
// Homepage: https://www.gnu.org/software/wget/

import (
	"fmt"
	
	"os/exec"
)

func installWget() {
	// Método 1: Descargar y extraer .tar.gz
	wget_tar_url := "https://ftp.gnu.org/gnu/wget/wget-1.24.5.tar.gz"
	wget_cmd_tar := exec.Command("curl", "-L", wget_tar_url, "-o", "package.tar.gz")
	err := wget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wget_zip_url := "https://ftp.gnu.org/gnu/wget/wget-1.24.5.zip"
	wget_cmd_zip := exec.Command("curl", "-L", wget_zip_url, "-o", "package.zip")
	err = wget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wget_bin_url := "https://ftp.gnu.org/gnu/wget/wget-1.24.5.bin"
	wget_cmd_bin := exec.Command("curl", "-L", wget_bin_url, "-o", "binary.bin")
	err = wget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wget_src_url := "https://ftp.gnu.org/gnu/wget/wget-1.24.5.src.tar.gz"
	wget_cmd_src := exec.Command("curl", "-L", wget_src_url, "-o", "source.tar.gz")
	err = wget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wget_cmd_direct := exec.Command("./binary")
	err = wget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
