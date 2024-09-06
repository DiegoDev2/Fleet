package main

// OsmGpsMap - GTK+ library to embed OpenStreetMap maps
// Homepage: https://github.com/nzjrs/osm-gps-map

import (
	"fmt"
	
	"os/exec"
)

func installOsmGpsMap() {
	// Método 1: Descargar y extraer .tar.gz
	osmgpsmap_tar_url := "https://github.com/nzjrs/osm-gps-map/releases/download/1.2.0/osm-gps-map-1.2.0.tar.gz"
	osmgpsmap_cmd_tar := exec.Command("curl", "-L", osmgpsmap_tar_url, "-o", "package.tar.gz")
	err := osmgpsmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmgpsmap_zip_url := "https://github.com/nzjrs/osm-gps-map/releases/download/1.2.0/osm-gps-map-1.2.0.zip"
	osmgpsmap_cmd_zip := exec.Command("curl", "-L", osmgpsmap_zip_url, "-o", "package.zip")
	err = osmgpsmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmgpsmap_bin_url := "https://github.com/nzjrs/osm-gps-map/releases/download/1.2.0/osm-gps-map-1.2.0.bin"
	osmgpsmap_cmd_bin := exec.Command("curl", "-L", osmgpsmap_bin_url, "-o", "binary.bin")
	err = osmgpsmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmgpsmap_src_url := "https://github.com/nzjrs/osm-gps-map/releases/download/1.2.0/osm-gps-map-1.2.0.src.tar.gz"
	osmgpsmap_cmd_src := exec.Command("curl", "-L", osmgpsmap_src_url, "-o", "source.tar.gz")
	err = osmgpsmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmgpsmap_cmd_direct := exec.Command("./binary")
	err = osmgpsmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsoup@2")
exec.Command("latte", "install", "libsoup@2")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
