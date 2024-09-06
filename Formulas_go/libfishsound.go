package main

// Libfishsound - Decode and encode audio data using the Xiph.org codecs
// Homepage: https://xiph.org/fishsound/

import (
	"fmt"
	
	"os/exec"
)

func installLibfishsound() {
	// Método 1: Descargar y extraer .tar.gz
	libfishsound_tar_url := "https://downloads.xiph.org/releases/libfishsound/libfishsound-1.0.0.tar.gz"
	libfishsound_cmd_tar := exec.Command("curl", "-L", libfishsound_tar_url, "-o", "package.tar.gz")
	err := libfishsound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfishsound_zip_url := "https://downloads.xiph.org/releases/libfishsound/libfishsound-1.0.0.zip"
	libfishsound_cmd_zip := exec.Command("curl", "-L", libfishsound_zip_url, "-o", "package.zip")
	err = libfishsound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfishsound_bin_url := "https://downloads.xiph.org/releases/libfishsound/libfishsound-1.0.0.bin"
	libfishsound_cmd_bin := exec.Command("curl", "-L", libfishsound_bin_url, "-o", "binary.bin")
	err = libfishsound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfishsound_src_url := "https://downloads.xiph.org/releases/libfishsound/libfishsound-1.0.0.src.tar.gz"
	libfishsound_cmd_src := exec.Command("curl", "-L", libfishsound_src_url, "-o", "source.tar.gz")
	err = libfishsound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfishsound_cmd_direct := exec.Command("./binary")
	err = libfishsound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
}
