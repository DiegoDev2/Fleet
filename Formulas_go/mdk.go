package main

// Mdk - GNU MIX development kit
// Homepage: https://www.gnu.org/software/mdk/mdk.html

import (
	"fmt"
	
	"os/exec"
)

func installMdk() {
	// Método 1: Descargar y extraer .tar.gz
	mdk_tar_url := "https://ftp.gnu.org/gnu/mdk/v1.3.0/mdk-1.3.0.tar.gz"
	mdk_cmd_tar := exec.Command("curl", "-L", mdk_tar_url, "-o", "package.tar.gz")
	err := mdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdk_zip_url := "https://ftp.gnu.org/gnu/mdk/v1.3.0/mdk-1.3.0.zip"
	mdk_cmd_zip := exec.Command("curl", "-L", mdk_zip_url, "-o", "package.zip")
	err = mdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdk_bin_url := "https://ftp.gnu.org/gnu/mdk/v1.3.0/mdk-1.3.0.bin"
	mdk_cmd_bin := exec.Command("curl", "-L", mdk_bin_url, "-o", "binary.bin")
	err = mdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdk_src_url := "https://ftp.gnu.org/gnu/mdk/v1.3.0/mdk-1.3.0.src.tar.gz"
	mdk_cmd_src := exec.Command("curl", "-L", mdk_src_url, "-o", "source.tar.gz")
	err = mdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdk_cmd_direct := exec.Command("./binary")
	err = mdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: guile")
exec.Command("latte", "install", "guile")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
