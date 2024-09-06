package main

// WoofDoom - Woof! is a continuation of the Boom/MBF bloodline of Doom source ports
// Homepage: https://github.com/fabiangreffrath/woof

import (
	"fmt"
	
	"os/exec"
)

func installWoofDoom() {
	// Método 1: Descargar y extraer .tar.gz
	woofdoom_tar_url := "https://github.com/fabiangreffrath/woof/archive/refs/tags/woof_14.5.0.tar.gz"
	woofdoom_cmd_tar := exec.Command("curl", "-L", woofdoom_tar_url, "-o", "package.tar.gz")
	err := woofdoom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	woofdoom_zip_url := "https://github.com/fabiangreffrath/woof/archive/refs/tags/woof_14.5.0.zip"
	woofdoom_cmd_zip := exec.Command("curl", "-L", woofdoom_zip_url, "-o", "package.zip")
	err = woofdoom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	woofdoom_bin_url := "https://github.com/fabiangreffrath/woof/archive/refs/tags/woof_14.5.0.bin"
	woofdoom_cmd_bin := exec.Command("curl", "-L", woofdoom_bin_url, "-o", "binary.bin")
	err = woofdoom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	woofdoom_src_url := "https://github.com/fabiangreffrath/woof/archive/refs/tags/woof_14.5.0.src.tar.gz"
	woofdoom_cmd_src := exec.Command("curl", "-L", woofdoom_src_url, "-o", "source.tar.gz")
	err = woofdoom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	woofdoom_cmd_direct := exec.Command("./binary")
	err = woofdoom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libxmp")
	exec.Command("latte", "install", "libxmp").Run()
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_net")
	exec.Command("latte", "install", "sdl2_net").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
