package main

// Pulseaudio - Sound system for POSIX OSes
// Homepage: https://wiki.freedesktop.org/www/Software/PulseAudio/

import (
	"fmt"
	
	"os/exec"
)

func installPulseaudio() {
	// Método 1: Descargar y extraer .tar.gz
	pulseaudio_tar_url := "https://www.freedesktop.org/software/pulseaudio/releases/pulseaudio-17.0.tar.xz"
	pulseaudio_cmd_tar := exec.Command("curl", "-L", pulseaudio_tar_url, "-o", "package.tar.gz")
	err := pulseaudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pulseaudio_zip_url := "https://www.freedesktop.org/software/pulseaudio/releases/pulseaudio-17.0.tar.xz"
	pulseaudio_cmd_zip := exec.Command("curl", "-L", pulseaudio_zip_url, "-o", "package.zip")
	err = pulseaudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pulseaudio_bin_url := "https://www.freedesktop.org/software/pulseaudio/releases/pulseaudio-17.0.tar.xz"
	pulseaudio_cmd_bin := exec.Command("curl", "-L", pulseaudio_bin_url, "-o", "binary.bin")
	err = pulseaudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pulseaudio_src_url := "https://www.freedesktop.org/software/pulseaudio/releases/pulseaudio-17.0.tar.xz"
	pulseaudio_cmd_src := exec.Command("curl", "-L", pulseaudio_src_url, "-o", "source.tar.gz")
	err = pulseaudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pulseaudio_cmd_direct := exec.Command("./binary")
	err = pulseaudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libsoxr")
	exec.Command("latte", "install", "libsoxr").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: orc")
	exec.Command("latte", "install", "orc").Run()
	fmt.Println("Instalando dependencia: speexdsp")
	exec.Command("latte", "install", "speexdsp").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: perl-xml-parser")
	exec.Command("latte", "install", "perl-xml-parser").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
}
