package main

// Audiofile - Reads and writes many common audio file formats
// Homepage: https://audiofile.68k.org/

import (
	"fmt"
	
	"os/exec"
)

func installAudiofile() {
	// Método 1: Descargar y extraer .tar.gz
	audiofile_tar_url := "https://audiofile.68k.org/audiofile-0.3.6.tar.gz"
	audiofile_cmd_tar := exec.Command("curl", "-L", audiofile_tar_url, "-o", "package.tar.gz")
	err := audiofile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	audiofile_zip_url := "https://audiofile.68k.org/audiofile-0.3.6.zip"
	audiofile_cmd_zip := exec.Command("curl", "-L", audiofile_zip_url, "-o", "package.zip")
	err = audiofile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	audiofile_bin_url := "https://audiofile.68k.org/audiofile-0.3.6.bin"
	audiofile_cmd_bin := exec.Command("curl", "-L", audiofile_bin_url, "-o", "binary.bin")
	err = audiofile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	audiofile_src_url := "https://audiofile.68k.org/audiofile-0.3.6.src.tar.gz"
	audiofile_cmd_src := exec.Command("curl", "-L", audiofile_src_url, "-o", "source.tar.gz")
	err = audiofile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	audiofile_cmd_direct := exec.Command("./binary")
	err = audiofile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
