package main

// Libxpm - X.Org: X Pixmap (XPM) image file format library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxpm() {
	// Método 1: Descargar y extraer .tar.gz
	libxpm_tar_url := "https://www.x.org/archive/individual/lib/libXpm-3.5.17.tar.gz"
	libxpm_cmd_tar := exec.Command("curl", "-L", libxpm_tar_url, "-o", "package.tar.gz")
	err := libxpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxpm_zip_url := "https://www.x.org/archive/individual/lib/libXpm-3.5.17.zip"
	libxpm_cmd_zip := exec.Command("curl", "-L", libxpm_zip_url, "-o", "package.zip")
	err = libxpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxpm_bin_url := "https://www.x.org/archive/individual/lib/libXpm-3.5.17.bin"
	libxpm_cmd_bin := exec.Command("curl", "-L", libxpm_bin_url, "-o", "binary.bin")
	err = libxpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxpm_src_url := "https://www.x.org/archive/individual/lib/libXpm-3.5.17.src.tar.gz"
	libxpm_cmd_src := exec.Command("curl", "-L", libxpm_src_url, "-o", "source.tar.gz")
	err = libxpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxpm_cmd_direct := exec.Command("./binary")
	err = libxpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
