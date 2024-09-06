package main

// Pioneers - Settlers of Catan clone
// Homepage: https://pio.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPioneers() {
	// Método 1: Descargar y extraer .tar.gz
	pioneers_tar_url := "https://downloads.sourceforge.net/project/pio/Source/pioneers-15.6.tar.gz"
	pioneers_cmd_tar := exec.Command("curl", "-L", pioneers_tar_url, "-o", "package.tar.gz")
	err := pioneers_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pioneers_zip_url := "https://downloads.sourceforge.net/project/pio/Source/pioneers-15.6.zip"
	pioneers_cmd_zip := exec.Command("curl", "-L", pioneers_zip_url, "-o", "package.zip")
	err = pioneers_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pioneers_bin_url := "https://downloads.sourceforge.net/project/pio/Source/pioneers-15.6.bin"
	pioneers_cmd_bin := exec.Command("curl", "-L", pioneers_bin_url, "-o", "binary.bin")
	err = pioneers_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pioneers_src_url := "https://downloads.sourceforge.net/project/pio/Source/pioneers-15.6.src.tar.gz"
	pioneers_cmd_src := exec.Command("curl", "-L", pioneers_src_url, "-o", "source.tar.gz")
	err = pioneers_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pioneers_cmd_direct := exec.Command("./binary")
	err = pioneers_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: itstool")
exec.Command("latte", "install", "itstool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
