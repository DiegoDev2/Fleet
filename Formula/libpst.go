package main

// Libpst - Utilities for the PST file format
// Homepage: https://www.five-ten-sg.com/libpst/

import (
	"fmt"
	
	"os/exec"
)

func installLibpst() {
	// Método 1: Descargar y extraer .tar.gz
	libpst_tar_url := "https://www.five-ten-sg.com/libpst/packages/libpst-0.6.76.tar.gz"
	libpst_cmd_tar := exec.Command("curl", "-L", libpst_tar_url, "-o", "package.tar.gz")
	err := libpst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpst_zip_url := "https://www.five-ten-sg.com/libpst/packages/libpst-0.6.76.zip"
	libpst_cmd_zip := exec.Command("curl", "-L", libpst_zip_url, "-o", "package.zip")
	err = libpst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpst_bin_url := "https://www.five-ten-sg.com/libpst/packages/libpst-0.6.76.bin"
	libpst_cmd_bin := exec.Command("curl", "-L", libpst_bin_url, "-o", "binary.bin")
	err = libpst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpst_src_url := "https://www.five-ten-sg.com/libpst/packages/libpst-0.6.76.src.tar.gz"
	libpst_cmd_src := exec.Command("curl", "-L", libpst_src_url, "-o", "source.tar.gz")
	err = libpst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpst_cmd_direct := exec.Command("./binary")
	err = libpst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libgsf")
	exec.Command("latte", "install", "libgsf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
