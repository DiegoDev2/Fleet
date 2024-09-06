package main

// Babl - Dynamic, any-to-any, pixel format translation library
// Homepage: https://www.gegl.org/babl/

import (
	"fmt"
	
	"os/exec"
)

func installBabl() {
	// Método 1: Descargar y extraer .tar.gz
	babl_tar_url := "https://download.gimp.org/pub/babl/0.1/babl-0.1.108.tar.xz"
	babl_cmd_tar := exec.Command("curl", "-L", babl_tar_url, "-o", "package.tar.gz")
	err := babl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	babl_zip_url := "https://download.gimp.org/pub/babl/0.1/babl-0.1.108.tar.xz"
	babl_cmd_zip := exec.Command("curl", "-L", babl_zip_url, "-o", "package.zip")
	err = babl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	babl_bin_url := "https://download.gimp.org/pub/babl/0.1/babl-0.1.108.tar.xz"
	babl_cmd_bin := exec.Command("curl", "-L", babl_bin_url, "-o", "binary.bin")
	err = babl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	babl_src_url := "https://download.gimp.org/pub/babl/0.1/babl-0.1.108.tar.xz"
	babl_cmd_src := exec.Command("curl", "-L", babl_src_url, "-o", "source.tar.gz")
	err = babl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	babl_cmd_direct := exec.Command("./binary")
	err = babl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
}
