package main

// Dmenu - Dynamic menu for X11
// Homepage: https://tools.suckless.org/dmenu/

import (
	"fmt"
	
	"os/exec"
)

func installDmenu() {
	// Método 1: Descargar y extraer .tar.gz
	dmenu_tar_url := "https://dl.suckless.org/tools/dmenu-5.3.tar.gz"
	dmenu_cmd_tar := exec.Command("curl", "-L", dmenu_tar_url, "-o", "package.tar.gz")
	err := dmenu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmenu_zip_url := "https://dl.suckless.org/tools/dmenu-5.3.zip"
	dmenu_cmd_zip := exec.Command("curl", "-L", dmenu_zip_url, "-o", "package.zip")
	err = dmenu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmenu_bin_url := "https://dl.suckless.org/tools/dmenu-5.3.bin"
	dmenu_cmd_bin := exec.Command("curl", "-L", dmenu_bin_url, "-o", "binary.bin")
	err = dmenu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmenu_src_url := "https://dl.suckless.org/tools/dmenu-5.3.src.tar.gz"
	dmenu_cmd_src := exec.Command("curl", "-L", dmenu_src_url, "-o", "source.tar.gz")
	err = dmenu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmenu_cmd_direct := exec.Command("./binary")
	err = dmenu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxft")
exec.Command("latte", "install", "libxft")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
}
