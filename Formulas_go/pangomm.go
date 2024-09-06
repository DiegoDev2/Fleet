package main

// Pangomm - C++ interface to Pango
// Homepage: https://www.pango.org/

import (
	"fmt"
	
	"os/exec"
)

func installPangomm() {
	// Método 1: Descargar y extraer .tar.gz
	pangomm_tar_url := "https://download.gnome.org/sources/pangomm/2.54/pangomm-2.54.0.tar.xz"
	pangomm_cmd_tar := exec.Command("curl", "-L", pangomm_tar_url, "-o", "package.tar.gz")
	err := pangomm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pangomm_zip_url := "https://download.gnome.org/sources/pangomm/2.54/pangomm-2.54.0.tar.xz"
	pangomm_cmd_zip := exec.Command("curl", "-L", pangomm_zip_url, "-o", "package.zip")
	err = pangomm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pangomm_bin_url := "https://download.gnome.org/sources/pangomm/2.54/pangomm-2.54.0.tar.xz"
	pangomm_cmd_bin := exec.Command("curl", "-L", pangomm_bin_url, "-o", "binary.bin")
	err = pangomm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pangomm_src_url := "https://download.gnome.org/sources/pangomm/2.54/pangomm-2.54.0.tar.xz"
	pangomm_cmd_src := exec.Command("curl", "-L", pangomm_src_url, "-o", "source.tar.gz")
	err = pangomm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pangomm_cmd_direct := exec.Command("./binary")
	err = pangomm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairomm")
exec.Command("latte", "install", "cairomm")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glibmm")
exec.Command("latte", "install", "glibmm")
	fmt.Println("Instalando dependencia: libsigc++")
exec.Command("latte", "install", "libsigc++")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
