package main

// Libraw - Library for reading RAW files from digital photo cameras
// Homepage: https://www.libraw.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibraw() {
	// Método 1: Descargar y extraer .tar.gz
	libraw_tar_url := "https://www.libraw.org/data/LibRaw-0.21.2.tar.gz"
	libraw_cmd_tar := exec.Command("curl", "-L", libraw_tar_url, "-o", "package.tar.gz")
	err := libraw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libraw_zip_url := "https://www.libraw.org/data/LibRaw-0.21.2.zip"
	libraw_cmd_zip := exec.Command("curl", "-L", libraw_zip_url, "-o", "package.zip")
	err = libraw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libraw_bin_url := "https://www.libraw.org/data/LibRaw-0.21.2.bin"
	libraw_cmd_bin := exec.Command("curl", "-L", libraw_bin_url, "-o", "binary.bin")
	err = libraw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libraw_src_url := "https://www.libraw.org/data/LibRaw-0.21.2.src.tar.gz"
	libraw_cmd_src := exec.Command("curl", "-L", libraw_src_url, "-o", "source.tar.gz")
	err = libraw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libraw_cmd_direct := exec.Command("./binary")
	err = libraw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
