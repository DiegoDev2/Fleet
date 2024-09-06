package main

// Fox - Toolkit for developing Graphical User Interfaces easily
// Homepage: http://fox-toolkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installFox() {
	// Método 1: Descargar y extraer .tar.gz
	fox_tar_url := "http://fox-toolkit.org/ftp/fox-1.6.58.tar.gz"
	fox_cmd_tar := exec.Command("curl", "-L", fox_tar_url, "-o", "package.tar.gz")
	err := fox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fox_zip_url := "http://fox-toolkit.org/ftp/fox-1.6.58.zip"
	fox_cmd_zip := exec.Command("curl", "-L", fox_zip_url, "-o", "package.zip")
	err = fox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fox_bin_url := "http://fox-toolkit.org/ftp/fox-1.6.58.bin"
	fox_cmd_bin := exec.Command("curl", "-L", fox_bin_url, "-o", "binary.bin")
	err = fox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fox_src_url := "http://fox-toolkit.org/ftp/fox-1.6.58.src.tar.gz"
	fox_cmd_src := exec.Command("curl", "-L", fox_src_url, "-o", "source.tar.gz")
	err = fox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fox_cmd_direct := exec.Command("./binary")
	err = fox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
	fmt.Println("Instalando dependencia: libxfixes")
	exec.Command("latte", "install", "libxfixes").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
}
