package main

// Imake - Build automation system written for X11
// Homepage: https://xorg.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installImake() {
	// Método 1: Descargar y extraer .tar.gz
	imake_tar_url := "https://xorg.freedesktop.org/releases/individual/util/imake-1.0.10.tar.xz"
	imake_cmd_tar := exec.Command("curl", "-L", imake_tar_url, "-o", "package.tar.gz")
	err := imake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imake_zip_url := "https://xorg.freedesktop.org/releases/individual/util/imake-1.0.10.tar.xz"
	imake_cmd_zip := exec.Command("curl", "-L", imake_zip_url, "-o", "package.zip")
	err = imake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imake_bin_url := "https://xorg.freedesktop.org/releases/individual/util/imake-1.0.10.tar.xz"
	imake_cmd_bin := exec.Command("curl", "-L", imake_bin_url, "-o", "binary.bin")
	err = imake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imake_src_url := "https://xorg.freedesktop.org/releases/individual/util/imake-1.0.10.tar.xz"
	imake_cmd_src := exec.Command("curl", "-L", imake_src_url, "-o", "source.tar.gz")
	err = imake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imake_cmd_direct := exec.Command("./binary")
	err = imake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: tradcpp")
exec.Command("latte", "install", "tradcpp")
}
