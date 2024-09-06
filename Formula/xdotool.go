package main

// Xdotool - Fake keyboard/mouse input and window management for X
// Homepage: https://www.semicomplete.com/projects/xdotool/

import (
	"fmt"
	
	"os/exec"
)

func installXdotool() {
	// Método 1: Descargar y extraer .tar.gz
	xdotool_tar_url := "https://github.com/jordansissel/xdotool/releases/download/v3.20211022.1/xdotool-3.20211022.1.tar.gz"
	xdotool_cmd_tar := exec.Command("curl", "-L", xdotool_tar_url, "-o", "package.tar.gz")
	err := xdotool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xdotool_zip_url := "https://github.com/jordansissel/xdotool/releases/download/v3.20211022.1/xdotool-3.20211022.1.zip"
	xdotool_cmd_zip := exec.Command("curl", "-L", xdotool_zip_url, "-o", "package.zip")
	err = xdotool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xdotool_bin_url := "https://github.com/jordansissel/xdotool/releases/download/v3.20211022.1/xdotool-3.20211022.1.bin"
	xdotool_cmd_bin := exec.Command("curl", "-L", xdotool_bin_url, "-o", "binary.bin")
	err = xdotool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xdotool_src_url := "https://github.com/jordansissel/xdotool/releases/download/v3.20211022.1/xdotool-3.20211022.1.src.tar.gz"
	xdotool_cmd_src := exec.Command("curl", "-L", xdotool_src_url, "-o", "source.tar.gz")
	err = xdotool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xdotool_cmd_direct := exec.Command("./binary")
	err = xdotool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: libxtst")
	exec.Command("latte", "install", "libxtst").Run()
}
