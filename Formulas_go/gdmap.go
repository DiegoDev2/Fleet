package main

// Gdmap - Tool to inspect the used space of folders
// Homepage: https://sourceforge.net/projects/gdmap/

import (
	"fmt"
	
	"os/exec"
)

func installGdmap() {
	// Método 1: Descargar y extraer .tar.gz
	gdmap_tar_url := "https://downloads.sourceforge.net/project/gdmap/gdmap/0.8.1/gdmap-0.8.1.tar.gz"
	gdmap_cmd_tar := exec.Command("curl", "-L", gdmap_tar_url, "-o", "package.tar.gz")
	err := gdmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdmap_zip_url := "https://downloads.sourceforge.net/project/gdmap/gdmap/0.8.1/gdmap-0.8.1.zip"
	gdmap_cmd_zip := exec.Command("curl", "-L", gdmap_zip_url, "-o", "package.zip")
	err = gdmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdmap_bin_url := "https://downloads.sourceforge.net/project/gdmap/gdmap/0.8.1/gdmap-0.8.1.bin"
	gdmap_cmd_bin := exec.Command("curl", "-L", gdmap_bin_url, "-o", "binary.bin")
	err = gdmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdmap_src_url := "https://downloads.sourceforge.net/project/gdmap/gdmap/0.8.1/gdmap-0.8.1.src.tar.gz"
	gdmap_cmd_src := exec.Command("curl", "-L", gdmap_src_url, "-o", "source.tar.gz")
	err = gdmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdmap_cmd_direct := exec.Command("./binary")
	err = gdmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
}
