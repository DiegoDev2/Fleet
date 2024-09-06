package main

// Jack - Audio Connection Kit
// Homepage: https://jackaudio.org/

import (
	"fmt"
	
	"os/exec"
)

func installJack() {
	// Método 1: Descargar y extraer .tar.gz
	jack_tar_url := "https://github.com/jackaudio/jack2/archive/refs/tags/v1.9.22.tar.gz"
	jack_cmd_tar := exec.Command("curl", "-L", jack_tar_url, "-o", "package.tar.gz")
	err := jack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jack_zip_url := "https://github.com/jackaudio/jack2/archive/refs/tags/v1.9.22.zip"
	jack_cmd_zip := exec.Command("curl", "-L", jack_zip_url, "-o", "package.zip")
	err = jack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jack_bin_url := "https://github.com/jackaudio/jack2/archive/refs/tags/v1.9.22.bin"
	jack_cmd_bin := exec.Command("curl", "-L", jack_bin_url, "-o", "binary.bin")
	err = jack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jack_src_url := "https://github.com/jackaudio/jack2/archive/refs/tags/v1.9.22.src.tar.gz"
	jack_cmd_src := exec.Command("curl", "-L", jack_src_url, "-o", "source.tar.gz")
	err = jack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jack_cmd_direct := exec.Command("./binary")
	err = jack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: aften")
	exec.Command("latte", "install", "aften").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
