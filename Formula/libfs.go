package main

// Libfs - X.Org: X Font Service client library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibfs() {
	// Método 1: Descargar y extraer .tar.gz
	libfs_tar_url := "https://xorg.freedesktop.org/archive/individual/lib/libFS-1.0.10.tar.gz"
	libfs_cmd_tar := exec.Command("curl", "-L", libfs_tar_url, "-o", "package.tar.gz")
	err := libfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfs_zip_url := "https://xorg.freedesktop.org/archive/individual/lib/libFS-1.0.10.zip"
	libfs_cmd_zip := exec.Command("curl", "-L", libfs_zip_url, "-o", "package.zip")
	err = libfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfs_bin_url := "https://xorg.freedesktop.org/archive/individual/lib/libFS-1.0.10.bin"
	libfs_cmd_bin := exec.Command("curl", "-L", libfs_bin_url, "-o", "binary.bin")
	err = libfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfs_src_url := "https://xorg.freedesktop.org/archive/individual/lib/libFS-1.0.10.src.tar.gz"
	libfs_cmd_src := exec.Command("curl", "-L", libfs_src_url, "-o", "source.tar.gz")
	err = libfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfs_cmd_direct := exec.Command("./binary")
	err = libfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xtrans")
	exec.Command("latte", "install", "xtrans").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
