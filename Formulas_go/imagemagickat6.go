package main

// ImagemagickAT6 - Tools and libraries to manipulate images in many formats
// Homepage: https://legacy.imagemagick.org/

import (
	"fmt"
	
	"os/exec"
)

func installImagemagickAT6() {
	// Método 1: Descargar y extraer .tar.gz
	imagemagickat6_tar_url := "https://imagemagick.org/archive/releases/ImageMagick-6.9.13-16.tar.xz"
	imagemagickat6_cmd_tar := exec.Command("curl", "-L", imagemagickat6_tar_url, "-o", "package.tar.gz")
	err := imagemagickat6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imagemagickat6_zip_url := "https://imagemagick.org/archive/releases/ImageMagick-6.9.13-16.tar.xz"
	imagemagickat6_cmd_zip := exec.Command("curl", "-L", imagemagickat6_zip_url, "-o", "package.zip")
	err = imagemagickat6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imagemagickat6_bin_url := "https://imagemagick.org/archive/releases/ImageMagick-6.9.13-16.tar.xz"
	imagemagickat6_cmd_bin := exec.Command("curl", "-L", imagemagickat6_bin_url, "-o", "binary.bin")
	err = imagemagickat6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imagemagickat6_src_url := "https://imagemagick.org/archive/releases/ImageMagick-6.9.13-16.tar.xz"
	imagemagickat6_cmd_src := exec.Command("curl", "-L", imagemagickat6_src_url, "-o", "source.tar.gz")
	err = imagemagickat6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imagemagickat6_cmd_direct := exec.Command("./binary")
	err = imagemagickat6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
