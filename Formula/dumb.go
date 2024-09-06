package main

// Dumb - IT, XM, S3M and MOD player library
// Homepage: https://dumb.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDumb() {
	// Método 1: Descargar y extraer .tar.gz
	dumb_tar_url := "https://github.com/kode54/dumb/archive/refs/tags/2.0.3.tar.gz"
	dumb_cmd_tar := exec.Command("curl", "-L", dumb_tar_url, "-o", "package.tar.gz")
	err := dumb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dumb_zip_url := "https://github.com/kode54/dumb/archive/refs/tags/2.0.3.zip"
	dumb_cmd_zip := exec.Command("curl", "-L", dumb_zip_url, "-o", "package.zip")
	err = dumb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dumb_bin_url := "https://github.com/kode54/dumb/archive/refs/tags/2.0.3.bin"
	dumb_cmd_bin := exec.Command("curl", "-L", dumb_bin_url, "-o", "binary.bin")
	err = dumb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dumb_src_url := "https://github.com/kode54/dumb/archive/refs/tags/2.0.3.src.tar.gz"
	dumb_cmd_src := exec.Command("curl", "-L", dumb_src_url, "-o", "source.tar.gz")
	err = dumb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dumb_cmd_direct := exec.Command("./binary")
	err = dumb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: argtable")
	exec.Command("latte", "install", "argtable").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
