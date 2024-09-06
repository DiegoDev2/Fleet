package main

// Uade - Play Amiga tunes through UAE emulation
// Homepage: https://zakalwe.fi/uade/

import (
	"fmt"
	
	"os/exec"
)

func installUade() {
	// Método 1: Descargar y extraer .tar.gz
	uade_tar_url := "https://zakalwe.fi/uade/uade3/uade-3.03.tar.bz2"
	uade_cmd_tar := exec.Command("curl", "-L", uade_tar_url, "-o", "package.tar.gz")
	err := uade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uade_zip_url := "https://zakalwe.fi/uade/uade3/uade-3.03.tar.bz2"
	uade_cmd_zip := exec.Command("curl", "-L", uade_zip_url, "-o", "package.zip")
	err = uade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uade_bin_url := "https://zakalwe.fi/uade/uade3/uade-3.03.tar.bz2"
	uade_cmd_bin := exec.Command("curl", "-L", uade_bin_url, "-o", "binary.bin")
	err = uade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uade_src_url := "https://zakalwe.fi/uade/uade3/uade-3.03.tar.bz2"
	uade_cmd_src := exec.Command("curl", "-L", uade_src_url, "-o", "source.tar.gz")
	err = uade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uade_cmd_direct := exec.Command("./binary")
	err = uade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
}
