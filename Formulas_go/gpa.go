package main

// Gpa - Graphical user interface for the GnuPG
// Homepage: https://www.gnupg.org/related_software/gpa/

import (
	"fmt"
	
	"os/exec"
)

func installGpa() {
	// Método 1: Descargar y extraer .tar.gz
	gpa_tar_url := "https://gnupg.org/ftp/gcrypt/gpa/gpa-0.10.0.tar.bz2"
	gpa_cmd_tar := exec.Command("curl", "-L", gpa_tar_url, "-o", "package.tar.gz")
	err := gpa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpa_zip_url := "https://gnupg.org/ftp/gcrypt/gpa/gpa-0.10.0.tar.bz2"
	gpa_cmd_zip := exec.Command("curl", "-L", gpa_zip_url, "-o", "package.zip")
	err = gpa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpa_bin_url := "https://gnupg.org/ftp/gcrypt/gpa/gpa-0.10.0.tar.bz2"
	gpa_cmd_bin := exec.Command("curl", "-L", gpa_bin_url, "-o", "binary.bin")
	err = gpa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpa_src_url := "https://gnupg.org/ftp/gcrypt/gpa/gpa-0.10.0.tar.bz2"
	gpa_cmd_src := exec.Command("curl", "-L", gpa_src_url, "-o", "source.tar.gz")
	err = gpa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpa_cmd_direct := exec.Command("./binary")
	err = gpa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: desktop-file-utils")
exec.Command("latte", "install", "desktop-file-utils")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
