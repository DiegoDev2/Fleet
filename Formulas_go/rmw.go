package main

// Rmw - Trashcan/recycle bin utility for the command-line
// Homepage: https://theimpossibleastronaut.github.io/rmw-website/

import (
	"fmt"
	
	"os/exec"
)

func installRmw() {
	// Método 1: Descargar y extraer .tar.gz
	rmw_tar_url := "https://github.com/theimpossibleastronaut/rmw/releases/download/v0.9.2/rmw-0.9.2.tar.xz"
	rmw_cmd_tar := exec.Command("curl", "-L", rmw_tar_url, "-o", "package.tar.gz")
	err := rmw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rmw_zip_url := "https://github.com/theimpossibleastronaut/rmw/releases/download/v0.9.2/rmw-0.9.2.tar.xz"
	rmw_cmd_zip := exec.Command("curl", "-L", rmw_zip_url, "-o", "package.zip")
	err = rmw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rmw_bin_url := "https://github.com/theimpossibleastronaut/rmw/releases/download/v0.9.2/rmw-0.9.2.tar.xz"
	rmw_cmd_bin := exec.Command("curl", "-L", rmw_bin_url, "-o", "binary.bin")
	err = rmw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rmw_src_url := "https://github.com/theimpossibleastronaut/rmw/releases/download/v0.9.2/rmw-0.9.2.tar.xz"
	rmw_cmd_src := exec.Command("curl", "-L", rmw_src_url, "-o", "source.tar.gz")
	err = rmw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rmw_cmd_direct := exec.Command("./binary")
	err = rmw_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: canfigger")
exec.Command("latte", "install", "canfigger")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
