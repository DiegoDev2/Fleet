package main

// Qtads - TADS multimedia interpreter
// Homepage: https://realnc.github.io/qtads/

import (
	"fmt"
	
	"os/exec"
)

func installQtads() {
	// Método 1: Descargar y extraer .tar.gz
	qtads_tar_url := "https://github.com/realnc/qtads/releases/download/v3.4.0/qtads-3.4.0-source.tar.xz"
	qtads_cmd_tar := exec.Command("curl", "-L", qtads_tar_url, "-o", "package.tar.gz")
	err := qtads_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qtads_zip_url := "https://github.com/realnc/qtads/releases/download/v3.4.0/qtads-3.4.0-source.tar.xz"
	qtads_cmd_zip := exec.Command("curl", "-L", qtads_zip_url, "-o", "package.zip")
	err = qtads_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qtads_bin_url := "https://github.com/realnc/qtads/releases/download/v3.4.0/qtads-3.4.0-source.tar.xz"
	qtads_cmd_bin := exec.Command("curl", "-L", qtads_bin_url, "-o", "binary.bin")
	err = qtads_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qtads_src_url := "https://github.com/realnc/qtads/releases/download/v3.4.0/qtads-3.4.0-source.tar.xz"
	qtads_cmd_src := exec.Command("curl", "-L", qtads_src_url, "-o", "source.tar.gz")
	err = qtads_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qtads_cmd_direct := exec.Command("./binary")
	err = qtads_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: mpg123")
	exec.Command("latte", "install", "mpg123").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
