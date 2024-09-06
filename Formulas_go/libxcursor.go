package main

// Libxcursor - X.Org: X Window System Cursor management library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxcursor() {
	// Método 1: Descargar y extraer .tar.gz
	libxcursor_tar_url := "https://www.x.org/archive/individual/lib/libXcursor-1.2.2.tar.xz"
	libxcursor_cmd_tar := exec.Command("curl", "-L", libxcursor_tar_url, "-o", "package.tar.gz")
	err := libxcursor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxcursor_zip_url := "https://www.x.org/archive/individual/lib/libXcursor-1.2.2.tar.xz"
	libxcursor_cmd_zip := exec.Command("curl", "-L", libxcursor_zip_url, "-o", "package.zip")
	err = libxcursor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxcursor_bin_url := "https://www.x.org/archive/individual/lib/libXcursor-1.2.2.tar.xz"
	libxcursor_cmd_bin := exec.Command("curl", "-L", libxcursor_bin_url, "-o", "binary.bin")
	err = libxcursor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxcursor_src_url := "https://www.x.org/archive/individual/lib/libXcursor-1.2.2.tar.xz"
	libxcursor_cmd_src := exec.Command("curl", "-L", libxcursor_src_url, "-o", "source.tar.gz")
	err = libxcursor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxcursor_cmd_direct := exec.Command("./binary")
	err = libxcursor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
}
