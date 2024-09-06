package main

// Nxengine - Rewrite of Cave Story (Doukutsu Monogatari)
// Homepage: https://nxengine.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installNxengine() {
	// Método 1: Descargar y extraer .tar.gz
	nxengine_tar_url := "https://nxengine.sourceforge.io/dl/nx-src-1006.tar.bz2"
	nxengine_cmd_tar := exec.Command("curl", "-L", nxengine_tar_url, "-o", "package.tar.gz")
	err := nxengine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nxengine_zip_url := "https://nxengine.sourceforge.io/dl/nx-src-1006.tar.bz2"
	nxengine_cmd_zip := exec.Command("curl", "-L", nxengine_zip_url, "-o", "package.zip")
	err = nxengine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nxengine_bin_url := "https://nxengine.sourceforge.io/dl/nx-src-1006.tar.bz2"
	nxengine_cmd_bin := exec.Command("curl", "-L", nxengine_bin_url, "-o", "binary.bin")
	err = nxengine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nxengine_src_url := "https://nxengine.sourceforge.io/dl/nx-src-1006.tar.bz2"
	nxengine_cmd_src := exec.Command("curl", "-L", nxengine_src_url, "-o", "source.tar.gz")
	err = nxengine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nxengine_cmd_direct := exec.Command("./binary")
	err = nxengine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_ttf")
	exec.Command("latte", "install", "sdl_ttf").Run()
}
