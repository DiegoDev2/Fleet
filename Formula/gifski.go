package main

// Gifski - Highest-quality GIF encoder based on pngquant
// Homepage: https://gif.ski/

import (
	"fmt"
	
	"os/exec"
)

func installGifski() {
	// Método 1: Descargar y extraer .tar.gz
	gifski_tar_url := "https://github.com/ImageOptim/gifski/archive/refs/tags/1.32.0.tar.gz"
	gifski_cmd_tar := exec.Command("curl", "-L", gifski_tar_url, "-o", "package.tar.gz")
	err := gifski_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gifski_zip_url := "https://github.com/ImageOptim/gifski/archive/refs/tags/1.32.0.zip"
	gifski_cmd_zip := exec.Command("curl", "-L", gifski_zip_url, "-o", "package.zip")
	err = gifski_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gifski_bin_url := "https://github.com/ImageOptim/gifski/archive/refs/tags/1.32.0.bin"
	gifski_cmd_bin := exec.Command("curl", "-L", gifski_bin_url, "-o", "binary.bin")
	err = gifski_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gifski_src_url := "https://github.com/ImageOptim/gifski/archive/refs/tags/1.32.0.src.tar.gz"
	gifski_cmd_src := exec.Command("curl", "-L", gifski_src_url, "-o", "source.tar.gz")
	err = gifski_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gifski_cmd_direct := exec.Command("./binary")
	err = gifski_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: ffmpeg@6")
	exec.Command("latte", "install", "ffmpeg@6").Run()
}
