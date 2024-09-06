package main

// RxvtUnicode - Rxvt fork with Unicode support
// Homepage: http://software.schmorp.de/pkg/rxvt-unicode.html

import (
	"fmt"
	
	"os/exec"
)

func installRxvtUnicode() {
	// Método 1: Descargar y extraer .tar.gz
	rxvtunicode_tar_url := "http://dist.schmorp.de/rxvt-unicode/rxvt-unicode-9.31.tar.bz2"
	rxvtunicode_cmd_tar := exec.Command("curl", "-L", rxvtunicode_tar_url, "-o", "package.tar.gz")
	err := rxvtunicode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rxvtunicode_zip_url := "http://dist.schmorp.de/rxvt-unicode/rxvt-unicode-9.31.tar.bz2"
	rxvtunicode_cmd_zip := exec.Command("curl", "-L", rxvtunicode_zip_url, "-o", "package.zip")
	err = rxvtunicode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rxvtunicode_bin_url := "http://dist.schmorp.de/rxvt-unicode/rxvt-unicode-9.31.tar.bz2"
	rxvtunicode_cmd_bin := exec.Command("curl", "-L", rxvtunicode_bin_url, "-o", "binary.bin")
	err = rxvtunicode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rxvtunicode_src_url := "http://dist.schmorp.de/rxvt-unicode/rxvt-unicode-9.31.tar.bz2"
	rxvtunicode_cmd_src := exec.Command("curl", "-L", rxvtunicode_src_url, "-o", "source.tar.gz")
	err = rxvtunicode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rxvtunicode_cmd_direct := exec.Command("./binary")
	err = rxvtunicode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
}
