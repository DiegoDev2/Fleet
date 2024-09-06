package main

// Opencv - Open source computer vision library
// Homepage: https://opencv.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpencv() {
	// Método 1: Descargar y extraer .tar.gz
	opencv_tar_url := "https://github.com/opencv/opencv/archive/refs/tags/4.10.0.tar.gz"
	opencv_cmd_tar := exec.Command("curl", "-L", opencv_tar_url, "-o", "package.tar.gz")
	err := opencv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencv_zip_url := "https://github.com/opencv/opencv/archive/refs/tags/4.10.0.zip"
	opencv_cmd_zip := exec.Command("curl", "-L", opencv_zip_url, "-o", "package.zip")
	err = opencv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencv_bin_url := "https://github.com/opencv/opencv/archive/refs/tags/4.10.0.bin"
	opencv_cmd_bin := exec.Command("curl", "-L", opencv_bin_url, "-o", "binary.bin")
	err = opencv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencv_src_url := "https://github.com/opencv/opencv/archive/refs/tags/4.10.0.src.tar.gz"
	opencv_cmd_src := exec.Command("curl", "-L", opencv_src_url, "-o", "source.tar.gz")
	err = opencv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencv_cmd_direct := exec.Command("./binary")
	err = opencv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: ceres-solver")
exec.Command("latte", "install", "ceres-solver")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: openvino")
exec.Command("latte", "install", "openvino")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: tesseract")
exec.Command("latte", "install", "tesseract")
	fmt.Println("Instalando dependencia: vtk")
exec.Command("latte", "install", "vtk")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
}
