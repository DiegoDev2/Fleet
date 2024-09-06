package main

// Dillo - Fast and small graphical web browser
// Homepage: https://dillo-browser.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installDillo() {
	// Método 1: Descargar y extraer .tar.gz
	dillo_tar_url := "https://github.com/dillo-browser/dillo/releases/download/v3.1.1/dillo-3.1.1.tar.bz2"
	dillo_cmd_tar := exec.Command("curl", "-L", dillo_tar_url, "-o", "package.tar.gz")
	err := dillo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dillo_zip_url := "https://github.com/dillo-browser/dillo/releases/download/v3.1.1/dillo-3.1.1.tar.bz2"
	dillo_cmd_zip := exec.Command("curl", "-L", dillo_zip_url, "-o", "package.zip")
	err = dillo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dillo_bin_url := "https://github.com/dillo-browser/dillo/releases/download/v3.1.1/dillo-3.1.1.tar.bz2"
	dillo_cmd_bin := exec.Command("curl", "-L", dillo_bin_url, "-o", "binary.bin")
	err = dillo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dillo_src_url := "https://github.com/dillo-browser/dillo/releases/download/v3.1.1/dillo-3.1.1.tar.bz2"
	dillo_cmd_src := exec.Command("curl", "-L", dillo_src_url, "-o", "source.tar.gz")
	err = dillo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dillo_cmd_direct := exec.Command("./binary")
	err = dillo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: fltk")
	exec.Command("latte", "install", "fltk").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
