package main

// Wemux - Enhances tmux's to provide multiuser terminal multiplexing
// Homepage: https://github.com/zolrath/wemux

import (
	"fmt"
	
	"os/exec"
)

func installWemux() {
	// Método 1: Descargar y extraer .tar.gz
	wemux_tar_url := "https://github.com/zolrath/wemux/archive/refs/tags/v3.2.0.tar.gz"
	wemux_cmd_tar := exec.Command("curl", "-L", wemux_tar_url, "-o", "package.tar.gz")
	err := wemux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wemux_zip_url := "https://github.com/zolrath/wemux/archive/refs/tags/v3.2.0.zip"
	wemux_cmd_zip := exec.Command("curl", "-L", wemux_zip_url, "-o", "package.zip")
	err = wemux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wemux_bin_url := "https://github.com/zolrath/wemux/archive/refs/tags/v3.2.0.bin"
	wemux_cmd_bin := exec.Command("curl", "-L", wemux_bin_url, "-o", "binary.bin")
	err = wemux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wemux_src_url := "https://github.com/zolrath/wemux/archive/refs/tags/v3.2.0.src.tar.gz"
	wemux_cmd_src := exec.Command("curl", "-L", wemux_src_url, "-o", "source.tar.gz")
	err = wemux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wemux_cmd_direct := exec.Command("./binary")
	err = wemux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tmux")
exec.Command("latte", "install", "tmux")
}
