package main

// Libjcat - Library for reading Jcat files
// Homepage: https://github.com/hughsie/libjcat

import (
	"fmt"
	
	"os/exec"
)

func installLibjcat() {
	// Método 1: Descargar y extraer .tar.gz
	libjcat_tar_url := "https://github.com/hughsie/libjcat/releases/download/0.2.1/libjcat-0.2.1.tar.xz"
	libjcat_cmd_tar := exec.Command("curl", "-L", libjcat_tar_url, "-o", "package.tar.gz")
	err := libjcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libjcat_zip_url := "https://github.com/hughsie/libjcat/releases/download/0.2.1/libjcat-0.2.1.tar.xz"
	libjcat_cmd_zip := exec.Command("curl", "-L", libjcat_zip_url, "-o", "package.zip")
	err = libjcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libjcat_bin_url := "https://github.com/hughsie/libjcat/releases/download/0.2.1/libjcat-0.2.1.tar.xz"
	libjcat_cmd_bin := exec.Command("curl", "-L", libjcat_bin_url, "-o", "binary.bin")
	err = libjcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libjcat_src_url := "https://github.com/hughsie/libjcat/releases/download/0.2.1/libjcat-0.2.1.tar.xz"
	libjcat_cmd_src := exec.Command("curl", "-L", libjcat_src_url, "-o", "source.tar.gz")
	err = libjcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libjcat_cmd_direct := exec.Command("./binary")
	err = libjcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gi-docgen")
	exec.Command("latte", "install", "gi-docgen").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
