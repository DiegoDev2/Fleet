package main

// Cairo - Vector graphics library with cross-device output support
// Homepage: https://cairographics.org/

import (
	"fmt"
	
	"os/exec"
)

func installCairo() {
	// Método 1: Descargar y extraer .tar.gz
	cairo_tar_url := "https://cairographics.org/releases/cairo-1.18.2.tar.xz"
	cairo_cmd_tar := exec.Command("curl", "-L", cairo_tar_url, "-o", "package.tar.gz")
	err := cairo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cairo_zip_url := "https://cairographics.org/releases/cairo-1.18.2.tar.xz"
	cairo_cmd_zip := exec.Command("curl", "-L", cairo_zip_url, "-o", "package.zip")
	err = cairo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cairo_bin_url := "https://cairographics.org/releases/cairo-1.18.2.tar.xz"
	cairo_cmd_bin := exec.Command("curl", "-L", cairo_bin_url, "-o", "binary.bin")
	err = cairo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cairo_src_url := "https://cairographics.org/releases/cairo-1.18.2.tar.xz"
	cairo_cmd_src := exec.Command("curl", "-L", cairo_src_url, "-o", "source.tar.gz")
	err = cairo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cairo_cmd_direct := exec.Command("./binary")
	err = cairo_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
