package main

// Gawk - GNU awk utility
// Homepage: https://www.gnu.org/software/gawk/

import (
	"fmt"
	
	"os/exec"
)

func installGawk() {
	// Método 1: Descargar y extraer .tar.gz
	gawk_tar_url := "https://ftp.gnu.org/gnu/gawk/gawk-5.3.0.tar.xz"
	gawk_cmd_tar := exec.Command("curl", "-L", gawk_tar_url, "-o", "package.tar.gz")
	err := gawk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gawk_zip_url := "https://ftp.gnu.org/gnu/gawk/gawk-5.3.0.tar.xz"
	gawk_cmd_zip := exec.Command("curl", "-L", gawk_zip_url, "-o", "package.zip")
	err = gawk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gawk_bin_url := "https://ftp.gnu.org/gnu/gawk/gawk-5.3.0.tar.xz"
	gawk_cmd_bin := exec.Command("curl", "-L", gawk_bin_url, "-o", "binary.bin")
	err = gawk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gawk_src_url := "https://ftp.gnu.org/gnu/gawk/gawk-5.3.0.tar.xz"
	gawk_cmd_src := exec.Command("curl", "-L", gawk_src_url, "-o", "source.tar.gz")
	err = gawk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gawk_cmd_direct := exec.Command("./binary")
	err = gawk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
