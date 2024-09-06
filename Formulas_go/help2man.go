package main

// Help2man - Automatically generate simple man pages
// Homepage: https://www.gnu.org/software/help2man/

import (
	"fmt"
	
	"os/exec"
)

func installHelp2man() {
	// Método 1: Descargar y extraer .tar.gz
	help2man_tar_url := "https://ftp.gnu.org/gnu/help2man/help2man-1.49.3.tar.xz"
	help2man_cmd_tar := exec.Command("curl", "-L", help2man_tar_url, "-o", "package.tar.gz")
	err := help2man_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	help2man_zip_url := "https://ftp.gnu.org/gnu/help2man/help2man-1.49.3.tar.xz"
	help2man_cmd_zip := exec.Command("curl", "-L", help2man_zip_url, "-o", "package.zip")
	err = help2man_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	help2man_bin_url := "https://ftp.gnu.org/gnu/help2man/help2man-1.49.3.tar.xz"
	help2man_cmd_bin := exec.Command("curl", "-L", help2man_bin_url, "-o", "binary.bin")
	err = help2man_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	help2man_src_url := "https://ftp.gnu.org/gnu/help2man/help2man-1.49.3.tar.xz"
	help2man_cmd_src := exec.Command("curl", "-L", help2man_src_url, "-o", "source.tar.gz")
	err = help2man_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	help2man_cmd_direct := exec.Command("./binary")
	err = help2man_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
}
