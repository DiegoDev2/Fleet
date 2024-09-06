package main

// Libxaw3d - X.Org: 3D Athena widget set based on the Xt library
// Homepage: https://www.x.org

import (
	"fmt"
	
	"os/exec"
)

func installLibxaw3d() {
	// Método 1: Descargar y extraer .tar.gz
	libxaw3d_tar_url := "https://xorg.freedesktop.org/archive/individual/lib/libXaw3d-1.6.6.tar.gz"
	libxaw3d_cmd_tar := exec.Command("curl", "-L", libxaw3d_tar_url, "-o", "package.tar.gz")
	err := libxaw3d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxaw3d_zip_url := "https://xorg.freedesktop.org/archive/individual/lib/libXaw3d-1.6.6.zip"
	libxaw3d_cmd_zip := exec.Command("curl", "-L", libxaw3d_zip_url, "-o", "package.zip")
	err = libxaw3d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxaw3d_bin_url := "https://xorg.freedesktop.org/archive/individual/lib/libXaw3d-1.6.6.bin"
	libxaw3d_cmd_bin := exec.Command("curl", "-L", libxaw3d_bin_url, "-o", "binary.bin")
	err = libxaw3d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxaw3d_src_url := "https://xorg.freedesktop.org/archive/individual/lib/libXaw3d-1.6.6.src.tar.gz"
	libxaw3d_cmd_src := exec.Command("curl", "-L", libxaw3d_src_url, "-o", "source.tar.gz")
	err = libxaw3d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxaw3d_cmd_direct := exec.Command("./binary")
	err = libxaw3d_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
	fmt.Println("Instalando dependencia: libxpm")
exec.Command("latte", "install", "libxpm")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
}
