package main

// Exiftran - Transform digital camera jpegs and their EXIF data
// Homepage: https://www.kraxel.org/blog/linux/fbida/

import (
	"fmt"
	
	"os/exec"
)

func installExiftran() {
	// Método 1: Descargar y extraer .tar.gz
	exiftran_tar_url := "https://www.kraxel.org/releases/fbida/fbida-2.14.tar.gz"
	exiftran_cmd_tar := exec.Command("curl", "-L", exiftran_tar_url, "-o", "package.tar.gz")
	err := exiftran_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exiftran_zip_url := "https://www.kraxel.org/releases/fbida/fbida-2.14.zip"
	exiftran_cmd_zip := exec.Command("curl", "-L", exiftran_zip_url, "-o", "package.zip")
	err = exiftran_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exiftran_bin_url := "https://www.kraxel.org/releases/fbida/fbida-2.14.bin"
	exiftran_cmd_bin := exec.Command("curl", "-L", exiftran_bin_url, "-o", "binary.bin")
	err = exiftran_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exiftran_src_url := "https://www.kraxel.org/releases/fbida/fbida-2.14.src.tar.gz"
	exiftran_cmd_src := exec.Command("curl", "-L", exiftran_src_url, "-o", "source.tar.gz")
	err = exiftran_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exiftran_cmd_direct := exec.Command("./binary")
	err = exiftran_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libexif")
	exec.Command("latte", "install", "libexif").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libepoxy")
	exec.Command("latte", "install", "libepoxy").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: openmotif")
	exec.Command("latte", "install", "openmotif").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
}
