package main

// Libvisio - Interpret and import Visio diagrams
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libvisio

import (
	"fmt"
	
	"os/exec"
)

func installLibvisio() {
	// Método 1: Descargar y extraer .tar.gz
	libvisio_tar_url := "https://dev-www.libreoffice.org/src/libvisio/libvisio-0.1.7.tar.xz"
	libvisio_cmd_tar := exec.Command("curl", "-L", libvisio_tar_url, "-o", "package.tar.gz")
	err := libvisio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvisio_zip_url := "https://dev-www.libreoffice.org/src/libvisio/libvisio-0.1.7.tar.xz"
	libvisio_cmd_zip := exec.Command("curl", "-L", libvisio_zip_url, "-o", "package.zip")
	err = libvisio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvisio_bin_url := "https://dev-www.libreoffice.org/src/libvisio/libvisio-0.1.7.tar.xz"
	libvisio_cmd_bin := exec.Command("curl", "-L", libvisio_bin_url, "-o", "binary.bin")
	err = libvisio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvisio_src_url := "https://dev-www.libreoffice.org/src/libvisio/libvisio-0.1.7.tar.xz"
	libvisio_cmd_src := exec.Command("curl", "-L", libvisio_src_url, "-o", "source.tar.gz")
	err = libvisio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvisio_cmd_direct := exec.Command("./binary")
	err = libvisio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
}
