package main

// Libcdr - C++ library to parse the file format of CorelDRAW documents
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libcdr

import (
	"fmt"
	
	"os/exec"
)

func installLibcdr() {
	// Método 1: Descargar y extraer .tar.gz
	libcdr_tar_url := "https://dev-www.libreoffice.org/src/libcdr/libcdr-0.1.7.tar.xz"
	libcdr_cmd_tar := exec.Command("curl", "-L", libcdr_tar_url, "-o", "package.tar.gz")
	err := libcdr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcdr_zip_url := "https://dev-www.libreoffice.org/src/libcdr/libcdr-0.1.7.tar.xz"
	libcdr_cmd_zip := exec.Command("curl", "-L", libcdr_zip_url, "-o", "package.zip")
	err = libcdr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcdr_bin_url := "https://dev-www.libreoffice.org/src/libcdr/libcdr-0.1.7.tar.xz"
	libcdr_cmd_bin := exec.Command("curl", "-L", libcdr_bin_url, "-o", "binary.bin")
	err = libcdr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcdr_src_url := "https://dev-www.libreoffice.org/src/libcdr/libcdr-0.1.7.tar.xz"
	libcdr_cmd_src := exec.Command("curl", "-L", libcdr_src_url, "-o", "source.tar.gz")
	err = libcdr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcdr_cmd_direct := exec.Command("./binary")
	err = libcdr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: librevenge")
exec.Command("latte", "install", "librevenge")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
}
