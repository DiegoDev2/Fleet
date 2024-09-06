package main

// Handbrake - Open-source video transcoder available for Linux, Mac, and Windows
// Homepage: https://handbrake.fr/

import (
	"fmt"
	
	"os/exec"
)

func installHandbrake() {
	// Método 1: Descargar y extraer .tar.gz
	handbrake_tar_url := "https://github.com/HandBrake/HandBrake/releases/download/1.7.3/HandBrake-1.7.3-source.tar.bz2"
	handbrake_cmd_tar := exec.Command("curl", "-L", handbrake_tar_url, "-o", "package.tar.gz")
	err := handbrake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	handbrake_zip_url := "https://github.com/HandBrake/HandBrake/releases/download/1.7.3/HandBrake-1.7.3-source.tar.bz2"
	handbrake_cmd_zip := exec.Command("curl", "-L", handbrake_zip_url, "-o", "package.zip")
	err = handbrake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	handbrake_bin_url := "https://github.com/HandBrake/HandBrake/releases/download/1.7.3/HandBrake-1.7.3-source.tar.bz2"
	handbrake_cmd_bin := exec.Command("curl", "-L", handbrake_bin_url, "-o", "binary.bin")
	err = handbrake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	handbrake_src_url := "https://github.com/HandBrake/HandBrake/releases/download/1.7.3/HandBrake-1.7.3-source.tar.bz2"
	handbrake_cmd_src := exec.Command("curl", "-L", handbrake_src_url, "-o", "source.tar.gz")
	err = handbrake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	handbrake_cmd_direct := exec.Command("./binary")
	err = handbrake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: yasm")
	exec.Command("latte", "install", "yasm").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libass")
	exec.Command("latte", "install", "libass").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: numactl")
	exec.Command("latte", "install", "numactl").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: speex")
	exec.Command("latte", "install", "speex").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
	fmt.Println("Instalando dependencia: x264")
	exec.Command("latte", "install", "x264").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
