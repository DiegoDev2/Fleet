package main

// Libpagemaker - Imports file format of Aldus/Adobe PageMaker documents
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libpagemaker

import (
	"fmt"
	
	"os/exec"
)

func installLibpagemaker() {
	// Método 1: Descargar y extraer .tar.gz
	libpagemaker_tar_url := "https://dev-www.libreoffice.org/src/libpagemaker/libpagemaker-0.0.4.tar.xz"
	libpagemaker_cmd_tar := exec.Command("curl", "-L", libpagemaker_tar_url, "-o", "package.tar.gz")
	err := libpagemaker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpagemaker_zip_url := "https://dev-www.libreoffice.org/src/libpagemaker/libpagemaker-0.0.4.tar.xz"
	libpagemaker_cmd_zip := exec.Command("curl", "-L", libpagemaker_zip_url, "-o", "package.zip")
	err = libpagemaker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpagemaker_bin_url := "https://dev-www.libreoffice.org/src/libpagemaker/libpagemaker-0.0.4.tar.xz"
	libpagemaker_cmd_bin := exec.Command("curl", "-L", libpagemaker_bin_url, "-o", "binary.bin")
	err = libpagemaker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpagemaker_src_url := "https://dev-www.libreoffice.org/src/libpagemaker/libpagemaker-0.0.4.tar.xz"
	libpagemaker_cmd_src := exec.Command("curl", "-L", libpagemaker_src_url, "-o", "source.tar.gz")
	err = libpagemaker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpagemaker_cmd_direct := exec.Command("./binary")
	err = libpagemaker_cmd_direct.Run()
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
