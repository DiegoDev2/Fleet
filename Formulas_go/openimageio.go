package main

// Openimageio - Library for reading, processing and writing images
// Homepage: https://openimageio.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installOpenimageio() {
	// Método 1: Descargar y extraer .tar.gz
	openimageio_tar_url := "https://github.com/AcademySoftwareFoundation/OpenImageIO/archive/refs/tags/v2.5.15.0.tar.gz"
	openimageio_cmd_tar := exec.Command("curl", "-L", openimageio_tar_url, "-o", "package.tar.gz")
	err := openimageio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openimageio_zip_url := "https://github.com/AcademySoftwareFoundation/OpenImageIO/archive/refs/tags/v2.5.15.0.zip"
	openimageio_cmd_zip := exec.Command("curl", "-L", openimageio_zip_url, "-o", "package.zip")
	err = openimageio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openimageio_bin_url := "https://github.com/AcademySoftwareFoundation/OpenImageIO/archive/refs/tags/v2.5.15.0.bin"
	openimageio_cmd_bin := exec.Command("curl", "-L", openimageio_bin_url, "-o", "binary.bin")
	err = openimageio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openimageio_src_url := "https://github.com/AcademySoftwareFoundation/OpenImageIO/archive/refs/tags/v2.5.15.0.src.tar.gz"
	openimageio_cmd_src := exec.Command("curl", "-L", openimageio_src_url, "-o", "source.tar.gz")
	err = openimageio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openimageio_cmd_direct := exec.Command("./binary")
	err = openimageio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: boost-python3")
exec.Command("latte", "install", "boost-python3")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libraw")
exec.Command("latte", "install", "libraw")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: opencolorio")
exec.Command("latte", "install", "opencolorio")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: pugixml")
exec.Command("latte", "install", "pugixml")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
