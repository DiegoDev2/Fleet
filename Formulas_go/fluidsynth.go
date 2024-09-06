package main

// FluidSynth - Real-time software synthesizer based on the SoundFont 2 specs
// Homepage: https://www.fluidsynth.org

import (
	"fmt"
	
	"os/exec"
)

func installFluidSynth() {
	// Método 1: Descargar y extraer .tar.gz
	fluidsynth_tar_url := "https://github.com/FluidSynth/fluidsynth/archive/refs/tags/v2.3.6.tar.gz"
	fluidsynth_cmd_tar := exec.Command("curl", "-L", fluidsynth_tar_url, "-o", "package.tar.gz")
	err := fluidsynth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fluidsynth_zip_url := "https://github.com/FluidSynth/fluidsynth/archive/refs/tags/v2.3.6.zip"
	fluidsynth_cmd_zip := exec.Command("curl", "-L", fluidsynth_zip_url, "-o", "package.zip")
	err = fluidsynth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fluidsynth_bin_url := "https://github.com/FluidSynth/fluidsynth/archive/refs/tags/v2.3.6.bin"
	fluidsynth_cmd_bin := exec.Command("curl", "-L", fluidsynth_bin_url, "-o", "binary.bin")
	err = fluidsynth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fluidsynth_src_url := "https://github.com/FluidSynth/fluidsynth/archive/refs/tags/v2.3.6.src.tar.gz"
	fluidsynth_cmd_src := exec.Command("curl", "-L", fluidsynth_src_url, "-o", "source.tar.gz")
	err = fluidsynth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fluidsynth_cmd_direct := exec.Command("./binary")
	err = fluidsynth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
