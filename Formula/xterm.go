package main

// Xterm - Terminal emulator for the X Window System
// Homepage: https://invisible-island.net/xterm/

import (
	"fmt"
	
	"os/exec"
)

func installXterm() {
	// Método 1: Descargar y extraer .tar.gz
	xterm_tar_url := "https://invisible-mirror.net/archives/xterm/xterm-394.tgz"
	xterm_cmd_tar := exec.Command("curl", "-L", xterm_tar_url, "-o", "package.tar.gz")
	err := xterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xterm_zip_url := "https://invisible-mirror.net/archives/xterm/xterm-394.tgz"
	xterm_cmd_zip := exec.Command("curl", "-L", xterm_zip_url, "-o", "package.zip")
	err = xterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xterm_bin_url := "https://invisible-mirror.net/archives/xterm/xterm-394.tgz"
	xterm_cmd_bin := exec.Command("curl", "-L", xterm_bin_url, "-o", "binary.bin")
	err = xterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xterm_src_url := "https://invisible-mirror.net/archives/xterm/xterm-394.tgz"
	xterm_cmd_src := exec.Command("curl", "-L", xterm_src_url, "-o", "source.tar.gz")
	err = xterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xterm_cmd_direct := exec.Command("./binary")
	err = xterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxaw")
	exec.Command("latte", "install", "libxaw").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
}
