package main

// MonoLibgdiplus - GDI+-compatible API on non-Windows operating systems
// Homepage: https://www.mono-project.com/docs/gui/libgdiplus/

import (
	"fmt"
	
	"os/exec"
)

func installMonoLibgdiplus() {
	// Método 1: Descargar y extraer .tar.gz
	monolibgdiplus_tar_url := "https://download.mono-project.com/sources/libgdiplus/libgdiplus-6.1.tar.gz"
	monolibgdiplus_cmd_tar := exec.Command("curl", "-L", monolibgdiplus_tar_url, "-o", "package.tar.gz")
	err := monolibgdiplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monolibgdiplus_zip_url := "https://download.mono-project.com/sources/libgdiplus/libgdiplus-6.1.zip"
	monolibgdiplus_cmd_zip := exec.Command("curl", "-L", monolibgdiplus_zip_url, "-o", "package.zip")
	err = monolibgdiplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monolibgdiplus_bin_url := "https://download.mono-project.com/sources/libgdiplus/libgdiplus-6.1.bin"
	monolibgdiplus_cmd_bin := exec.Command("curl", "-L", monolibgdiplus_bin_url, "-o", "binary.bin")
	err = monolibgdiplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monolibgdiplus_src_url := "https://download.mono-project.com/sources/libgdiplus/libgdiplus-6.1.src.tar.gz"
	monolibgdiplus_cmd_src := exec.Command("curl", "-L", monolibgdiplus_src_url, "-o", "source.tar.gz")
	err = monolibgdiplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monolibgdiplus_cmd_direct := exec.Command("./binary")
	err = monolibgdiplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
