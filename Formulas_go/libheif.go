package main

// Libheif - ISO/IEC 23008-12:2017 HEIF file format decoder and encoder
// Homepage: https://www.libde265.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibheif() {
	// Método 1: Descargar y extraer .tar.gz
	libheif_tar_url := "https://github.com/strukturag/libheif/releases/download/v1.18.2/libheif-1.18.2.tar.gz"
	libheif_cmd_tar := exec.Command("curl", "-L", libheif_tar_url, "-o", "package.tar.gz")
	err := libheif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libheif_zip_url := "https://github.com/strukturag/libheif/releases/download/v1.18.2/libheif-1.18.2.zip"
	libheif_cmd_zip := exec.Command("curl", "-L", libheif_zip_url, "-o", "package.zip")
	err = libheif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libheif_bin_url := "https://github.com/strukturag/libheif/releases/download/v1.18.2/libheif-1.18.2.bin"
	libheif_cmd_bin := exec.Command("curl", "-L", libheif_bin_url, "-o", "binary.bin")
	err = libheif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libheif_src_url := "https://github.com/strukturag/libheif/releases/download/v1.18.2/libheif-1.18.2.src.tar.gz"
	libheif_cmd_src := exec.Command("curl", "-L", libheif_src_url, "-o", "source.tar.gz")
	err = libheif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libheif_cmd_direct := exec.Command("./binary")
	err = libheif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: aom")
exec.Command("latte", "install", "aom")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libde265")
exec.Command("latte", "install", "libde265")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: shared-mime-info")
exec.Command("latte", "install", "shared-mime-info")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: x265")
exec.Command("latte", "install", "x265")
}
