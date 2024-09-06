package main

// Fltk - Cross-platform C++ GUI toolkit
// Homepage: https://www.fltk.org/

import (
	"fmt"
	
	"os/exec"
)

func installFltk() {
	// Método 1: Descargar y extraer .tar.gz
	fltk_tar_url := "https://www.fltk.org/pub/fltk/1.3.9/fltk-1.3.9-source.tar.gz"
	fltk_cmd_tar := exec.Command("curl", "-L", fltk_tar_url, "-o", "package.tar.gz")
	err := fltk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fltk_zip_url := "https://www.fltk.org/pub/fltk/1.3.9/fltk-1.3.9-source.zip"
	fltk_cmd_zip := exec.Command("curl", "-L", fltk_zip_url, "-o", "package.zip")
	err = fltk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fltk_bin_url := "https://www.fltk.org/pub/fltk/1.3.9/fltk-1.3.9-source.bin"
	fltk_cmd_bin := exec.Command("curl", "-L", fltk_bin_url, "-o", "binary.bin")
	err = fltk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fltk_src_url := "https://www.fltk.org/pub/fltk/1.3.9/fltk-1.3.9-source.src.tar.gz"
	fltk_cmd_src := exec.Command("curl", "-L", fltk_src_url, "-o", "source.tar.gz")
	err = fltk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fltk_cmd_direct := exec.Command("./binary")
	err = fltk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxft")
exec.Command("latte", "install", "libxft")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
