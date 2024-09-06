package main

// Snapcast - Synchronous multiroom audio player
// Homepage: https://github.com/badaix/snapcast

import (
	"fmt"
	
	"os/exec"
)

func installSnapcast() {
	// Método 1: Descargar y extraer .tar.gz
	snapcast_tar_url := "https://github.com/badaix/snapcast/archive/refs/tags/v0.29.0.tar.gz"
	snapcast_cmd_tar := exec.Command("curl", "-L", snapcast_tar_url, "-o", "package.tar.gz")
	err := snapcast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snapcast_zip_url := "https://github.com/badaix/snapcast/archive/refs/tags/v0.29.0.zip"
	snapcast_cmd_zip := exec.Command("curl", "-L", snapcast_zip_url, "-o", "package.zip")
	err = snapcast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snapcast_bin_url := "https://github.com/badaix/snapcast/archive/refs/tags/v0.29.0.bin"
	snapcast_cmd_bin := exec.Command("curl", "-L", snapcast_bin_url, "-o", "binary.bin")
	err = snapcast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snapcast_src_url := "https://github.com/badaix/snapcast/archive/refs/tags/v0.29.0.src.tar.gz"
	snapcast_cmd_src := exec.Command("curl", "-L", snapcast_src_url, "-o", "source.tar.gz")
	err = snapcast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snapcast_cmd_direct := exec.Command("./binary")
	err = snapcast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libsoxr")
	exec.Command("latte", "install", "libsoxr").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: avahi")
	exec.Command("latte", "install", "avahi").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
}
