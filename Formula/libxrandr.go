package main

// Libxrandr - X.Org: X Resize, Rotate and Reflection extension library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxrandr() {
	// Método 1: Descargar y extraer .tar.gz
	libxrandr_tar_url := "https://www.x.org/archive/individual/lib/libXrandr-1.5.4.tar.xz"
	libxrandr_cmd_tar := exec.Command("curl", "-L", libxrandr_tar_url, "-o", "package.tar.gz")
	err := libxrandr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxrandr_zip_url := "https://www.x.org/archive/individual/lib/libXrandr-1.5.4.tar.xz"
	libxrandr_cmd_zip := exec.Command("curl", "-L", libxrandr_zip_url, "-o", "package.zip")
	err = libxrandr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxrandr_bin_url := "https://www.x.org/archive/individual/lib/libXrandr-1.5.4.tar.xz"
	libxrandr_cmd_bin := exec.Command("curl", "-L", libxrandr_bin_url, "-o", "binary.bin")
	err = libxrandr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxrandr_src_url := "https://www.x.org/archive/individual/lib/libXrandr-1.5.4.tar.xz"
	libxrandr_cmd_src := exec.Command("curl", "-L", libxrandr_src_url, "-o", "source.tar.gz")
	err = libxrandr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxrandr_cmd_direct := exec.Command("./binary")
	err = libxrandr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
