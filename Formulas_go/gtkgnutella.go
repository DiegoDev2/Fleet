package main

// GtkGnutella - Share files in a peer-to-peer (P2P) network
// Homepage: https://gtk-gnutella.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installGtkGnutella() {
	// Método 1: Descargar y extraer .tar.gz
	gtkgnutella_tar_url := "https://downloads.sourceforge.net/project/gtk-gnutella/gtk-gnutella/1.2.3/gtk-gnutella-1.2.3.tar.xz"
	gtkgnutella_cmd_tar := exec.Command("curl", "-L", gtkgnutella_tar_url, "-o", "package.tar.gz")
	err := gtkgnutella_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkgnutella_zip_url := "https://downloads.sourceforge.net/project/gtk-gnutella/gtk-gnutella/1.2.3/gtk-gnutella-1.2.3.tar.xz"
	gtkgnutella_cmd_zip := exec.Command("curl", "-L", gtkgnutella_zip_url, "-o", "package.zip")
	err = gtkgnutella_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkgnutella_bin_url := "https://downloads.sourceforge.net/project/gtk-gnutella/gtk-gnutella/1.2.3/gtk-gnutella-1.2.3.tar.xz"
	gtkgnutella_cmd_bin := exec.Command("curl", "-L", gtkgnutella_bin_url, "-o", "binary.bin")
	err = gtkgnutella_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkgnutella_src_url := "https://downloads.sourceforge.net/project/gtk-gnutella/gtk-gnutella/1.2.3/gtk-gnutella-1.2.3.tar.xz"
	gtkgnutella_cmd_src := exec.Command("curl", "-L", gtkgnutella_src_url, "-o", "source.tar.gz")
	err = gtkgnutella_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkgnutella_cmd_direct := exec.Command("./binary")
	err = gtkgnutella_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
