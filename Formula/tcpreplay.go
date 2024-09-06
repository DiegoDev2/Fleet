package main

// Tcpreplay - Replay saved tcpdump files at arbitrary speeds
// Homepage: https://tcpreplay.appneta.com/

import (
	"fmt"
	
	"os/exec"
)

func installTcpreplay() {
	// Método 1: Descargar y extraer .tar.gz
	tcpreplay_tar_url := "https://github.com/appneta/tcpreplay/releases/download/v4.5.1/tcpreplay-4.5.1.tar.gz"
	tcpreplay_cmd_tar := exec.Command("curl", "-L", tcpreplay_tar_url, "-o", "package.tar.gz")
	err := tcpreplay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcpreplay_zip_url := "https://github.com/appneta/tcpreplay/releases/download/v4.5.1/tcpreplay-4.5.1.zip"
	tcpreplay_cmd_zip := exec.Command("curl", "-L", tcpreplay_zip_url, "-o", "package.zip")
	err = tcpreplay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcpreplay_bin_url := "https://github.com/appneta/tcpreplay/releases/download/v4.5.1/tcpreplay-4.5.1.bin"
	tcpreplay_cmd_bin := exec.Command("curl", "-L", tcpreplay_bin_url, "-o", "binary.bin")
	err = tcpreplay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcpreplay_src_url := "https://github.com/appneta/tcpreplay/releases/download/v4.5.1/tcpreplay-4.5.1.src.tar.gz"
	tcpreplay_cmd_src := exec.Command("curl", "-L", tcpreplay_src_url, "-o", "source.tar.gz")
	err = tcpreplay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcpreplay_cmd_direct := exec.Command("./binary")
	err = tcpreplay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libdnet")
	exec.Command("latte", "install", "libdnet").Run()
}
