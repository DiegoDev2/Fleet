package main

// Speex - Audio codec designed for speech
// Homepage: https://speex.org/

import (
	"fmt"
	
	"os/exec"
)

func installSpeex() {
	// Método 1: Descargar y extraer .tar.gz
	speex_tar_url := "https://downloads.xiph.org/releases/speex/speex-1.2.1.tar.gz"
	speex_cmd_tar := exec.Command("curl", "-L", speex_tar_url, "-o", "package.tar.gz")
	err := speex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speex_zip_url := "https://downloads.xiph.org/releases/speex/speex-1.2.1.zip"
	speex_cmd_zip := exec.Command("curl", "-L", speex_zip_url, "-o", "package.zip")
	err = speex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speex_bin_url := "https://downloads.xiph.org/releases/speex/speex-1.2.1.bin"
	speex_cmd_bin := exec.Command("curl", "-L", speex_bin_url, "-o", "binary.bin")
	err = speex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speex_src_url := "https://downloads.xiph.org/releases/speex/speex-1.2.1.src.tar.gz"
	speex_cmd_src := exec.Command("curl", "-L", speex_src_url, "-o", "source.tar.gz")
	err = speex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speex_cmd_direct := exec.Command("./binary")
	err = speex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
}
