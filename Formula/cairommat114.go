package main

// CairommAT114 - Vector graphics library with cross-device output support
// Homepage: https://cairographics.org/cairomm/

import (
	"fmt"
	
	"os/exec"
)

func installCairommAT114() {
	// Método 1: Descargar y extraer .tar.gz
	cairommat114_tar_url := "https://cairographics.org/releases/cairomm-1.14.5.tar.xz"
	cairommat114_cmd_tar := exec.Command("curl", "-L", cairommat114_tar_url, "-o", "package.tar.gz")
	err := cairommat114_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cairommat114_zip_url := "https://cairographics.org/releases/cairomm-1.14.5.tar.xz"
	cairommat114_cmd_zip := exec.Command("curl", "-L", cairommat114_zip_url, "-o", "package.zip")
	err = cairommat114_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cairommat114_bin_url := "https://cairographics.org/releases/cairomm-1.14.5.tar.xz"
	cairommat114_cmd_bin := exec.Command("curl", "-L", cairommat114_bin_url, "-o", "binary.bin")
	err = cairommat114_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cairommat114_src_url := "https://cairographics.org/releases/cairomm-1.14.5.tar.xz"
	cairommat114_cmd_src := exec.Command("curl", "-L", cairommat114_src_url, "-o", "source.tar.gz")
	err = cairommat114_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cairommat114_cmd_direct := exec.Command("./binary")
	err = cairommat114_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
}
