package main

// Libxxf86dga - X.Org: XFree86-DGA X extension
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxxf86dga() {
	// Método 1: Descargar y extraer .tar.gz
	libxxf86dga_tar_url := "https://www.x.org/archive/individual/lib/libXxf86dga-1.1.6.tar.xz"
	libxxf86dga_cmd_tar := exec.Command("curl", "-L", libxxf86dga_tar_url, "-o", "package.tar.gz")
	err := libxxf86dga_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxxf86dga_zip_url := "https://www.x.org/archive/individual/lib/libXxf86dga-1.1.6.tar.xz"
	libxxf86dga_cmd_zip := exec.Command("curl", "-L", libxxf86dga_zip_url, "-o", "package.zip")
	err = libxxf86dga_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxxf86dga_bin_url := "https://www.x.org/archive/individual/lib/libXxf86dga-1.1.6.tar.xz"
	libxxf86dga_cmd_bin := exec.Command("curl", "-L", libxxf86dga_bin_url, "-o", "binary.bin")
	err = libxxf86dga_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxxf86dga_src_url := "https://www.x.org/archive/individual/lib/libXxf86dga-1.1.6.tar.xz"
	libxxf86dga_cmd_src := exec.Command("curl", "-L", libxxf86dga_src_url, "-o", "source.tar.gz")
	err = libxxf86dga_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxxf86dga_cmd_direct := exec.Command("./binary")
	err = libxxf86dga_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
}
