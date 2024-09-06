package main

// Mp3blaster - Text-based mp3 player
// Homepage: https://mp3blaster.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installMp3blaster() {
	// Método 1: Descargar y extraer .tar.gz
	mp3blaster_tar_url := "https://downloads.sourceforge.net/project/mp3blaster/mp3blaster/mp3blaster-3.2.6/mp3blaster-3.2.6.tar.gz"
	mp3blaster_cmd_tar := exec.Command("curl", "-L", mp3blaster_tar_url, "-o", "package.tar.gz")
	err := mp3blaster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3blaster_zip_url := "https://downloads.sourceforge.net/project/mp3blaster/mp3blaster/mp3blaster-3.2.6/mp3blaster-3.2.6.zip"
	mp3blaster_cmd_zip := exec.Command("curl", "-L", mp3blaster_zip_url, "-o", "package.zip")
	err = mp3blaster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3blaster_bin_url := "https://downloads.sourceforge.net/project/mp3blaster/mp3blaster/mp3blaster-3.2.6/mp3blaster-3.2.6.bin"
	mp3blaster_cmd_bin := exec.Command("curl", "-L", mp3blaster_bin_url, "-o", "binary.bin")
	err = mp3blaster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3blaster_src_url := "https://downloads.sourceforge.net/project/mp3blaster/mp3blaster/mp3blaster-3.2.6/mp3blaster-3.2.6.src.tar.gz"
	mp3blaster_cmd_src := exec.Command("curl", "-L", mp3blaster_src_url, "-o", "source.tar.gz")
	err = mp3blaster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3blaster_cmd_direct := exec.Command("./binary")
	err = mp3blaster_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
}
