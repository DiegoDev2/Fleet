package main

// Gegl - Graph based image processing framework
// Homepage: https://www.gegl.org/

import (
	"fmt"
	
	"os/exec"
)

func installGegl() {
	// Método 1: Descargar y extraer .tar.gz
	gegl_tar_url := "https://download.gimp.org/pub/gegl/0.4/gegl-0.4.48.tar.xz"
	gegl_cmd_tar := exec.Command("curl", "-L", gegl_tar_url, "-o", "package.tar.gz")
	err := gegl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gegl_zip_url := "https://download.gimp.org/pub/gegl/0.4/gegl-0.4.48.tar.xz"
	gegl_cmd_zip := exec.Command("curl", "-L", gegl_zip_url, "-o", "package.zip")
	err = gegl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gegl_bin_url := "https://download.gimp.org/pub/gegl/0.4/gegl-0.4.48.tar.xz"
	gegl_cmd_bin := exec.Command("curl", "-L", gegl_bin_url, "-o", "binary.bin")
	err = gegl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gegl_src_url := "https://download.gimp.org/pub/gegl/0.4/gegl-0.4.48.tar.xz"
	gegl_cmd_src := exec.Command("curl", "-L", gegl_src_url, "-o", "source.tar.gz")
	err = gegl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gegl_cmd_direct := exec.Command("./binary")
	err = gegl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: babl")
	exec.Command("latte", "install", "babl").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
}
