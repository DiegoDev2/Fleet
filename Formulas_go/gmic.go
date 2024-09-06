package main

// Gmic - Full-Featured Open-Source Framework for Image Processing
// Homepage: https://gmic.eu/

import (
	"fmt"
	
	"os/exec"
)

func installGmic() {
	// Método 1: Descargar y extraer .tar.gz
	gmic_tar_url := "https://gmic.eu/files/source/gmic_3.4.2.tar.gz"
	gmic_cmd_tar := exec.Command("curl", "-L", gmic_tar_url, "-o", "package.tar.gz")
	err := gmic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmic_zip_url := "https://gmic.eu/files/source/gmic_3.4.2.zip"
	gmic_cmd_zip := exec.Command("curl", "-L", gmic_zip_url, "-o", "package.zip")
	err = gmic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmic_bin_url := "https://gmic.eu/files/source/gmic_3.4.2.bin"
	gmic_cmd_bin := exec.Command("curl", "-L", gmic_bin_url, "-o", "binary.bin")
	err = gmic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmic_src_url := "https://gmic.eu/files/source/gmic_3.4.2.src.tar.gz"
	gmic_cmd_src := exec.Command("curl", "-L", gmic_src_url, "-o", "source.tar.gz")
	err = gmic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmic_cmd_direct := exec.Command("./binary")
	err = gmic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cimg")
exec.Command("latte", "install", "cimg")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
