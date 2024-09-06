package main

// Atop - Advanced system and process monitor for Linux using process events
// Homepage: https://www.atoptool.nl

import (
	"fmt"
	
	"os/exec"
)

func installAtop() {
	// Método 1: Descargar y extraer .tar.gz
	atop_tar_url := "https://github.com/Atoptool/atop/archive/refs/tags/v2.11.0.tar.gz"
	atop_cmd_tar := exec.Command("curl", "-L", atop_tar_url, "-o", "package.tar.gz")
	err := atop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atop_zip_url := "https://github.com/Atoptool/atop/archive/refs/tags/v2.11.0.zip"
	atop_cmd_zip := exec.Command("curl", "-L", atop_zip_url, "-o", "package.zip")
	err = atop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atop_bin_url := "https://github.com/Atoptool/atop/archive/refs/tags/v2.11.0.bin"
	atop_cmd_bin := exec.Command("curl", "-L", atop_bin_url, "-o", "binary.bin")
	err = atop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atop_src_url := "https://github.com/Atoptool/atop/archive/refs/tags/v2.11.0.src.tar.gz"
	atop_cmd_src := exec.Command("curl", "-L", atop_src_url, "-o", "source.tar.gz")
	err = atop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atop_cmd_direct := exec.Command("./binary")
	err = atop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: linux-headers@5.15")
exec.Command("latte", "install", "linux-headers@5.15")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
}
