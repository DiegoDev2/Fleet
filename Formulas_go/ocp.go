package main

// Ocp - UNIX port of the Open Cubic Player
// Homepage: https://stian.cubic.org/project-ocp.php

import (
	"fmt"
	
	"os/exec"
)

func installOcp() {
	// Método 1: Descargar y extraer .tar.gz
	ocp_tar_url := "https://stian.cubic.org/ocp/ocp-0.2.109.tar.xz"
	ocp_cmd_tar := exec.Command("curl", "-L", ocp_tar_url, "-o", "package.tar.gz")
	err := ocp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocp_zip_url := "https://stian.cubic.org/ocp/ocp-0.2.109.tar.xz"
	ocp_cmd_zip := exec.Command("curl", "-L", ocp_zip_url, "-o", "package.zip")
	err = ocp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocp_bin_url := "https://stian.cubic.org/ocp/ocp-0.2.109.tar.xz"
	ocp_cmd_bin := exec.Command("curl", "-L", ocp_bin_url, "-o", "binary.bin")
	err = ocp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocp_src_url := "https://stian.cubic.org/ocp/ocp-0.2.109.tar.xz"
	ocp_cmd_src := exec.Command("curl", "-L", ocp_src_url, "-o", "source.tar.gz")
	err = ocp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocp_cmd_direct := exec.Command("./binary")
	err = ocp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xa")
exec.Command("latte", "install", "xa")
	fmt.Println("Instalando dependencia: ancient")
exec.Command("latte", "install", "ancient")
	fmt.Println("Instalando dependencia: cjson")
exec.Command("latte", "install", "cjson")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: game-music-emu")
exec.Command("latte", "install", "game-music-emu")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libdiscid")
exec.Command("latte", "install", "libdiscid")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mad")
exec.Command("latte", "install", "mad")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
