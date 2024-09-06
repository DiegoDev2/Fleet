package main

// AppstreamGlib - Helper library for reading and writing AppStream metadata
// Homepage: https://github.com/hughsie/appstream-glib

import (
	"fmt"
	
	"os/exec"
)

func installAppstreamGlib() {
	// Método 1: Descargar y extraer .tar.gz
	appstreamglib_tar_url := "https://github.com/hughsie/appstream-glib/archive/refs/tags/appstream_glib_0_8_3.tar.gz"
	appstreamglib_cmd_tar := exec.Command("curl", "-L", appstreamglib_tar_url, "-o", "package.tar.gz")
	err := appstreamglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	appstreamglib_zip_url := "https://github.com/hughsie/appstream-glib/archive/refs/tags/appstream_glib_0_8_3.zip"
	appstreamglib_cmd_zip := exec.Command("curl", "-L", appstreamglib_zip_url, "-o", "package.zip")
	err = appstreamglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	appstreamglib_bin_url := "https://github.com/hughsie/appstream-glib/archive/refs/tags/appstream_glib_0_8_3.bin"
	appstreamglib_cmd_bin := exec.Command("curl", "-L", appstreamglib_bin_url, "-o", "binary.bin")
	err = appstreamglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	appstreamglib_src_url := "https://github.com/hughsie/appstream-glib/archive/refs/tags/appstream_glib_0_8_3.src.tar.gz"
	appstreamglib_cmd_src := exec.Command("curl", "-L", appstreamglib_src_url, "-o", "source.tar.gz")
	err = appstreamglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	appstreamglib_cmd_direct := exec.Command("./binary")
	err = appstreamglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
	exec.Command("latte", "install", "docbook").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
