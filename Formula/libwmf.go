package main

// Libwmf - Library for converting WMF (Window Metafile Format) files
// Homepage: https://wvware.sourceforge.net/libwmf.html

import (
	"fmt"
	
	"os/exec"
)

func installLibwmf() {
	// Método 1: Descargar y extraer .tar.gz
	libwmf_tar_url := "https://downloads.sourceforge.net/project/wvware/libwmf/0.2.8.4/libwmf-0.2.8.4.tar.gz"
	libwmf_cmd_tar := exec.Command("curl", "-L", libwmf_tar_url, "-o", "package.tar.gz")
	err := libwmf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwmf_zip_url := "https://downloads.sourceforge.net/project/wvware/libwmf/0.2.8.4/libwmf-0.2.8.4.zip"
	libwmf_cmd_zip := exec.Command("curl", "-L", libwmf_zip_url, "-o", "package.zip")
	err = libwmf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwmf_bin_url := "https://downloads.sourceforge.net/project/wvware/libwmf/0.2.8.4/libwmf-0.2.8.4.bin"
	libwmf_cmd_bin := exec.Command("curl", "-L", libwmf_bin_url, "-o", "binary.bin")
	err = libwmf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwmf_src_url := "https://downloads.sourceforge.net/project/wvware/libwmf/0.2.8.4/libwmf-0.2.8.4.src.tar.gz"
	libwmf_cmd_src := exec.Command("curl", "-L", libwmf_src_url, "-o", "source.tar.gz")
	err = libwmf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwmf_cmd_direct := exec.Command("./binary")
	err = libwmf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
