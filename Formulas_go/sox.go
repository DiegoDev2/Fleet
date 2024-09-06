package main

// Sox - SOund eXchange: universal sound sample translator
// Homepage: https://sox.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSox() {
	// Método 1: Descargar y extraer .tar.gz
	sox_tar_url := "https://downloads.sourceforge.net/project/sox/sox/14.4.2/sox-14.4.2.tar.gz"
	sox_cmd_tar := exec.Command("curl", "-L", sox_tar_url, "-o", "package.tar.gz")
	err := sox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sox_zip_url := "https://downloads.sourceforge.net/project/sox/sox/14.4.2/sox-14.4.2.zip"
	sox_cmd_zip := exec.Command("curl", "-L", sox_zip_url, "-o", "package.zip")
	err = sox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sox_bin_url := "https://downloads.sourceforge.net/project/sox/sox/14.4.2/sox-14.4.2.bin"
	sox_cmd_bin := exec.Command("curl", "-L", sox_bin_url, "-o", "binary.bin")
	err = sox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sox_src_url := "https://downloads.sourceforge.net/project/sox/sox/14.4.2/sox-14.4.2.src.tar.gz"
	sox_cmd_src := exec.Command("curl", "-L", sox_src_url, "-o", "source.tar.gz")
	err = sox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sox_cmd_direct := exec.Command("./binary")
	err = sox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mad")
exec.Command("latte", "install", "mad")
	fmt.Println("Instalando dependencia: opusfile")
exec.Command("latte", "install", "opusfile")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
