package main

// Libass - Subtitle renderer for the ASS/SSA subtitle format
// Homepage: https://github.com/libass/libass

import (
	"fmt"
	
	"os/exec"
)

func installLibass() {
	// Método 1: Descargar y extraer .tar.gz
	libass_tar_url := "https://github.com/libass/libass/releases/download/0.17.3/libass-0.17.3.tar.xz"
	libass_cmd_tar := exec.Command("curl", "-L", libass_tar_url, "-o", "package.tar.gz")
	err := libass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libass_zip_url := "https://github.com/libass/libass/releases/download/0.17.3/libass-0.17.3.tar.xz"
	libass_cmd_zip := exec.Command("curl", "-L", libass_zip_url, "-o", "package.zip")
	err = libass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libass_bin_url := "https://github.com/libass/libass/releases/download/0.17.3/libass-0.17.3.tar.xz"
	libass_cmd_bin := exec.Command("curl", "-L", libass_bin_url, "-o", "binary.bin")
	err = libass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libass_src_url := "https://github.com/libass/libass/releases/download/0.17.3/libass-0.17.3.tar.xz"
	libass_cmd_src := exec.Command("curl", "-L", libass_src_url, "-o", "source.tar.gz")
	err = libass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libass_cmd_direct := exec.Command("./binary")
	err = libass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libunibreak")
	exec.Command("latte", "install", "libunibreak").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}
