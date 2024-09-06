package main

// PangommAT246 - C++ interface to Pango
// Homepage: https://www.pango.org/

import (
	"fmt"
	
	"os/exec"
)

func installPangommAT246() {
	// Método 1: Descargar y extraer .tar.gz
	pangommat246_tar_url := "https://download.gnome.org/sources/pangomm/2.46/pangomm-2.46.4.tar.xz"
	pangommat246_cmd_tar := exec.Command("curl", "-L", pangommat246_tar_url, "-o", "package.tar.gz")
	err := pangommat246_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pangommat246_zip_url := "https://download.gnome.org/sources/pangomm/2.46/pangomm-2.46.4.tar.xz"
	pangommat246_cmd_zip := exec.Command("curl", "-L", pangommat246_zip_url, "-o", "package.zip")
	err = pangommat246_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pangommat246_bin_url := "https://download.gnome.org/sources/pangomm/2.46/pangomm-2.46.4.tar.xz"
	pangommat246_cmd_bin := exec.Command("curl", "-L", pangommat246_bin_url, "-o", "binary.bin")
	err = pangommat246_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pangommat246_src_url := "https://download.gnome.org/sources/pangomm/2.46/pangomm-2.46.4.tar.xz"
	pangommat246_cmd_src := exec.Command("curl", "-L", pangommat246_src_url, "-o", "source.tar.gz")
	err = pangommat246_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pangommat246_cmd_direct := exec.Command("./binary")
	err = pangommat246_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairomm@1.14")
exec.Command("latte", "install", "cairomm@1.14")
	fmt.Println("Instalando dependencia: glibmm@2.66")
exec.Command("latte", "install", "glibmm@2.66")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
