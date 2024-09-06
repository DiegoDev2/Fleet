package main

// JpegXl - New file format for still image compression
// Homepage: https://jpeg.org/jpegxl/index.html

import (
	"fmt"
	
	"os/exec"
)

func installJpegXl() {
	// Método 1: Descargar y extraer .tar.gz
	jpegxl_tar_url := "https://github.com/libjxl/libjxl/archive/refs/tags/v0.10.3.tar.gz"
	jpegxl_cmd_tar := exec.Command("curl", "-L", jpegxl_tar_url, "-o", "package.tar.gz")
	err := jpegxl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpegxl_zip_url := "https://github.com/libjxl/libjxl/archive/refs/tags/v0.10.3.zip"
	jpegxl_cmd_zip := exec.Command("curl", "-L", jpegxl_zip_url, "-o", "package.zip")
	err = jpegxl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpegxl_bin_url := "https://github.com/libjxl/libjxl/archive/refs/tags/v0.10.3.bin"
	jpegxl_cmd_bin := exec.Command("curl", "-L", jpegxl_bin_url, "-o", "binary.bin")
	err = jpegxl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpegxl_src_url := "https://github.com/libjxl/libjxl/archive/refs/tags/v0.10.3.src.tar.gz"
	jpegxl_cmd_src := exec.Command("curl", "-L", jpegxl_src_url, "-o", "source.tar.gz")
	err = jpegxl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpegxl_cmd_direct := exec.Command("./binary")
	err = jpegxl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: highway")
exec.Command("latte", "install", "highway")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
