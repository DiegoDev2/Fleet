package main

// DosboxStaging - Modernized DOSBox soft-fork
// Homepage: https://dosbox-staging.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installDosboxStaging() {
	// Método 1: Descargar y extraer .tar.gz
	dosboxstaging_tar_url := "https://github.com/dosbox-staging/dosbox-staging/archive/refs/tags/v0.81.2.tar.gz"
	dosboxstaging_cmd_tar := exec.Command("curl", "-L", dosboxstaging_tar_url, "-o", "package.tar.gz")
	err := dosboxstaging_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dosboxstaging_zip_url := "https://github.com/dosbox-staging/dosbox-staging/archive/refs/tags/v0.81.2.zip"
	dosboxstaging_cmd_zip := exec.Command("curl", "-L", dosboxstaging_zip_url, "-o", "package.zip")
	err = dosboxstaging_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dosboxstaging_bin_url := "https://github.com/dosbox-staging/dosbox-staging/archive/refs/tags/v0.81.2.bin"
	dosboxstaging_cmd_bin := exec.Command("curl", "-L", dosboxstaging_bin_url, "-o", "binary.bin")
	err = dosboxstaging_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dosboxstaging_src_url := "https://github.com/dosbox-staging/dosbox-staging/archive/refs/tags/v0.81.2.src.tar.gz"
	dosboxstaging_cmd_src := exec.Command("curl", "-L", dosboxstaging_src_url, "-o", "source.tar.gz")
	err = dosboxstaging_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dosboxstaging_cmd_direct := exec.Command("./binary")
	err = dosboxstaging_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: iir1")
exec.Command("latte", "install", "iir1")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libslirp")
exec.Command("latte", "install", "libslirp")
	fmt.Println("Instalando dependencia: mt32emu")
exec.Command("latte", "install", "mt32emu")
	fmt.Println("Instalando dependencia: opusfile")
exec.Command("latte", "install", "opusfile")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
	fmt.Println("Instalando dependencia: sdl2_net")
exec.Command("latte", "install", "sdl2_net")
	fmt.Println("Instalando dependencia: speexdsp")
exec.Command("latte", "install", "speexdsp")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
