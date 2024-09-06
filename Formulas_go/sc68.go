package main

// Sc68 - Play music originally designed for Atari ST and Amiga computers
// Homepage: https://sc68.atari.org/project.html

import (
	"fmt"
	
	"os/exec"
)

func installSc68() {
	// Método 1: Descargar y extraer .tar.gz
	sc68_tar_url := "https://downloads.sourceforge.net/project/sc68/sc68/2.2.1/sc68-2.2.1.tar.gz"
	sc68_cmd_tar := exec.Command("curl", "-L", sc68_tar_url, "-o", "package.tar.gz")
	err := sc68_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sc68_zip_url := "https://downloads.sourceforge.net/project/sc68/sc68/2.2.1/sc68-2.2.1.zip"
	sc68_cmd_zip := exec.Command("curl", "-L", sc68_zip_url, "-o", "package.zip")
	err = sc68_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sc68_bin_url := "https://downloads.sourceforge.net/project/sc68/sc68/2.2.1/sc68-2.2.1.bin"
	sc68_cmd_bin := exec.Command("curl", "-L", sc68_bin_url, "-o", "binary.bin")
	err = sc68_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sc68_src_url := "https://downloads.sourceforge.net/project/sc68/sc68/2.2.1/sc68-2.2.1.src.tar.gz"
	sc68_cmd_src := exec.Command("curl", "-L", sc68_src_url, "-o", "source.tar.gz")
	err = sc68_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sc68_cmd_direct := exec.Command("./binary")
	err = sc68_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
