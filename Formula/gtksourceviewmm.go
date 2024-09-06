package main

// Gtksourceviewmm - C++ bindings for gtksourceview
// Homepage: https://gitlab.gnome.org/GNOME/gtksourceviewmm

import (
	"fmt"
	
	"os/exec"
)

func installGtksourceviewmm() {
	// Método 1: Descargar y extraer .tar.gz
	gtksourceviewmm_tar_url := "https://download.gnome.org/sources/gtksourceviewmm/2.10/gtksourceviewmm-2.10.3.tar.xz"
	gtksourceviewmm_cmd_tar := exec.Command("curl", "-L", gtksourceviewmm_tar_url, "-o", "package.tar.gz")
	err := gtksourceviewmm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtksourceviewmm_zip_url := "https://download.gnome.org/sources/gtksourceviewmm/2.10/gtksourceviewmm-2.10.3.tar.xz"
	gtksourceviewmm_cmd_zip := exec.Command("curl", "-L", gtksourceviewmm_zip_url, "-o", "package.zip")
	err = gtksourceviewmm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtksourceviewmm_bin_url := "https://download.gnome.org/sources/gtksourceviewmm/2.10/gtksourceviewmm-2.10.3.tar.xz"
	gtksourceviewmm_cmd_bin := exec.Command("curl", "-L", gtksourceviewmm_bin_url, "-o", "binary.bin")
	err = gtksourceviewmm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtksourceviewmm_src_url := "https://download.gnome.org/sources/gtksourceviewmm/2.10/gtksourceviewmm-2.10.3.tar.xz"
	gtksourceviewmm_cmd_src := exec.Command("curl", "-L", gtksourceviewmm_src_url, "-o", "source.tar.gz")
	err = gtksourceviewmm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtksourceviewmm_cmd_direct := exec.Command("./binary")
	err = gtksourceviewmm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtkmm")
	exec.Command("latte", "install", "gtkmm").Run()
	fmt.Println("Instalando dependencia: gtksourceview")
	exec.Command("latte", "install", "gtksourceview").Run()
}
