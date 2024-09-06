package main

// Fheroes2 - Recreation of the Heroes of Might and Magic II game engine
// Homepage: https://ihhub.github.io/fheroes2/

import (
	"fmt"
	
	"os/exec"
)

func installFheroes2() {
	// Método 1: Descargar y extraer .tar.gz
	fheroes2_tar_url := "https://github.com/ihhub/fheroes2/archive/refs/tags/1.1.1.tar.gz"
	fheroes2_cmd_tar := exec.Command("curl", "-L", fheroes2_tar_url, "-o", "package.tar.gz")
	err := fheroes2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fheroes2_zip_url := "https://github.com/ihhub/fheroes2/archive/refs/tags/1.1.1.zip"
	fheroes2_cmd_zip := exec.Command("curl", "-L", fheroes2_zip_url, "-o", "package.zip")
	err = fheroes2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fheroes2_bin_url := "https://github.com/ihhub/fheroes2/archive/refs/tags/1.1.1.bin"
	fheroes2_cmd_bin := exec.Command("curl", "-L", fheroes2_bin_url, "-o", "binary.bin")
	err = fheroes2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fheroes2_src_url := "https://github.com/ihhub/fheroes2/archive/refs/tags/1.1.1.src.tar.gz"
	fheroes2_cmd_src := exec.Command("curl", "-L", fheroes2_src_url, "-o", "source.tar.gz")
	err = fheroes2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fheroes2_cmd_direct := exec.Command("./binary")
	err = fheroes2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: innoextract")
exec.Command("latte", "install", "innoextract")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
}
