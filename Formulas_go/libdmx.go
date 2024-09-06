package main

// Libdmx - X.Org: X Window System DMX (Distributed Multihead X) extension library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibdmx() {
	// Método 1: Descargar y extraer .tar.gz
	libdmx_tar_url := "https://www.x.org/archive/individual/lib/libdmx-1.1.5.tar.xz"
	libdmx_cmd_tar := exec.Command("curl", "-L", libdmx_tar_url, "-o", "package.tar.gz")
	err := libdmx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdmx_zip_url := "https://www.x.org/archive/individual/lib/libdmx-1.1.5.tar.xz"
	libdmx_cmd_zip := exec.Command("curl", "-L", libdmx_zip_url, "-o", "package.zip")
	err = libdmx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdmx_bin_url := "https://www.x.org/archive/individual/lib/libdmx-1.1.5.tar.xz"
	libdmx_cmd_bin := exec.Command("curl", "-L", libdmx_bin_url, "-o", "binary.bin")
	err = libdmx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdmx_src_url := "https://www.x.org/archive/individual/lib/libdmx-1.1.5.tar.xz"
	libdmx_cmd_src := exec.Command("curl", "-L", libdmx_src_url, "-o", "source.tar.gz")
	err = libdmx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdmx_cmd_direct := exec.Command("./binary")
	err = libdmx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
