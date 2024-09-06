package main

// CrispyDoom - Limit-removing enhanced-resolution Doom source port based on Chocolate Doom
// Homepage: https://github.com/fabiangreffrath/crispy-doom

import (
	"fmt"
	
	"os/exec"
)

func installCrispyDoom() {
	// Método 1: Descargar y extraer .tar.gz
	crispydoom_tar_url := "https://github.com/fabiangreffrath/crispy-doom/archive/refs/tags/crispy-doom-7.0.tar.gz"
	crispydoom_cmd_tar := exec.Command("curl", "-L", crispydoom_tar_url, "-o", "package.tar.gz")
	err := crispydoom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crispydoom_zip_url := "https://github.com/fabiangreffrath/crispy-doom/archive/refs/tags/crispy-doom-7.0.zip"
	crispydoom_cmd_zip := exec.Command("curl", "-L", crispydoom_zip_url, "-o", "package.zip")
	err = crispydoom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crispydoom_bin_url := "https://github.com/fabiangreffrath/crispy-doom/archive/refs/tags/crispy-doom-7.0.bin"
	crispydoom_cmd_bin := exec.Command("curl", "-L", crispydoom_bin_url, "-o", "binary.bin")
	err = crispydoom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crispydoom_src_url := "https://github.com/fabiangreffrath/crispy-doom/archive/refs/tags/crispy-doom-7.0.src.tar.gz"
	crispydoom_cmd_src := exec.Command("curl", "-L", crispydoom_src_url, "-o", "source.tar.gz")
	err = crispydoom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crispydoom_cmd_direct := exec.Command("./binary")
	err = crispydoom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: sdl2_net")
	exec.Command("latte", "install", "sdl2_net").Run()
}
