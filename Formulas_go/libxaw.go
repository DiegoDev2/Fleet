package main

// Libxaw - X.Org: X Athena Widget Set
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxaw() {
	// Método 1: Descargar y extraer .tar.gz
	libxaw_tar_url := "https://www.x.org/archive/individual/lib/libXaw-1.0.16.tar.xz"
	libxaw_cmd_tar := exec.Command("curl", "-L", libxaw_tar_url, "-o", "package.tar.gz")
	err := libxaw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxaw_zip_url := "https://www.x.org/archive/individual/lib/libXaw-1.0.16.tar.xz"
	libxaw_cmd_zip := exec.Command("curl", "-L", libxaw_zip_url, "-o", "package.zip")
	err = libxaw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxaw_bin_url := "https://www.x.org/archive/individual/lib/libXaw-1.0.16.tar.xz"
	libxaw_cmd_bin := exec.Command("curl", "-L", libxaw_bin_url, "-o", "binary.bin")
	err = libxaw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxaw_src_url := "https://www.x.org/archive/individual/lib/libXaw-1.0.16.tar.xz"
	libxaw_cmd_src := exec.Command("curl", "-L", libxaw_src_url, "-o", "source.tar.gz")
	err = libxaw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxaw_cmd_direct := exec.Command("./binary")
	err = libxaw_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
	fmt.Println("Instalando dependencia: libxpm")
exec.Command("latte", "install", "libxpm")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
}
