package main

// GnuChess - Chess-playing program
// Homepage: https://www.gnu.org/software/chess/

import (
	"fmt"
	
	"os/exec"
)

func installGnuChess() {
	// Método 1: Descargar y extraer .tar.gz
	gnuchess_tar_url := "https://ftp.gnu.org/gnu/chess/gnuchess-6.2.9.tar.gz"
	gnuchess_cmd_tar := exec.Command("curl", "-L", gnuchess_tar_url, "-o", "package.tar.gz")
	err := gnuchess_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuchess_zip_url := "https://ftp.gnu.org/gnu/chess/gnuchess-6.2.9.zip"
	gnuchess_cmd_zip := exec.Command("curl", "-L", gnuchess_zip_url, "-o", "package.zip")
	err = gnuchess_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuchess_bin_url := "https://ftp.gnu.org/gnu/chess/gnuchess-6.2.9.bin"
	gnuchess_cmd_bin := exec.Command("curl", "-L", gnuchess_bin_url, "-o", "binary.bin")
	err = gnuchess_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuchess_src_url := "https://ftp.gnu.org/gnu/chess/gnuchess-6.2.9.src.tar.gz"
	gnuchess_cmd_src := exec.Command("curl", "-L", gnuchess_src_url, "-o", "source.tar.gz")
	err = gnuchess_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuchess_cmd_direct := exec.Command("./binary")
	err = gnuchess_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
