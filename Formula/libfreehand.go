package main

// Libfreehand - Interpret and import Aldus/Macromedia/Adobe FreeHand documents
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libfreehand

import (
	"fmt"
	
	"os/exec"
)

func installLibfreehand() {
	// Método 1: Descargar y extraer .tar.gz
	libfreehand_tar_url := "https://dev-www.libreoffice.org/src/libfreehand/libfreehand-0.1.2.tar.xz"
	libfreehand_cmd_tar := exec.Command("curl", "-L", libfreehand_tar_url, "-o", "package.tar.gz")
	err := libfreehand_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfreehand_zip_url := "https://dev-www.libreoffice.org/src/libfreehand/libfreehand-0.1.2.tar.xz"
	libfreehand_cmd_zip := exec.Command("curl", "-L", libfreehand_zip_url, "-o", "package.zip")
	err = libfreehand_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfreehand_bin_url := "https://dev-www.libreoffice.org/src/libfreehand/libfreehand-0.1.2.tar.xz"
	libfreehand_cmd_bin := exec.Command("curl", "-L", libfreehand_bin_url, "-o", "binary.bin")
	err = libfreehand_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfreehand_src_url := "https://dev-www.libreoffice.org/src/libfreehand/libfreehand-0.1.2.tar.xz"
	libfreehand_cmd_src := exec.Command("curl", "-L", libfreehand_src_url, "-o", "source.tar.gz")
	err = libfreehand_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfreehand_cmd_direct := exec.Command("./binary")
	err = libfreehand_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}
