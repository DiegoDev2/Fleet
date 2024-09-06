package main

// Ncdc - NCurses direct connect
// Homepage: https://dev.yorhel.nl/ncdc

import (
	"fmt"
	
	"os/exec"
)

func installNcdc() {
	// Método 1: Descargar y extraer .tar.gz
	ncdc_tar_url := "https://dev.yorhel.nl/download/ncdc-1.24.1.tar.gz"
	ncdc_cmd_tar := exec.Command("curl", "-L", ncdc_tar_url, "-o", "package.tar.gz")
	err := ncdc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncdc_zip_url := "https://dev.yorhel.nl/download/ncdc-1.24.1.zip"
	ncdc_cmd_zip := exec.Command("curl", "-L", ncdc_zip_url, "-o", "package.zip")
	err = ncdc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncdc_bin_url := "https://dev.yorhel.nl/download/ncdc-1.24.1.bin"
	ncdc_cmd_bin := exec.Command("curl", "-L", ncdc_bin_url, "-o", "binary.bin")
	err = ncdc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncdc_src_url := "https://dev.yorhel.nl/download/ncdc-1.24.1.src.tar.gz"
	ncdc_cmd_src := exec.Command("curl", "-L", ncdc_src_url, "-o", "source.tar.gz")
	err = ncdc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncdc_cmd_direct := exec.Command("./binary")
	err = ncdc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
