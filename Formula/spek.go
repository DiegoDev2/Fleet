package main

// Spek - Acoustic spectrum analyser
// Homepage: https://www.spek.cc

import (
	"fmt"
	
	"os/exec"
)

func installSpek() {
	// Método 1: Descargar y extraer .tar.gz
	spek_tar_url := "https://github.com/alexkay/spek/releases/download/v0.8.5/spek-0.8.5.tar.xz"
	spek_cmd_tar := exec.Command("curl", "-L", spek_tar_url, "-o", "package.tar.gz")
	err := spek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spek_zip_url := "https://github.com/alexkay/spek/releases/download/v0.8.5/spek-0.8.5.tar.xz"
	spek_cmd_zip := exec.Command("curl", "-L", spek_zip_url, "-o", "package.zip")
	err = spek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spek_bin_url := "https://github.com/alexkay/spek/releases/download/v0.8.5/spek-0.8.5.tar.xz"
	spek_cmd_bin := exec.Command("curl", "-L", spek_bin_url, "-o", "binary.bin")
	err = spek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spek_src_url := "https://github.com/alexkay/spek/releases/download/v0.8.5/spek-0.8.5.tar.xz"
	spek_cmd_src := exec.Command("curl", "-L", spek_src_url, "-o", "source.tar.gz")
	err = spek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spek_cmd_direct := exec.Command("./binary")
	err = spek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: wxwidgets")
	exec.Command("latte", "install", "wxwidgets").Run()
}
