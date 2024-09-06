package main

// OsinfoDbTools - Tools for managing the libosinfo database files
// Homepage: https://libosinfo.org/

import (
	"fmt"
	
	"os/exec"
)

func installOsinfoDbTools() {
	// Método 1: Descargar y extraer .tar.gz
	osinfodbtools_tar_url := "https://releases.pagure.org/libosinfo/osinfo-db-tools-1.11.0.tar.xz"
	osinfodbtools_cmd_tar := exec.Command("curl", "-L", osinfodbtools_tar_url, "-o", "package.tar.gz")
	err := osinfodbtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osinfodbtools_zip_url := "https://releases.pagure.org/libosinfo/osinfo-db-tools-1.11.0.tar.xz"
	osinfodbtools_cmd_zip := exec.Command("curl", "-L", osinfodbtools_zip_url, "-o", "package.zip")
	err = osinfodbtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osinfodbtools_bin_url := "https://releases.pagure.org/libosinfo/osinfo-db-tools-1.11.0.tar.xz"
	osinfodbtools_cmd_bin := exec.Command("curl", "-L", osinfodbtools_bin_url, "-o", "binary.bin")
	err = osinfodbtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osinfodbtools_src_url := "https://releases.pagure.org/libosinfo/osinfo-db-tools-1.11.0.tar.xz"
	osinfodbtools_cmd_src := exec.Command("curl", "-L", osinfodbtools_src_url, "-o", "source.tar.gz")
	err = osinfodbtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osinfodbtools_cmd_direct := exec.Command("./binary")
	err = osinfodbtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
}
