package main

// Screen - Terminal multiplexer with VT100/ANSI terminal emulation
// Homepage: https://www.gnu.org/software/screen/

import (
	"fmt"
	
	"os/exec"
)

func installScreen() {
	// Método 1: Descargar y extraer .tar.gz
	screen_tar_url := "https://ftp.gnu.org/gnu/screen/screen-5.0.0.tar.gz"
	screen_cmd_tar := exec.Command("curl", "-L", screen_tar_url, "-o", "package.tar.gz")
	err := screen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	screen_zip_url := "https://ftp.gnu.org/gnu/screen/screen-5.0.0.zip"
	screen_cmd_zip := exec.Command("curl", "-L", screen_zip_url, "-o", "package.zip")
	err = screen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	screen_bin_url := "https://ftp.gnu.org/gnu/screen/screen-5.0.0.bin"
	screen_cmd_bin := exec.Command("curl", "-L", screen_bin_url, "-o", "binary.bin")
	err = screen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	screen_src_url := "https://ftp.gnu.org/gnu/screen/screen-5.0.0.src.tar.gz"
	screen_cmd_src := exec.Command("curl", "-L", screen_src_url, "-o", "source.tar.gz")
	err = screen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	screen_cmd_direct := exec.Command("./binary")
	err = screen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
