package main

// Libxfixes - X.Org: Header files for the XFIXES extension
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxfixes() {
	// Método 1: Descargar y extraer .tar.gz
	libxfixes_tar_url := "https://www.x.org/archive/individual/lib/libXfixes-6.0.1.tar.xz"
	libxfixes_cmd_tar := exec.Command("curl", "-L", libxfixes_tar_url, "-o", "package.tar.gz")
	err := libxfixes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxfixes_zip_url := "https://www.x.org/archive/individual/lib/libXfixes-6.0.1.tar.xz"
	libxfixes_cmd_zip := exec.Command("curl", "-L", libxfixes_zip_url, "-o", "package.zip")
	err = libxfixes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxfixes_bin_url := "https://www.x.org/archive/individual/lib/libXfixes-6.0.1.tar.xz"
	libxfixes_cmd_bin := exec.Command("curl", "-L", libxfixes_bin_url, "-o", "binary.bin")
	err = libxfixes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxfixes_src_url := "https://www.x.org/archive/individual/lib/libXfixes-6.0.1.tar.xz"
	libxfixes_cmd_src := exec.Command("curl", "-L", libxfixes_src_url, "-o", "source.tar.gz")
	err = libxfixes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxfixes_cmd_direct := exec.Command("./binary")
	err = libxfixes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
