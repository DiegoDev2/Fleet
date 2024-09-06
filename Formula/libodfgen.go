package main

// Libodfgen - ODF export library for projects using librevenge
// Homepage: https://sourceforge.net/p/libwpd/wiki/libodfgen/

import (
	"fmt"
	
	"os/exec"
)

func installLibodfgen() {
	// Método 1: Descargar y extraer .tar.gz
	libodfgen_tar_url := "https://dev-www.libreoffice.org/src/libodfgen-0.1.8.tar.xz"
	libodfgen_cmd_tar := exec.Command("curl", "-L", libodfgen_tar_url, "-o", "package.tar.gz")
	err := libodfgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libodfgen_zip_url := "https://dev-www.libreoffice.org/src/libodfgen-0.1.8.tar.xz"
	libodfgen_cmd_zip := exec.Command("curl", "-L", libodfgen_zip_url, "-o", "package.zip")
	err = libodfgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libodfgen_bin_url := "https://dev-www.libreoffice.org/src/libodfgen-0.1.8.tar.xz"
	libodfgen_cmd_bin := exec.Command("curl", "-L", libodfgen_bin_url, "-o", "binary.bin")
	err = libodfgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libodfgen_src_url := "https://dev-www.libreoffice.org/src/libodfgen-0.1.8.tar.xz"
	libodfgen_cmd_src := exec.Command("curl", "-L", libodfgen_src_url, "-o", "source.tar.gz")
	err = libodfgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libodfgen_cmd_direct := exec.Command("./binary")
	err = libodfgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libetonyek")
	exec.Command("latte", "install", "libetonyek").Run()
	fmt.Println("Instalando dependencia: libwpg")
	exec.Command("latte", "install", "libwpg").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
	fmt.Println("Instalando dependencia: libwpd")
	exec.Command("latte", "install", "libwpd").Run()
}
