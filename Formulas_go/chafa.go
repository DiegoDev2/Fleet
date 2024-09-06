package main

// Chafa - Versatile and fast Unicode/ASCII/ANSI graphics renderer
// Homepage: https://hpjansson.org/chafa/

import (
	"fmt"
	
	"os/exec"
)

func installChafa() {
	// Método 1: Descargar y extraer .tar.gz
	chafa_tar_url := "https://hpjansson.org/chafa/releases/chafa-1.14.2.tar.xz"
	chafa_cmd_tar := exec.Command("curl", "-L", chafa_tar_url, "-o", "package.tar.gz")
	err := chafa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chafa_zip_url := "https://hpjansson.org/chafa/releases/chafa-1.14.2.tar.xz"
	chafa_cmd_zip := exec.Command("curl", "-L", chafa_zip_url, "-o", "package.zip")
	err = chafa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chafa_bin_url := "https://hpjansson.org/chafa/releases/chafa-1.14.2.tar.xz"
	chafa_cmd_bin := exec.Command("curl", "-L", chafa_bin_url, "-o", "binary.bin")
	err = chafa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chafa_src_url := "https://hpjansson.org/chafa/releases/chafa-1.14.2.tar.xz"
	chafa_cmd_src := exec.Command("curl", "-L", chafa_src_url, "-o", "source.tar.gz")
	err = chafa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chafa_cmd_direct := exec.Command("./binary")
	err = chafa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
