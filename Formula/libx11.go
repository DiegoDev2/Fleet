package main

// Libx11 - X.Org: Core X11 protocol client library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibx11() {
	// Método 1: Descargar y extraer .tar.gz
	libx11_tar_url := "https://www.x.org/archive/individual/lib/libX11-1.8.10.tar.gz"
	libx11_cmd_tar := exec.Command("curl", "-L", libx11_tar_url, "-o", "package.tar.gz")
	err := libx11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libx11_zip_url := "https://www.x.org/archive/individual/lib/libX11-1.8.10.zip"
	libx11_cmd_zip := exec.Command("curl", "-L", libx11_zip_url, "-o", "package.zip")
	err = libx11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libx11_bin_url := "https://www.x.org/archive/individual/lib/libX11-1.8.10.bin"
	libx11_cmd_bin := exec.Command("curl", "-L", libx11_bin_url, "-o", "binary.bin")
	err = libx11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libx11_src_url := "https://www.x.org/archive/individual/lib/libX11-1.8.10.src.tar.gz"
	libx11_cmd_src := exec.Command("curl", "-L", libx11_src_url, "-o", "source.tar.gz")
	err = libx11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libx11_cmd_direct := exec.Command("./binary")
	err = libx11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: xtrans")
	exec.Command("latte", "install", "xtrans").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
