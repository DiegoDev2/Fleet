package main

// Xroar - Dragon and Tandy 8-bit computer emulator
// Homepage: https://www.6809.org.uk/xroar/

import (
	"fmt"
	
	"os/exec"
)

func installXroar() {
	// Método 1: Descargar y extraer .tar.gz
	xroar_tar_url := "https://www.6809.org.uk/xroar/dl/xroar-1.5.5.tar.gz"
	xroar_cmd_tar := exec.Command("curl", "-L", xroar_tar_url, "-o", "package.tar.gz")
	err := xroar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xroar_zip_url := "https://www.6809.org.uk/xroar/dl/xroar-1.5.5.zip"
	xroar_cmd_zip := exec.Command("curl", "-L", xroar_zip_url, "-o", "package.zip")
	err = xroar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xroar_bin_url := "https://www.6809.org.uk/xroar/dl/xroar-1.5.5.bin"
	xroar_cmd_bin := exec.Command("curl", "-L", xroar_bin_url, "-o", "binary.bin")
	err = xroar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xroar_src_url := "https://www.6809.org.uk/xroar/dl/xroar-1.5.5.src.tar.gz"
	xroar_cmd_src := exec.Command("curl", "-L", xroar_src_url, "-o", "source.tar.gz")
	err = xroar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xroar_cmd_direct := exec.Command("./binary")
	err = xroar_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
