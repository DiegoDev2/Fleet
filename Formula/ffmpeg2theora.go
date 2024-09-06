package main

// Ffmpeg2theora - Convert video files to Ogg Theora format
// Homepage: https://gitlab.xiph.org/xiph/ffmpeg2theora

import (
	"fmt"
	
	"os/exec"
)

func installFfmpeg2theora() {
	// Método 1: Descargar y extraer .tar.gz
	ffmpeg2theora_tar_url := "https://gitlab.xiph.org/xiph/ffmpeg2theora/-/archive/0.30/ffmpeg2theora-0.30.tar.gz"
	ffmpeg2theora_cmd_tar := exec.Command("curl", "-L", ffmpeg2theora_tar_url, "-o", "package.tar.gz")
	err := ffmpeg2theora_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffmpeg2theora_zip_url := "https://gitlab.xiph.org/xiph/ffmpeg2theora/-/archive/0.30/ffmpeg2theora-0.30.zip"
	ffmpeg2theora_cmd_zip := exec.Command("curl", "-L", ffmpeg2theora_zip_url, "-o", "package.zip")
	err = ffmpeg2theora_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffmpeg2theora_bin_url := "https://gitlab.xiph.org/xiph/ffmpeg2theora/-/archive/0.30/ffmpeg2theora-0.30.bin"
	ffmpeg2theora_cmd_bin := exec.Command("curl", "-L", ffmpeg2theora_bin_url, "-o", "binary.bin")
	err = ffmpeg2theora_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffmpeg2theora_src_url := "https://gitlab.xiph.org/xiph/ffmpeg2theora/-/archive/0.30/ffmpeg2theora-0.30.src.tar.gz"
	ffmpeg2theora_cmd_src := exec.Command("curl", "-L", ffmpeg2theora_src_url, "-o", "source.tar.gz")
	err = ffmpeg2theora_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffmpeg2theora_cmd_direct := exec.Command("./binary")
	err = ffmpeg2theora_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: scons")
	exec.Command("latte", "install", "scons").Run()
	fmt.Println("Instalando dependencia: ffmpeg@4")
	exec.Command("latte", "install", "ffmpeg@4").Run()
	fmt.Println("Instalando dependencia: libkate")
	exec.Command("latte", "install", "libkate").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
}
