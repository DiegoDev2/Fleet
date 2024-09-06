package main

// GtkChtheme - GTK+ 2.0 theme changer GUI
// Homepage: http://plasmasturm.org/code/gtk-chtheme/

import (
	"fmt"
	
	"os/exec"
)

func installGtkChtheme() {
	// Método 1: Descargar y extraer .tar.gz
	gtkchtheme_tar_url := "http://plasmasturm.org/code/gtk-chtheme/gtk-chtheme-0.3.1.tar.bz2"
	gtkchtheme_cmd_tar := exec.Command("curl", "-L", gtkchtheme_tar_url, "-o", "package.tar.gz")
	err := gtkchtheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkchtheme_zip_url := "http://plasmasturm.org/code/gtk-chtheme/gtk-chtheme-0.3.1.tar.bz2"
	gtkchtheme_cmd_zip := exec.Command("curl", "-L", gtkchtheme_zip_url, "-o", "package.zip")
	err = gtkchtheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkchtheme_bin_url := "http://plasmasturm.org/code/gtk-chtheme/gtk-chtheme-0.3.1.tar.bz2"
	gtkchtheme_cmd_bin := exec.Command("curl", "-L", gtkchtheme_bin_url, "-o", "binary.bin")
	err = gtkchtheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkchtheme_src_url := "http://plasmasturm.org/code/gtk-chtheme/gtk-chtheme-0.3.1.tar.bz2"
	gtkchtheme_cmd_src := exec.Command("curl", "-L", gtkchtheme_src_url, "-o", "source.tar.gz")
	err = gtkchtheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkchtheme_cmd_direct := exec.Command("./binary")
	err = gtkchtheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
}
