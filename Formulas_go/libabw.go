package main

// Libabw - Library for parsing AbiWord documents
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libabw

import (
	"fmt"
	
	"os/exec"
)

func installLibabw() {
	// Método 1: Descargar y extraer .tar.gz
	libabw_tar_url := "https://dev-www.libreoffice.org/src/libabw/libabw-0.1.3.tar.xz"
	libabw_cmd_tar := exec.Command("curl", "-L", libabw_tar_url, "-o", "package.tar.gz")
	err := libabw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libabw_zip_url := "https://dev-www.libreoffice.org/src/libabw/libabw-0.1.3.tar.xz"
	libabw_cmd_zip := exec.Command("curl", "-L", libabw_zip_url, "-o", "package.zip")
	err = libabw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libabw_bin_url := "https://dev-www.libreoffice.org/src/libabw/libabw-0.1.3.tar.xz"
	libabw_cmd_bin := exec.Command("curl", "-L", libabw_bin_url, "-o", "binary.bin")
	err = libabw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libabw_src_url := "https://dev-www.libreoffice.org/src/libabw/libabw-0.1.3.tar.xz"
	libabw_cmd_src := exec.Command("curl", "-L", libabw_src_url, "-o", "source.tar.gz")
	err = libabw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libabw_cmd_direct := exec.Command("./binary")
	err = libabw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: librevenge")
exec.Command("latte", "install", "librevenge")
}
