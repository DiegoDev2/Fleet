package main

// Icoutils - Create and extract MS Windows icons and cursors
// Homepage: https://www.nongnu.org/icoutils/

import (
	"fmt"
	
	"os/exec"
)

func installIcoutils() {
	// Método 1: Descargar y extraer .tar.gz
	icoutils_tar_url := "https://savannah.nongnu.org/download/icoutils/icoutils-0.32.3.tar.bz2"
	icoutils_cmd_tar := exec.Command("curl", "-L", icoutils_tar_url, "-o", "package.tar.gz")
	err := icoutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icoutils_zip_url := "https://savannah.nongnu.org/download/icoutils/icoutils-0.32.3.tar.bz2"
	icoutils_cmd_zip := exec.Command("curl", "-L", icoutils_zip_url, "-o", "package.zip")
	err = icoutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icoutils_bin_url := "https://savannah.nongnu.org/download/icoutils/icoutils-0.32.3.tar.bz2"
	icoutils_cmd_bin := exec.Command("curl", "-L", icoutils_bin_url, "-o", "binary.bin")
	err = icoutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icoutils_src_url := "https://savannah.nongnu.org/download/icoutils/icoutils-0.32.3.tar.bz2"
	icoutils_cmd_src := exec.Command("curl", "-L", icoutils_src_url, "-o", "source.tar.gz")
	err = icoutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icoutils_cmd_direct := exec.Command("./binary")
	err = icoutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
