package main

// DosboxX - DOSBox with accurate emulation and wide testing
// Homepage: https://dosbox-x.com/

import (
	"fmt"
	
	"os/exec"
)

func installDosboxX() {
	// Método 1: Descargar y extraer .tar.gz
	dosboxx_tar_url := "https://github.com/joncampbell123/dosbox-x/archive/refs/tags/dosbox-x-v2024.07.01.tar.gz"
	dosboxx_cmd_tar := exec.Command("curl", "-L", dosboxx_tar_url, "-o", "package.tar.gz")
	err := dosboxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dosboxx_zip_url := "https://github.com/joncampbell123/dosbox-x/archive/refs/tags/dosbox-x-v2024.07.01.zip"
	dosboxx_cmd_zip := exec.Command("curl", "-L", dosboxx_zip_url, "-o", "package.zip")
	err = dosboxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dosboxx_bin_url := "https://github.com/joncampbell123/dosbox-x/archive/refs/tags/dosbox-x-v2024.07.01.bin"
	dosboxx_cmd_bin := exec.Command("curl", "-L", dosboxx_bin_url, "-o", "binary.bin")
	err = dosboxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dosboxx_src_url := "https://github.com/joncampbell123/dosbox-x/archive/refs/tags/dosbox-x-v2024.07.01.src.tar.gz"
	dosboxx_cmd_src := exec.Command("curl", "-L", dosboxx_src_url, "-o", "source.tar.gz")
	err = dosboxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dosboxx_cmd_direct := exec.Command("./binary")
	err = dosboxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libslirp")
exec.Command("latte", "install", "libslirp")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: linux-headers@5.15")
exec.Command("latte", "install", "linux-headers@5.15")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
}
