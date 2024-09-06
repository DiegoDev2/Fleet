package main

// Libgr - GR framework: a graphics library for visualisation applications
// Homepage: https://gr-framework.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibgr() {
	// Método 1: Descargar y extraer .tar.gz
	libgr_tar_url := "https://github.com/sciapp/gr/archive/refs/tags/v0.73.7.tar.gz"
	libgr_cmd_tar := exec.Command("curl", "-L", libgr_tar_url, "-o", "package.tar.gz")
	err := libgr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgr_zip_url := "https://github.com/sciapp/gr/archive/refs/tags/v0.73.7.zip"
	libgr_cmd_zip := exec.Command("curl", "-L", libgr_zip_url, "-o", "package.zip")
	err = libgr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgr_bin_url := "https://github.com/sciapp/gr/archive/refs/tags/v0.73.7.bin"
	libgr_cmd_bin := exec.Command("curl", "-L", libgr_bin_url, "-o", "binary.bin")
	err = libgr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgr_src_url := "https://github.com/sciapp/gr/archive/refs/tags/v0.73.7.src.tar.gz"
	libgr_cmd_src := exec.Command("curl", "-L", libgr_src_url, "-o", "source.tar.gz")
	err = libgr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgr_cmd_direct := exec.Command("./binary")
	err = libgr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
