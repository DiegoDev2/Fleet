package main

// Csound - Sound and music computing system
// Homepage: https://csound.com

import (
	"fmt"
	
	"os/exec"
)

func installCsound() {
	// Método 1: Descargar y extraer .tar.gz
	csound_tar_url := "https://github.com/csound/csound.git"
	csound_cmd_tar := exec.Command("curl", "-L", csound_tar_url, "-o", "package.tar.gz")
	err := csound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csound_zip_url := "https://github.com/csound/csound.git"
	csound_cmd_zip := exec.Command("curl", "-L", csound_zip_url, "-o", "package.zip")
	err = csound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csound_bin_url := "https://github.com/csound/csound.git"
	csound_cmd_bin := exec.Command("curl", "-L", csound_bin_url, "-o", "binary.bin")
	err = csound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csound_src_url := "https://github.com/csound/csound.git"
	csound_cmd_src := exec.Command("curl", "-L", csound_src_url, "-o", "source.tar.gz")
	err = csound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csound_cmd_direct := exec.Command("./binary")
	err = csound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asio")
exec.Command("latte", "install", "asio")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: faust")
exec.Command("latte", "install", "faust")
	fmt.Println("Instalando dependencia: fltk")
exec.Command("latte", "install", "fltk")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: liblo")
exec.Command("latte", "install", "liblo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libwebsockets")
exec.Command("latte", "install", "libwebsockets")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: portmidi")
exec.Command("latte", "install", "portmidi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: stk")
exec.Command("latte", "install", "stk")
	fmt.Println("Instalando dependencia: wiiuse")
exec.Command("latte", "install", "wiiuse")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
