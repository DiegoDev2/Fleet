package main

// Gitg - GNOME GUI client to view git repositories
// Homepage: https://wiki.gnome.org/Apps/Gitg

import (
	"fmt"
	
	"os/exec"
)

func installGitg() {
	// Método 1: Descargar y extraer .tar.gz
	gitg_tar_url := "https://download.gnome.org/sources/gitg/44/gitg-44.tar.xz"
	gitg_cmd_tar := exec.Command("curl", "-L", gitg_tar_url, "-o", "package.tar.gz")
	err := gitg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitg_zip_url := "https://download.gnome.org/sources/gitg/44/gitg-44.tar.xz"
	gitg_cmd_zip := exec.Command("curl", "-L", gitg_zip_url, "-o", "package.zip")
	err = gitg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitg_bin_url := "https://download.gnome.org/sources/gitg/44/gitg-44.tar.xz"
	gitg_cmd_bin := exec.Command("curl", "-L", gitg_bin_url, "-o", "binary.bin")
	err = gitg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitg_src_url := "https://download.gnome.org/sources/gitg/44/gitg-44.tar.xz"
	gitg_cmd_src := exec.Command("curl", "-L", gitg_src_url, "-o", "source.tar.gz")
	err = gitg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitg_cmd_direct := exec.Command("./binary")
	err = gitg_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: gspell")
	exec.Command("latte", "install", "gspell").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: gtksourceview4")
	exec.Command("latte", "install", "gtksourceview4").Run()
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
	exec.Command("latte", "install", "hicolor-icon-theme").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libdazzle")
	exec.Command("latte", "install", "libdazzle").Run()
	fmt.Println("Instalando dependencia: libgee")
	exec.Command("latte", "install", "libgee").Run()
	fmt.Println("Instalando dependencia: libgit2-glib")
	exec.Command("latte", "install", "libgit2-glib").Run()
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
	fmt.Println("Instalando dependencia: libhandy")
	exec.Command("latte", "install", "libhandy").Run()
	fmt.Println("Instalando dependencia: libpeas@1")
	exec.Command("latte", "install", "libpeas@1").Run()
	fmt.Println("Instalando dependencia: libsecret")
	exec.Command("latte", "install", "libsecret").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
