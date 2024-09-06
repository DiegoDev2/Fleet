package main

// Xmodmap - Modify keymaps and pointer button mappings in X
// Homepage: https://gitlab.freedesktop.org/xorg/app/xmodmap

import (
	"fmt"
	
	"os/exec"
)

func installXmodmap() {
	// Método 1: Descargar y extraer .tar.gz
	xmodmap_tar_url := "https://www.x.org/releases/individual/app/xmodmap-1.0.11.tar.xz"
	xmodmap_cmd_tar := exec.Command("curl", "-L", xmodmap_tar_url, "-o", "package.tar.gz")
	err := xmodmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmodmap_zip_url := "https://www.x.org/releases/individual/app/xmodmap-1.0.11.tar.xz"
	xmodmap_cmd_zip := exec.Command("curl", "-L", xmodmap_zip_url, "-o", "package.zip")
	err = xmodmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmodmap_bin_url := "https://www.x.org/releases/individual/app/xmodmap-1.0.11.tar.xz"
	xmodmap_cmd_bin := exec.Command("curl", "-L", xmodmap_bin_url, "-o", "binary.bin")
	err = xmodmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmodmap_src_url := "https://www.x.org/releases/individual/app/xmodmap-1.0.11.tar.xz"
	xmodmap_cmd_src := exec.Command("curl", "-L", xmodmap_src_url, "-o", "source.tar.gz")
	err = xmodmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmodmap_cmd_direct := exec.Command("./binary")
	err = xmodmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
	fmt.Println("Instalando dependencia: xorg-server")
	exec.Command("latte", "install", "xorg-server").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
