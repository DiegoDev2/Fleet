package main

// Imlib2 - Image loading and rendering library
// Homepage: https://sourceforge.net/projects/enlightenment/

import (
	"fmt"
	
	"os/exec"
)

func installImlib2() {
	// Método 1: Descargar y extraer .tar.gz
	imlib2_tar_url := "https://downloads.sourceforge.net/project/enlightenment/imlib2-src/1.12.3/imlib2-1.12.3.tar.gz"
	imlib2_cmd_tar := exec.Command("curl", "-L", imlib2_tar_url, "-o", "package.tar.gz")
	err := imlib2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imlib2_zip_url := "https://downloads.sourceforge.net/project/enlightenment/imlib2-src/1.12.3/imlib2-1.12.3.zip"
	imlib2_cmd_zip := exec.Command("curl", "-L", imlib2_zip_url, "-o", "package.zip")
	err = imlib2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imlib2_bin_url := "https://downloads.sourceforge.net/project/enlightenment/imlib2-src/1.12.3/imlib2-1.12.3.bin"
	imlib2_cmd_bin := exec.Command("curl", "-L", imlib2_bin_url, "-o", "binary.bin")
	err = imlib2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imlib2_src_url := "https://downloads.sourceforge.net/project/enlightenment/imlib2-src/1.12.3/imlib2-1.12.3.src.tar.gz"
	imlib2_cmd_src := exec.Command("curl", "-L", imlib2_src_url, "-o", "source.tar.gz")
	err = imlib2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imlib2_cmd_direct := exec.Command("./binary")
	err = imlib2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
