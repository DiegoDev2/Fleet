package main

// YelpTools - Tools that help create and edit Mallard or DocBook documentation
// Homepage: https://gitlab.gnome.org/GNOME/yelp-tools

import (
	"fmt"
	
	"os/exec"
)

func installYelpTools() {
	// Método 1: Descargar y extraer .tar.gz
	yelptools_tar_url := "https://download.gnome.org/sources/yelp-tools/42/yelp-tools-42.1.tar.xz"
	yelptools_cmd_tar := exec.Command("curl", "-L", yelptools_tar_url, "-o", "package.tar.gz")
	err := yelptools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yelptools_zip_url := "https://download.gnome.org/sources/yelp-tools/42/yelp-tools-42.1.tar.xz"
	yelptools_cmd_zip := exec.Command("curl", "-L", yelptools_zip_url, "-o", "package.zip")
	err = yelptools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yelptools_bin_url := "https://download.gnome.org/sources/yelp-tools/42/yelp-tools-42.1.tar.xz"
	yelptools_cmd_bin := exec.Command("curl", "-L", yelptools_bin_url, "-o", "binary.bin")
	err = yelptools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yelptools_src_url := "https://download.gnome.org/sources/yelp-tools/42/yelp-tools-42.1.tar.xz"
	yelptools_cmd_src := exec.Command("curl", "-L", yelptools_src_url, "-o", "source.tar.gz")
	err = yelptools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yelptools_cmd_direct := exec.Command("./binary")
	err = yelptools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
