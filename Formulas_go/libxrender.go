package main

// Libxrender - X.Org: Library for the Render Extension to the X11 protocol
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxrender() {
	// Método 1: Descargar y extraer .tar.gz
	libxrender_tar_url := "https://www.x.org/archive/individual/lib/libXrender-0.9.11.tar.gz"
	libxrender_cmd_tar := exec.Command("curl", "-L", libxrender_tar_url, "-o", "package.tar.gz")
	err := libxrender_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxrender_zip_url := "https://www.x.org/archive/individual/lib/libXrender-0.9.11.zip"
	libxrender_cmd_zip := exec.Command("curl", "-L", libxrender_zip_url, "-o", "package.zip")
	err = libxrender_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxrender_bin_url := "https://www.x.org/archive/individual/lib/libXrender-0.9.11.bin"
	libxrender_cmd_bin := exec.Command("curl", "-L", libxrender_bin_url, "-o", "binary.bin")
	err = libxrender_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxrender_src_url := "https://www.x.org/archive/individual/lib/libXrender-0.9.11.src.tar.gz"
	libxrender_cmd_src := exec.Command("curl", "-L", libxrender_src_url, "-o", "source.tar.gz")
	err = libxrender_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxrender_cmd_direct := exec.Command("./binary")
	err = libxrender_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
