package main

// Libmspub - Interpret and import Microsoft Publisher content
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libmspub

import (
	"fmt"
	
	"os/exec"
)

func installLibmspub() {
	// Método 1: Descargar y extraer .tar.gz
	libmspub_tar_url := "https://dev-www.libreoffice.org/src/libmspub/libmspub-0.1.4.tar.xz"
	libmspub_cmd_tar := exec.Command("curl", "-L", libmspub_tar_url, "-o", "package.tar.gz")
	err := libmspub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmspub_zip_url := "https://dev-www.libreoffice.org/src/libmspub/libmspub-0.1.4.tar.xz"
	libmspub_cmd_zip := exec.Command("curl", "-L", libmspub_zip_url, "-o", "package.zip")
	err = libmspub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmspub_bin_url := "https://dev-www.libreoffice.org/src/libmspub/libmspub-0.1.4.tar.xz"
	libmspub_cmd_bin := exec.Command("curl", "-L", libmspub_bin_url, "-o", "binary.bin")
	err = libmspub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmspub_src_url := "https://dev-www.libreoffice.org/src/libmspub/libmspub-0.1.4.tar.xz"
	libmspub_cmd_src := exec.Command("curl", "-L", libmspub_src_url, "-o", "source.tar.gz")
	err = libmspub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmspub_cmd_direct := exec.Command("./binary")
	err = libmspub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libwpg")
	exec.Command("latte", "install", "libwpg").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
	fmt.Println("Instalando dependencia: libwpd")
	exec.Command("latte", "install", "libwpd").Run()
}
