package main

// Scummvm - Graphic adventure game interpreter
// Homepage: https://www.scummvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installScummvm() {
	// Método 1: Descargar y extraer .tar.gz
	scummvm_tar_url := "https://downloads.scummvm.org/frs/scummvm/2.8.1/scummvm-2.8.1.tar.xz"
	scummvm_cmd_tar := exec.Command("curl", "-L", scummvm_tar_url, "-o", "package.tar.gz")
	err := scummvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scummvm_zip_url := "https://downloads.scummvm.org/frs/scummvm/2.8.1/scummvm-2.8.1.tar.xz"
	scummvm_cmd_zip := exec.Command("curl", "-L", scummvm_zip_url, "-o", "package.zip")
	err = scummvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scummvm_bin_url := "https://downloads.scummvm.org/frs/scummvm/2.8.1/scummvm-2.8.1.tar.xz"
	scummvm_cmd_bin := exec.Command("curl", "-L", scummvm_bin_url, "-o", "binary.bin")
	err = scummvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scummvm_src_url := "https://downloads.scummvm.org/frs/scummvm/2.8.1/scummvm-2.8.1.tar.xz"
	scummvm_cmd_src := exec.Command("curl", "-L", scummvm_src_url, "-o", "source.tar.gz")
	err = scummvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scummvm_cmd_direct := exec.Command("./binary")
	err = scummvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: a52dec")
	exec.Command("latte", "install", "a52dec").Run()
	fmt.Println("Instalando dependencia: faad2")
	exec.Command("latte", "install", "faad2").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libmikmod")
	exec.Command("latte", "install", "libmikmod").Run()
	fmt.Println("Instalando dependencia: libmpeg2")
	exec.Command("latte", "install", "libmpeg2").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
