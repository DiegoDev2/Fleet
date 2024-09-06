package main

// SoundTouch - Audio processing library
// Homepage: https://www.surina.net/soundtouch/

import (
	"fmt"
	
	"os/exec"
)

func installSoundTouch() {
	// Método 1: Descargar y extraer .tar.gz
	soundtouch_tar_url := "https://codeberg.org/soundtouch/soundtouch/archive/2.3.3.tar.gz"
	soundtouch_cmd_tar := exec.Command("curl", "-L", soundtouch_tar_url, "-o", "package.tar.gz")
	err := soundtouch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soundtouch_zip_url := "https://codeberg.org/soundtouch/soundtouch/archive/2.3.3.zip"
	soundtouch_cmd_zip := exec.Command("curl", "-L", soundtouch_zip_url, "-o", "package.zip")
	err = soundtouch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soundtouch_bin_url := "https://codeberg.org/soundtouch/soundtouch/archive/2.3.3.bin"
	soundtouch_cmd_bin := exec.Command("curl", "-L", soundtouch_bin_url, "-o", "binary.bin")
	err = soundtouch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soundtouch_src_url := "https://codeberg.org/soundtouch/soundtouch/archive/2.3.3.src.tar.gz"
	soundtouch_cmd_src := exec.Command("curl", "-L", soundtouch_src_url, "-o", "source.tar.gz")
	err = soundtouch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soundtouch_cmd_direct := exec.Command("./binary")
	err = soundtouch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
