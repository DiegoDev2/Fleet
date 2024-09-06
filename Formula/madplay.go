package main

// Madplay - MPEG Audio Decoder
// Homepage: https://www.underbit.com/products/mad/

import (
	"fmt"
	
	"os/exec"
)

func installMadplay() {
	// Método 1: Descargar y extraer .tar.gz
	madplay_tar_url := "https://downloads.sourceforge.net/project/mad/madplay/0.15.2b/madplay-0.15.2b.tar.gz"
	madplay_cmd_tar := exec.Command("curl", "-L", madplay_tar_url, "-o", "package.tar.gz")
	err := madplay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	madplay_zip_url := "https://downloads.sourceforge.net/project/mad/madplay/0.15.2b/madplay-0.15.2b.zip"
	madplay_cmd_zip := exec.Command("curl", "-L", madplay_zip_url, "-o", "package.zip")
	err = madplay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	madplay_bin_url := "https://downloads.sourceforge.net/project/mad/madplay/0.15.2b/madplay-0.15.2b.bin"
	madplay_cmd_bin := exec.Command("curl", "-L", madplay_bin_url, "-o", "binary.bin")
	err = madplay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	madplay_src_url := "https://downloads.sourceforge.net/project/mad/madplay/0.15.2b/madplay-0.15.2b.src.tar.gz"
	madplay_cmd_src := exec.Command("curl", "-L", madplay_src_url, "-o", "source.tar.gz")
	err = madplay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	madplay_cmd_direct := exec.Command("./binary")
	err = madplay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libid3tag")
	exec.Command("latte", "install", "libid3tag").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
}
