package main

// Lnav - Curses-based tool for viewing and analyzing log files
// Homepage: https://lnav.org/

import (
	"fmt"
	
	"os/exec"
)

func installLnav() {
	// Método 1: Descargar y extraer .tar.gz
	lnav_tar_url := "https://github.com/tstack/lnav/releases/download/v0.12.2/lnav-0.12.2.tar.gz"
	lnav_cmd_tar := exec.Command("curl", "-L", lnav_tar_url, "-o", "package.tar.gz")
	err := lnav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lnav_zip_url := "https://github.com/tstack/lnav/releases/download/v0.12.2/lnav-0.12.2.zip"
	lnav_cmd_zip := exec.Command("curl", "-L", lnav_zip_url, "-o", "package.zip")
	err = lnav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lnav_bin_url := "https://github.com/tstack/lnav/releases/download/v0.12.2/lnav-0.12.2.bin"
	lnav_cmd_bin := exec.Command("curl", "-L", lnav_bin_url, "-o", "binary.bin")
	err = lnav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lnav_src_url := "https://github.com/tstack/lnav/releases/download/v0.12.2/lnav-0.12.2.src.tar.gz"
	lnav_cmd_src := exec.Command("curl", "-L", lnav_src_url, "-o", "source.tar.gz")
	err = lnav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lnav_cmd_direct := exec.Command("./binary")
	err = lnav_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: re2c")
exec.Command("latte", "install", "re2c")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
}
