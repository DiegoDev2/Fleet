package main

// Timg - Terminal image and video viewer
// Homepage: https://timg.sh/

import (
	"fmt"
	
	"os/exec"
)

func installTimg() {
	// Método 1: Descargar y extraer .tar.gz
	timg_tar_url := "https://github.com/hzeller/timg/archive/refs/tags/v1.6.0.tar.gz"
	timg_cmd_tar := exec.Command("curl", "-L", timg_tar_url, "-o", "package.tar.gz")
	err := timg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	timg_zip_url := "https://github.com/hzeller/timg/archive/refs/tags/v1.6.0.zip"
	timg_cmd_zip := exec.Command("curl", "-L", timg_zip_url, "-o", "package.zip")
	err = timg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	timg_bin_url := "https://github.com/hzeller/timg/archive/refs/tags/v1.6.0.bin"
	timg_cmd_bin := exec.Command("curl", "-L", timg_bin_url, "-o", "binary.bin")
	err = timg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	timg_src_url := "https://github.com/hzeller/timg/archive/refs/tags/v1.6.0.src.tar.gz"
	timg_cmd_src := exec.Command("curl", "-L", timg_src_url, "-o", "source.tar.gz")
	err = timg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	timg_cmd_direct := exec.Command("./binary")
	err = timg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: graphicsmagick")
exec.Command("latte", "install", "graphicsmagick")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libdeflate")
exec.Command("latte", "install", "libdeflate")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: libsixel")
exec.Command("latte", "install", "libsixel")
	fmt.Println("Instalando dependencia: openslide")
exec.Command("latte", "install", "openslide")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
