package main

// Opus - Audio codec
// Homepage: https://www.opus-codec.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpus() {
	// Método 1: Descargar y extraer .tar.gz
	opus_tar_url := "https://ftp.osuosl.org/pub/xiph/releases/opus/opus-1.5.2.tar.gz"
	opus_cmd_tar := exec.Command("curl", "-L", opus_tar_url, "-o", "package.tar.gz")
	err := opus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opus_zip_url := "https://ftp.osuosl.org/pub/xiph/releases/opus/opus-1.5.2.zip"
	opus_cmd_zip := exec.Command("curl", "-L", opus_zip_url, "-o", "package.zip")
	err = opus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opus_bin_url := "https://ftp.osuosl.org/pub/xiph/releases/opus/opus-1.5.2.bin"
	opus_cmd_bin := exec.Command("curl", "-L", opus_bin_url, "-o", "binary.bin")
	err = opus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opus_src_url := "https://ftp.osuosl.org/pub/xiph/releases/opus/opus-1.5.2.src.tar.gz"
	opus_cmd_src := exec.Command("curl", "-L", opus_src_url, "-o", "source.tar.gz")
	err = opus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opus_cmd_direct := exec.Command("./binary")
	err = opus_cmd_direct.Run()
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
