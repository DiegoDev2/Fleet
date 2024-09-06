package main

// Dosbox - DOS Emulator
// Homepage: https://www.dosbox.com/

import (
	"fmt"
	
	"os/exec"
)

func installDosbox() {
	// Método 1: Descargar y extraer .tar.gz
	dosbox_tar_url := "https://downloads.sourceforge.net/project/dosbox/dosbox/0.74-3/dosbox-0.74-3.tar.gz"
	dosbox_cmd_tar := exec.Command("curl", "-L", dosbox_tar_url, "-o", "package.tar.gz")
	err := dosbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dosbox_zip_url := "https://downloads.sourceforge.net/project/dosbox/dosbox/0.74-3/dosbox-0.74-3.zip"
	dosbox_cmd_zip := exec.Command("curl", "-L", dosbox_zip_url, "-o", "package.zip")
	err = dosbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dosbox_bin_url := "https://downloads.sourceforge.net/project/dosbox/dosbox/0.74-3/dosbox-0.74-3.bin"
	dosbox_cmd_bin := exec.Command("curl", "-L", dosbox_bin_url, "-o", "binary.bin")
	err = dosbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dosbox_src_url := "https://downloads.sourceforge.net/project/dosbox/dosbox/0.74-3/dosbox-0.74-3.src.tar.gz"
	dosbox_cmd_src := exec.Command("curl", "-L", dosbox_src_url, "-o", "source.tar.gz")
	err = dosbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dosbox_cmd_direct := exec.Command("./binary")
	err = dosbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_net")
	exec.Command("latte", "install", "sdl_net").Run()
	fmt.Println("Instalando dependencia: sdl_sound")
	exec.Command("latte", "install", "sdl_sound").Run()
}
