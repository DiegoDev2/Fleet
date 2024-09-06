package main

// Libxmlb - Library for querying compressed XML metadata
// Homepage: https://github.com/hughsie/libxmlb

import (
	"fmt"
	
	"os/exec"
)

func installLibxmlb() {
	// Método 1: Descargar y extraer .tar.gz
	libxmlb_tar_url := "https://github.com/hughsie/libxmlb/releases/download/0.3.19/libxmlb-0.3.19.tar.xz"
	libxmlb_cmd_tar := exec.Command("curl", "-L", libxmlb_tar_url, "-o", "package.tar.gz")
	err := libxmlb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmlb_zip_url := "https://github.com/hughsie/libxmlb/releases/download/0.3.19/libxmlb-0.3.19.tar.xz"
	libxmlb_cmd_zip := exec.Command("curl", "-L", libxmlb_zip_url, "-o", "package.zip")
	err = libxmlb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmlb_bin_url := "https://github.com/hughsie/libxmlb/releases/download/0.3.19/libxmlb-0.3.19.tar.xz"
	libxmlb_cmd_bin := exec.Command("curl", "-L", libxmlb_bin_url, "-o", "binary.bin")
	err = libxmlb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmlb_src_url := "https://github.com/hughsie/libxmlb/releases/download/0.3.19/libxmlb-0.3.19.tar.xz"
	libxmlb_cmd_src := exec.Command("curl", "-L", libxmlb_src_url, "-o", "source.tar.gz")
	err = libxmlb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmlb_cmd_direct := exec.Command("./binary")
	err = libxmlb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gi-docgen")
	exec.Command("latte", "install", "gi-docgen").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
