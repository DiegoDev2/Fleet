package main

// Cairomm - Vector graphics library with cross-device output support
// Homepage: https://cairographics.org/cairomm/

import (
	"fmt"
	
	"os/exec"
)

func installCairomm() {
	// Método 1: Descargar y extraer .tar.gz
	cairomm_tar_url := "https://cairographics.org/releases/cairomm-1.18.0.tar.xz"
	cairomm_cmd_tar := exec.Command("curl", "-L", cairomm_tar_url, "-o", "package.tar.gz")
	err := cairomm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cairomm_zip_url := "https://cairographics.org/releases/cairomm-1.18.0.tar.xz"
	cairomm_cmd_zip := exec.Command("curl", "-L", cairomm_zip_url, "-o", "package.zip")
	err = cairomm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cairomm_bin_url := "https://cairographics.org/releases/cairomm-1.18.0.tar.xz"
	cairomm_cmd_bin := exec.Command("curl", "-L", cairomm_bin_url, "-o", "binary.bin")
	err = cairomm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cairomm_src_url := "https://cairographics.org/releases/cairomm-1.18.0.tar.xz"
	cairomm_cmd_src := exec.Command("curl", "-L", cairomm_src_url, "-o", "source.tar.gz")
	err = cairomm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cairomm_cmd_direct := exec.Command("./binary")
	err = cairomm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsigc++")
exec.Command("latte", "install", "libsigc++")
}
