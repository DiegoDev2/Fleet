package main

// Wxmaxima - Cross platform GUI for Maxima
// Homepage: https://wxmaxima-developers.github.io/wxmaxima/

import (
	"fmt"
	
	"os/exec"
)

func installWxmaxima() {
	// Método 1: Descargar y extraer .tar.gz
	wxmaxima_tar_url := "https://github.com/wxMaxima-developers/wxmaxima/archive/refs/tags/Version-24.08.0.tar.gz"
	wxmaxima_cmd_tar := exec.Command("curl", "-L", wxmaxima_tar_url, "-o", "package.tar.gz")
	err := wxmaxima_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wxmaxima_zip_url := "https://github.com/wxMaxima-developers/wxmaxima/archive/refs/tags/Version-24.08.0.zip"
	wxmaxima_cmd_zip := exec.Command("curl", "-L", wxmaxima_zip_url, "-o", "package.zip")
	err = wxmaxima_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wxmaxima_bin_url := "https://github.com/wxMaxima-developers/wxmaxima/archive/refs/tags/Version-24.08.0.bin"
	wxmaxima_cmd_bin := exec.Command("curl", "-L", wxmaxima_bin_url, "-o", "binary.bin")
	err = wxmaxima_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wxmaxima_src_url := "https://github.com/wxMaxima-developers/wxmaxima/archive/refs/tags/Version-24.08.0.src.tar.gz"
	wxmaxima_cmd_src := exec.Command("curl", "-L", wxmaxima_src_url, "-o", "source.tar.gz")
	err = wxmaxima_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wxmaxima_cmd_direct := exec.Command("./binary")
	err = wxmaxima_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: maxima")
	exec.Command("latte", "install", "maxima").Run()
	fmt.Println("Instalando dependencia: wxwidgets")
	exec.Command("latte", "install", "wxwidgets").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
