package main

// Pianobar - Command-line player for https://pandora.com
// Homepage: https://6xq.net/pianobar/

import (
	"fmt"
	
	"os/exec"
)

func installPianobar() {
	// Método 1: Descargar y extraer .tar.gz
	pianobar_tar_url := "https://6xq.net/pianobar/pianobar-2022.04.01.tar.bz2"
	pianobar_cmd_tar := exec.Command("curl", "-L", pianobar_tar_url, "-o", "package.tar.gz")
	err := pianobar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pianobar_zip_url := "https://6xq.net/pianobar/pianobar-2022.04.01.tar.bz2"
	pianobar_cmd_zip := exec.Command("curl", "-L", pianobar_zip_url, "-o", "package.zip")
	err = pianobar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pianobar_bin_url := "https://6xq.net/pianobar/pianobar-2022.04.01.tar.bz2"
	pianobar_cmd_bin := exec.Command("curl", "-L", pianobar_bin_url, "-o", "binary.bin")
	err = pianobar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pianobar_src_url := "https://6xq.net/pianobar/pianobar-2022.04.01.tar.bz2"
	pianobar_cmd_src := exec.Command("curl", "-L", pianobar_src_url, "-o", "source.tar.gz")
	err = pianobar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pianobar_cmd_direct := exec.Command("./binary")
	err = pianobar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libao")
	exec.Command("latte", "install", "libao").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
