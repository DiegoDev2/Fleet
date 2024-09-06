package main

// Libreplaygain - Library to implement ReplayGain standard for audio
// Homepage: https://www.musepack.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibreplaygain() {
	// Método 1: Descargar y extraer .tar.gz
	libreplaygain_tar_url := "https://files.musepack.net/source/libreplaygain_r475.tar.gz"
	libreplaygain_cmd_tar := exec.Command("curl", "-L", libreplaygain_tar_url, "-o", "package.tar.gz")
	err := libreplaygain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libreplaygain_zip_url := "https://files.musepack.net/source/libreplaygain_r475.zip"
	libreplaygain_cmd_zip := exec.Command("curl", "-L", libreplaygain_zip_url, "-o", "package.zip")
	err = libreplaygain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libreplaygain_bin_url := "https://files.musepack.net/source/libreplaygain_r475.bin"
	libreplaygain_cmd_bin := exec.Command("curl", "-L", libreplaygain_bin_url, "-o", "binary.bin")
	err = libreplaygain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libreplaygain_src_url := "https://files.musepack.net/source/libreplaygain_r475.src.tar.gz"
	libreplaygain_cmd_src := exec.Command("curl", "-L", libreplaygain_src_url, "-o", "source.tar.gz")
	err = libreplaygain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libreplaygain_cmd_direct := exec.Command("./binary")
	err = libreplaygain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
