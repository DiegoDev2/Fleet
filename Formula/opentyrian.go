package main

// OpenTyrian - Open-source port of Tyrian
// Homepage: https://github.com/opentyrian/opentyrian

import (
	"fmt"
	
	"os/exec"
)

func installOpenTyrian() {
	// Método 1: Descargar y extraer .tar.gz
	opentyrian_tar_url := "https://github.com/opentyrian/opentyrian/archive/refs/tags/v2.1.20221123.tar.gz"
	opentyrian_cmd_tar := exec.Command("curl", "-L", opentyrian_tar_url, "-o", "package.tar.gz")
	err := opentyrian_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opentyrian_zip_url := "https://github.com/opentyrian/opentyrian/archive/refs/tags/v2.1.20221123.zip"
	opentyrian_cmd_zip := exec.Command("curl", "-L", opentyrian_zip_url, "-o", "package.zip")
	err = opentyrian_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opentyrian_bin_url := "https://github.com/opentyrian/opentyrian/archive/refs/tags/v2.1.20221123.bin"
	opentyrian_cmd_bin := exec.Command("curl", "-L", opentyrian_bin_url, "-o", "binary.bin")
	err = opentyrian_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opentyrian_src_url := "https://github.com/opentyrian/opentyrian/archive/refs/tags/v2.1.20221123.src.tar.gz"
	opentyrian_cmd_src := exec.Command("curl", "-L", opentyrian_src_url, "-o", "source.tar.gz")
	err = opentyrian_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opentyrian_cmd_direct := exec.Command("./binary")
	err = opentyrian_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_net")
	exec.Command("latte", "install", "sdl2_net").Run()
}
