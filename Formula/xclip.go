package main

// Xclip - Access X11 clipboards from the command-line
// Homepage: https://github.com/astrand/xclip

import (
	"fmt"
	
	"os/exec"
)

func installXclip() {
	// Método 1: Descargar y extraer .tar.gz
	xclip_tar_url := "https://github.com/astrand/xclip/archive/refs/tags/0.13.tar.gz"
	xclip_cmd_tar := exec.Command("curl", "-L", xclip_tar_url, "-o", "package.tar.gz")
	err := xclip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xclip_zip_url := "https://github.com/astrand/xclip/archive/refs/tags/0.13.zip"
	xclip_cmd_zip := exec.Command("curl", "-L", xclip_zip_url, "-o", "package.zip")
	err = xclip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xclip_bin_url := "https://github.com/astrand/xclip/archive/refs/tags/0.13.bin"
	xclip_cmd_bin := exec.Command("curl", "-L", xclip_bin_url, "-o", "binary.bin")
	err = xclip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xclip_src_url := "https://github.com/astrand/xclip/archive/refs/tags/0.13.src.tar.gz"
	xclip_cmd_src := exec.Command("curl", "-L", xclip_src_url, "-o", "source.tar.gz")
	err = xclip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xclip_cmd_direct := exec.Command("./binary")
	err = xclip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
}
