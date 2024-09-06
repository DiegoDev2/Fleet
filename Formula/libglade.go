package main

// Libglade - RAD tool to help build GTK+ interfaces
// Homepage: https://glade.gnome.org

import (
	"fmt"
	
	"os/exec"
)

func installLibglade() {
	// Método 1: Descargar y extraer .tar.gz
	libglade_tar_url := "https://download.gnome.org/sources/libglade/2.6/libglade-2.6.4.tar.gz"
	libglade_cmd_tar := exec.Command("curl", "-L", libglade_tar_url, "-o", "package.tar.gz")
	err := libglade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libglade_zip_url := "https://download.gnome.org/sources/libglade/2.6/libglade-2.6.4.zip"
	libglade_cmd_zip := exec.Command("curl", "-L", libglade_zip_url, "-o", "package.zip")
	err = libglade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libglade_bin_url := "https://download.gnome.org/sources/libglade/2.6/libglade-2.6.4.bin"
	libglade_cmd_bin := exec.Command("curl", "-L", libglade_bin_url, "-o", "binary.bin")
	err = libglade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libglade_src_url := "https://download.gnome.org/sources/libglade/2.6/libglade-2.6.4.src.tar.gz"
	libglade_cmd_src := exec.Command("curl", "-L", libglade_src_url, "-o", "source.tar.gz")
	err = libglade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libglade_cmd_direct := exec.Command("./binary")
	err = libglade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
}
