package main

// OpencvAT3 - Open source computer vision library
// Homepage: https://opencv.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpencvAT3() {
	// Método 1: Descargar y extraer .tar.gz
	opencvat3_tar_url := "https://github.com/opencv/opencv/archive/refs/tags/3.4.20.tar.gz"
	opencvat3_cmd_tar := exec.Command("curl", "-L", opencvat3_tar_url, "-o", "package.tar.gz")
	err := opencvat3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencvat3_zip_url := "https://github.com/opencv/opencv/archive/refs/tags/3.4.20.zip"
	opencvat3_cmd_zip := exec.Command("curl", "-L", opencvat3_zip_url, "-o", "package.zip")
	err = opencvat3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencvat3_bin_url := "https://github.com/opencv/opencv/archive/refs/tags/3.4.20.bin"
	opencvat3_cmd_bin := exec.Command("curl", "-L", opencvat3_bin_url, "-o", "binary.bin")
	err = opencvat3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencvat3_src_url := "https://github.com/opencv/opencv/archive/refs/tags/3.4.20.src.tar.gz"
	opencvat3_cmd_src := exec.Command("curl", "-L", opencvat3_src_url, "-o", "source.tar.gz")
	err = opencvat3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencvat3_cmd_direct := exec.Command("./binary")
	err = opencvat3_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: ceres-solver")
exec.Command("latte", "install", "ceres-solver")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: ffmpeg@4")
exec.Command("latte", "install", "ffmpeg@4")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
