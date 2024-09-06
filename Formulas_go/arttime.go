package main

// Arttime - Clock, timer, time manager and ASCII+ text-art viewer for the terminal
// Homepage: https://github.com/poetaman/arttime

import (
	"fmt"
	
	"os/exec"
)

func installArttime() {
	// Método 1: Descargar y extraer .tar.gz
	arttime_tar_url := "https://github.com/poetaman/arttime/archive/refs/tags/v2.3.4.tar.gz"
	arttime_cmd_tar := exec.Command("curl", "-L", arttime_tar_url, "-o", "package.tar.gz")
	err := arttime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arttime_zip_url := "https://github.com/poetaman/arttime/archive/refs/tags/v2.3.4.zip"
	arttime_cmd_zip := exec.Command("curl", "-L", arttime_zip_url, "-o", "package.zip")
	err = arttime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arttime_bin_url := "https://github.com/poetaman/arttime/archive/refs/tags/v2.3.4.bin"
	arttime_cmd_bin := exec.Command("curl", "-L", arttime_bin_url, "-o", "binary.bin")
	err = arttime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arttime_src_url := "https://github.com/poetaman/arttime/archive/refs/tags/v2.3.4.src.tar.gz"
	arttime_cmd_src := exec.Command("curl", "-L", arttime_src_url, "-o", "source.tar.gz")
	err = arttime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arttime_cmd_direct := exec.Command("./binary")
	err = arttime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fzf")
exec.Command("latte", "install", "fzf")
	fmt.Println("Instalando dependencia: diffutils")
exec.Command("latte", "install", "diffutils")
	fmt.Println("Instalando dependencia: less")
exec.Command("latte", "install", "less")
	fmt.Println("Instalando dependencia: libnotify")
exec.Command("latte", "install", "libnotify")
	fmt.Println("Instalando dependencia: vorbis-tools")
exec.Command("latte", "install", "vorbis-tools")
	fmt.Println("Instalando dependencia: zsh")
exec.Command("latte", "install", "zsh")
}
