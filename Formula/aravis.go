package main

// Aravis - Vision library for genicam based cameras
// Homepage: https://github.com/AravisProject/aravis

import (
	"fmt"
	
	"os/exec"
)

func installAravis() {
	// Método 1: Descargar y extraer .tar.gz
	aravis_tar_url := "https://github.com/AravisProject/aravis/releases/download/0.8.33/aravis-0.8.33.tar.xz"
	aravis_cmd_tar := exec.Command("curl", "-L", aravis_tar_url, "-o", "package.tar.gz")
	err := aravis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aravis_zip_url := "https://github.com/AravisProject/aravis/releases/download/0.8.33/aravis-0.8.33.tar.xz"
	aravis_cmd_zip := exec.Command("curl", "-L", aravis_zip_url, "-o", "package.zip")
	err = aravis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aravis_bin_url := "https://github.com/AravisProject/aravis/releases/download/0.8.33/aravis-0.8.33.tar.xz"
	aravis_cmd_bin := exec.Command("curl", "-L", aravis_bin_url, "-o", "binary.bin")
	err = aravis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aravis_src_url := "https://github.com/AravisProject/aravis/releases/download/0.8.33/aravis-0.8.33.tar.xz"
	aravis_cmd_src := exec.Command("curl", "-L", aravis_src_url, "-o", "source.tar.gz")
	err = aravis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aravis_cmd_direct := exec.Command("./binary")
	err = aravis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gstreamer")
	exec.Command("latte", "install", "gstreamer").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: libnotify")
	exec.Command("latte", "install", "libnotify").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
