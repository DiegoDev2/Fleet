package main

// Rarian - Documentation metadata library
// Homepage: https://rarian.freedesktop.org/

import (
	"fmt"
	
	"os/exec"
)

func installRarian() {
	// Método 1: Descargar y extraer .tar.gz
	rarian_tar_url := "https://gitlab.freedesktop.org/rarian/rarian/-/releases/0.8.5/downloads/assets/rarian-0.8.5.tar.bz2"
	rarian_cmd_tar := exec.Command("curl", "-L", rarian_tar_url, "-o", "package.tar.gz")
	err := rarian_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rarian_zip_url := "https://gitlab.freedesktop.org/rarian/rarian/-/releases/0.8.5/downloads/assets/rarian-0.8.5.tar.bz2"
	rarian_cmd_zip := exec.Command("curl", "-L", rarian_zip_url, "-o", "package.zip")
	err = rarian_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rarian_bin_url := "https://gitlab.freedesktop.org/rarian/rarian/-/releases/0.8.5/downloads/assets/rarian-0.8.5.tar.bz2"
	rarian_cmd_bin := exec.Command("curl", "-L", rarian_bin_url, "-o", "binary.bin")
	err = rarian_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rarian_src_url := "https://gitlab.freedesktop.org/rarian/rarian/-/releases/0.8.5/downloads/assets/rarian-0.8.5.tar.bz2"
	rarian_cmd_src := exec.Command("curl", "-L", rarian_src_url, "-o", "source.tar.gz")
	err = rarian_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rarian_cmd_direct := exec.Command("./binary")
	err = rarian_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: tinyxml")
exec.Command("latte", "install", "tinyxml")
}
