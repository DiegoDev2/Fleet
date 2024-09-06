package main

// DmtxUtils - Read and write data matrix barcodes
// Homepage: https://github.com/dmtx/dmtx-utils

import (
	"fmt"
	
	"os/exec"
)

func installDmtxUtils() {
	// Método 1: Descargar y extraer .tar.gz
	dmtxutils_tar_url := "https://github.com/dmtx/dmtx-utils/archive/refs/tags/v0.7.6.tar.gz"
	dmtxutils_cmd_tar := exec.Command("curl", "-L", dmtxutils_tar_url, "-o", "package.tar.gz")
	err := dmtxutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmtxutils_zip_url := "https://github.com/dmtx/dmtx-utils/archive/refs/tags/v0.7.6.zip"
	dmtxutils_cmd_zip := exec.Command("curl", "-L", dmtxutils_zip_url, "-o", "package.zip")
	err = dmtxutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmtxutils_bin_url := "https://github.com/dmtx/dmtx-utils/archive/refs/tags/v0.7.6.bin"
	dmtxutils_cmd_bin := exec.Command("curl", "-L", dmtxutils_bin_url, "-o", "binary.bin")
	err = dmtxutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmtxutils_src_url := "https://github.com/dmtx/dmtx-utils/archive/refs/tags/v0.7.6.src.tar.gz"
	dmtxutils_cmd_src := exec.Command("curl", "-L", dmtxutils_src_url, "-o", "source.tar.gz")
	err = dmtxutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmtxutils_cmd_direct := exec.Command("./binary")
	err = dmtxutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
	fmt.Println("Instalando dependencia: libdmtx")
	exec.Command("latte", "install", "libdmtx").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: liblqr")
	exec.Command("latte", "install", "liblqr").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}
