package main

// Gifify - Turn movies into GIFs
// Homepage: https://github.com/jclem/gifify

import (
	"fmt"
	
	"os/exec"
)

func installGifify() {
	// Método 1: Descargar y extraer .tar.gz
	gifify_tar_url := "https://github.com/jclem/gifify/archive/refs/tags/v4.0.tar.gz"
	gifify_cmd_tar := exec.Command("curl", "-L", gifify_tar_url, "-o", "package.tar.gz")
	err := gifify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gifify_zip_url := "https://github.com/jclem/gifify/archive/refs/tags/v4.0.zip"
	gifify_cmd_zip := exec.Command("curl", "-L", gifify_zip_url, "-o", "package.zip")
	err = gifify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gifify_bin_url := "https://github.com/jclem/gifify/archive/refs/tags/v4.0.bin"
	gifify_cmd_bin := exec.Command("curl", "-L", gifify_bin_url, "-o", "binary.bin")
	err = gifify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gifify_src_url := "https://github.com/jclem/gifify/archive/refs/tags/v4.0.src.tar.gz"
	gifify_cmd_src := exec.Command("curl", "-L", gifify_src_url, "-o", "source.tar.gz")
	err = gifify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gifify_cmd_direct := exec.Command("./binary")
	err = gifify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
}
