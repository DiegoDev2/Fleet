package main

// Openjazz - Open source Jazz Jackrabit engine
// Homepage: https://www.alister.eu/jazz/oj/

import (
	"fmt"
	
	"os/exec"
)

func installOpenjazz() {
	// Método 1: Descargar y extraer .tar.gz
	openjazz_tar_url := "https://github.com/AlisterT/openjazz/archive/refs/tags/20231028.tar.gz"
	openjazz_cmd_tar := exec.Command("curl", "-L", openjazz_tar_url, "-o", "package.tar.gz")
	err := openjazz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjazz_zip_url := "https://github.com/AlisterT/openjazz/archive/refs/tags/20231028.zip"
	openjazz_cmd_zip := exec.Command("curl", "-L", openjazz_zip_url, "-o", "package.zip")
	err = openjazz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjazz_bin_url := "https://github.com/AlisterT/openjazz/archive/refs/tags/20231028.bin"
	openjazz_cmd_bin := exec.Command("curl", "-L", openjazz_bin_url, "-o", "binary.bin")
	err = openjazz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjazz_src_url := "https://github.com/AlisterT/openjazz/archive/refs/tags/20231028.src.tar.gz"
	openjazz_cmd_src := exec.Command("curl", "-L", openjazz_src_url, "-o", "source.tar.gz")
	err = openjazz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjazz_cmd_direct := exec.Command("./binary")
	err = openjazz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_net")
exec.Command("latte", "install", "sdl2_net")
}
