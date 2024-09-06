package main

// Navidrome - Modern Music Server and Streamer compatible with Subsonic/Airsonic
// Homepage: https://www.navidrome.org

import (
	"fmt"
	
	"os/exec"
)

func installNavidrome() {
	// Método 1: Descargar y extraer .tar.gz
	navidrome_tar_url := "https://github.com/navidrome/navidrome/archive/refs/tags/v0.52.5.tar.gz"
	navidrome_cmd_tar := exec.Command("curl", "-L", navidrome_tar_url, "-o", "package.tar.gz")
	err := navidrome_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	navidrome_zip_url := "https://github.com/navidrome/navidrome/archive/refs/tags/v0.52.5.zip"
	navidrome_cmd_zip := exec.Command("curl", "-L", navidrome_zip_url, "-o", "package.zip")
	err = navidrome_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	navidrome_bin_url := "https://github.com/navidrome/navidrome/archive/refs/tags/v0.52.5.bin"
	navidrome_cmd_bin := exec.Command("curl", "-L", navidrome_bin_url, "-o", "binary.bin")
	err = navidrome_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	navidrome_src_url := "https://github.com/navidrome/navidrome/archive/refs/tags/v0.52.5.src.tar.gz"
	navidrome_cmd_src := exec.Command("curl", "-L", navidrome_src_url, "-o", "source.tar.gz")
	err = navidrome_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	navidrome_cmd_direct := exec.Command("./binary")
	err = navidrome_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
}
