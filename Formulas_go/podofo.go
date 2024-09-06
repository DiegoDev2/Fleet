package main

// Podofo - Library to work with the PDF file format
// Homepage: https://github.com/podofo/podofo

import (
	"fmt"
	
	"os/exec"
)

func installPodofo() {
	// Método 1: Descargar y extraer .tar.gz
	podofo_tar_url := "https://github.com/podofo/podofo/archive/refs/tags/0.10.3.tar.gz"
	podofo_cmd_tar := exec.Command("curl", "-L", podofo_tar_url, "-o", "package.tar.gz")
	err := podofo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podofo_zip_url := "https://github.com/podofo/podofo/archive/refs/tags/0.10.3.zip"
	podofo_cmd_zip := exec.Command("curl", "-L", podofo_zip_url, "-o", "package.zip")
	err = podofo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podofo_bin_url := "https://github.com/podofo/podofo/archive/refs/tags/0.10.3.bin"
	podofo_cmd_bin := exec.Command("curl", "-L", podofo_bin_url, "-o", "binary.bin")
	err = podofo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podofo_src_url := "https://github.com/podofo/podofo/archive/refs/tags/0.10.3.src.tar.gz"
	podofo_cmd_src := exec.Command("curl", "-L", podofo_src_url, "-o", "source.tar.gz")
	err = podofo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podofo_cmd_direct := exec.Command("./binary")
	err = podofo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
