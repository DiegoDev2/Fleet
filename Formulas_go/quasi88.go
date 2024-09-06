package main

// Quasi88 - PC-8801 emulator
// Homepage: https://www.eonet.ne.jp/~showtime/quasi88/

import (
	"fmt"
	
	"os/exec"
)

func installQuasi88() {
	// Método 1: Descargar y extraer .tar.gz
	quasi88_tar_url := "https://www.eonet.ne.jp/~showtime/quasi88/release/quasi88-0.7.1.tgz"
	quasi88_cmd_tar := exec.Command("curl", "-L", quasi88_tar_url, "-o", "package.tar.gz")
	err := quasi88_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quasi88_zip_url := "https://www.eonet.ne.jp/~showtime/quasi88/release/quasi88-0.7.1.tgz"
	quasi88_cmd_zip := exec.Command("curl", "-L", quasi88_zip_url, "-o", "package.zip")
	err = quasi88_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quasi88_bin_url := "https://www.eonet.ne.jp/~showtime/quasi88/release/quasi88-0.7.1.tgz"
	quasi88_cmd_bin := exec.Command("curl", "-L", quasi88_bin_url, "-o", "binary.bin")
	err = quasi88_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quasi88_src_url := "https://www.eonet.ne.jp/~showtime/quasi88/release/quasi88-0.7.1.tgz"
	quasi88_cmd_src := exec.Command("curl", "-L", quasi88_src_url, "-o", "source.tar.gz")
	err = quasi88_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quasi88_cmd_direct := exec.Command("./binary")
	err = quasi88_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
