package main

// Openslide - C library to read whole-slide images (a.k.a. virtual slides)
// Homepage: https://openslide.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenslide() {
	// Método 1: Descargar y extraer .tar.gz
	openslide_tar_url := "https://github.com/openslide/openslide/releases/download/v4.0.0/openslide-4.0.0.tar.xz"
	openslide_cmd_tar := exec.Command("curl", "-L", openslide_tar_url, "-o", "package.tar.gz")
	err := openslide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openslide_zip_url := "https://github.com/openslide/openslide/releases/download/v4.0.0/openslide-4.0.0.tar.xz"
	openslide_cmd_zip := exec.Command("curl", "-L", openslide_zip_url, "-o", "package.zip")
	err = openslide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openslide_bin_url := "https://github.com/openslide/openslide/releases/download/v4.0.0/openslide-4.0.0.tar.xz"
	openslide_cmd_bin := exec.Command("curl", "-L", openslide_bin_url, "-o", "binary.bin")
	err = openslide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openslide_src_url := "https://github.com/openslide/openslide/releases/download/v4.0.0/openslide-4.0.0.tar.xz"
	openslide_cmd_src := exec.Command("curl", "-L", openslide_src_url, "-o", "source.tar.gz")
	err = openslide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openslide_cmd_direct := exec.Command("./binary")
	err = openslide_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libdicom")
exec.Command("latte", "install", "libdicom")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
