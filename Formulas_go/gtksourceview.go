package main

// Gtksourceview - Text view with syntax, undo/redo, and text marks
// Homepage: https://projects.gnome.org/gtksourceview/

import (
	"fmt"
	
	"os/exec"
)

func installGtksourceview() {
	// Método 1: Descargar y extraer .tar.gz
	gtksourceview_tar_url := "https://download.gnome.org/sources/gtksourceview/2.10/gtksourceview-2.10.5.tar.gz"
	gtksourceview_cmd_tar := exec.Command("curl", "-L", gtksourceview_tar_url, "-o", "package.tar.gz")
	err := gtksourceview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtksourceview_zip_url := "https://download.gnome.org/sources/gtksourceview/2.10/gtksourceview-2.10.5.zip"
	gtksourceview_cmd_zip := exec.Command("curl", "-L", gtksourceview_zip_url, "-o", "package.zip")
	err = gtksourceview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtksourceview_bin_url := "https://download.gnome.org/sources/gtksourceview/2.10/gtksourceview-2.10.5.bin"
	gtksourceview_cmd_bin := exec.Command("curl", "-L", gtksourceview_bin_url, "-o", "binary.bin")
	err = gtksourceview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtksourceview_src_url := "https://download.gnome.org/sources/gtksourceview/2.10/gtksourceview-2.10.5.src.tar.gz"
	gtksourceview_cmd_src := exec.Command("curl", "-L", gtksourceview_src_url, "-o", "source.tar.gz")
	err = gtksourceview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtksourceview_cmd_direct := exec.Command("./binary")
	err = gtksourceview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
