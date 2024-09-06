package main

// Tgif - Xlib-based interactive 2D drawing tool
// Homepage: https://sourceforge.net/projects/tgif/

import (
	"fmt"
	
	"os/exec"
)

func installTgif() {
	// Método 1: Descargar y extraer .tar.gz
	tgif_tar_url := "https://downloads.sourceforge.net/project/tgif/tgif/4.2.5/tgif-QPL-4.2.5.tar.gz"
	tgif_cmd_tar := exec.Command("curl", "-L", tgif_tar_url, "-o", "package.tar.gz")
	err := tgif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tgif_zip_url := "https://downloads.sourceforge.net/project/tgif/tgif/4.2.5/tgif-QPL-4.2.5.zip"
	tgif_cmd_zip := exec.Command("curl", "-L", tgif_zip_url, "-o", "package.zip")
	err = tgif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tgif_bin_url := "https://downloads.sourceforge.net/project/tgif/tgif/4.2.5/tgif-QPL-4.2.5.bin"
	tgif_cmd_bin := exec.Command("curl", "-L", tgif_bin_url, "-o", "binary.bin")
	err = tgif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tgif_src_url := "https://downloads.sourceforge.net/project/tgif/tgif/4.2.5/tgif-QPL-4.2.5.src.tar.gz"
	tgif_cmd_src := exec.Command("curl", "-L", tgif_src_url, "-o", "source.tar.gz")
	err = tgif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tgif_cmd_direct := exec.Command("./binary")
	err = tgif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libice")
exec.Command("latte", "install", "libice")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
}
