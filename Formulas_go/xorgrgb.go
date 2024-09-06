package main

// Xorgrgb - X.Org: color names database
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installXorgrgb() {
	// Método 1: Descargar y extraer .tar.gz
	xorgrgb_tar_url := "https://xorg.freedesktop.org/archive/individual/app/rgb-1.1.0.tar.gz"
	xorgrgb_cmd_tar := exec.Command("curl", "-L", xorgrgb_tar_url, "-o", "package.tar.gz")
	err := xorgrgb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xorgrgb_zip_url := "https://xorg.freedesktop.org/archive/individual/app/rgb-1.1.0.zip"
	xorgrgb_cmd_zip := exec.Command("curl", "-L", xorgrgb_zip_url, "-o", "package.zip")
	err = xorgrgb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xorgrgb_bin_url := "https://xorg.freedesktop.org/archive/individual/app/rgb-1.1.0.bin"
	xorgrgb_cmd_bin := exec.Command("curl", "-L", xorgrgb_bin_url, "-o", "binary.bin")
	err = xorgrgb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xorgrgb_src_url := "https://xorg.freedesktop.org/archive/individual/app/rgb-1.1.0.src.tar.gz"
	xorgrgb_cmd_src := exec.Command("curl", "-L", xorgrgb_src_url, "-o", "source.tar.gz")
	err = xorgrgb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xorgrgb_cmd_direct := exec.Command("./binary")
	err = xorgrgb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
