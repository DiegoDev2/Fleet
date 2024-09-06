package main

// Wxwidgets - Cross-platform C++ GUI toolkit
// Homepage: https://www.wxwidgets.org

import (
	"fmt"
	
	"os/exec"
)

func installWxwidgets() {
	// Método 1: Descargar y extraer .tar.gz
	wxwidgets_tar_url := "https://github.com/wxWidgets/wxWidgets/releases/download/v3.2.5/wxWidgets-3.2.5.tar.bz2"
	wxwidgets_cmd_tar := exec.Command("curl", "-L", wxwidgets_tar_url, "-o", "package.tar.gz")
	err := wxwidgets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wxwidgets_zip_url := "https://github.com/wxWidgets/wxWidgets/releases/download/v3.2.5/wxWidgets-3.2.5.tar.bz2"
	wxwidgets_cmd_zip := exec.Command("curl", "-L", wxwidgets_zip_url, "-o", "package.zip")
	err = wxwidgets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wxwidgets_bin_url := "https://github.com/wxWidgets/wxWidgets/releases/download/v3.2.5/wxWidgets-3.2.5.tar.bz2"
	wxwidgets_cmd_bin := exec.Command("curl", "-L", wxwidgets_bin_url, "-o", "binary.bin")
	err = wxwidgets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wxwidgets_src_url := "https://github.com/wxWidgets/wxWidgets/releases/download/v3.2.5/wxWidgets-3.2.5.tar.bz2"
	wxwidgets_cmd_src := exec.Command("curl", "-L", wxwidgets_src_url, "-o", "source.tar.gz")
	err = wxwidgets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wxwidgets_cmd_direct := exec.Command("./binary")
	err = wxwidgets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
