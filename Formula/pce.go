package main

// Pce - PC emulator
// Homepage: http://www.hampa.ch/pce/

import (
	"fmt"
	
	"os/exec"
)

func installPce() {
	// Método 1: Descargar y extraer .tar.gz
	pce_tar_url := "http://www.hampa.ch/pub/pce/pce-0.2.2.tar.gz"
	pce_cmd_tar := exec.Command("curl", "-L", pce_tar_url, "-o", "package.tar.gz")
	err := pce_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pce_zip_url := "http://www.hampa.ch/pub/pce/pce-0.2.2.zip"
	pce_cmd_zip := exec.Command("curl", "-L", pce_zip_url, "-o", "package.zip")
	err = pce_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pce_bin_url := "http://www.hampa.ch/pub/pce/pce-0.2.2.bin"
	pce_cmd_bin := exec.Command("curl", "-L", pce_bin_url, "-o", "binary.bin")
	err = pce_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pce_src_url := "http://www.hampa.ch/pub/pce/pce-0.2.2.src.tar.gz"
	pce_cmd_src := exec.Command("curl", "-L", pce_src_url, "-o", "source.tar.gz")
	err = pce_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pce_cmd_direct := exec.Command("./binary")
	err = pce_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}
