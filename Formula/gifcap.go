package main

// Gifcap - Capture video from an Android device and make a gif
// Homepage: https://github.com/outlook/gifcap

import (
	"fmt"
	
	"os/exec"
)

func installGifcap() {
	// Método 1: Descargar y extraer .tar.gz
	gifcap_tar_url := "https://github.com/outlook/gifcap/archive/refs/tags/1.0.4.tar.gz"
	gifcap_cmd_tar := exec.Command("curl", "-L", gifcap_tar_url, "-o", "package.tar.gz")
	err := gifcap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gifcap_zip_url := "https://github.com/outlook/gifcap/archive/refs/tags/1.0.4.zip"
	gifcap_cmd_zip := exec.Command("curl", "-L", gifcap_zip_url, "-o", "package.zip")
	err = gifcap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gifcap_bin_url := "https://github.com/outlook/gifcap/archive/refs/tags/1.0.4.bin"
	gifcap_cmd_bin := exec.Command("curl", "-L", gifcap_bin_url, "-o", "binary.bin")
	err = gifcap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gifcap_src_url := "https://github.com/outlook/gifcap/archive/refs/tags/1.0.4.src.tar.gz"
	gifcap_cmd_src := exec.Command("curl", "-L", gifcap_src_url, "-o", "source.tar.gz")
	err = gifcap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gifcap_cmd_direct := exec.Command("./binary")
	err = gifcap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
}
