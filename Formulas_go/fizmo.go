package main

// Fizmo - Z-Machine interpreter
// Homepage: https://fizmo.spellbreaker.org

import (
	"fmt"
	
	"os/exec"
)

func installFizmo() {
	// Método 1: Descargar y extraer .tar.gz
	fizmo_tar_url := "https://fizmo.spellbreaker.org/source/fizmo-0.8.5.tar.gz"
	fizmo_cmd_tar := exec.Command("curl", "-L", fizmo_tar_url, "-o", "package.tar.gz")
	err := fizmo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fizmo_zip_url := "https://fizmo.spellbreaker.org/source/fizmo-0.8.5.zip"
	fizmo_cmd_zip := exec.Command("curl", "-L", fizmo_zip_url, "-o", "package.zip")
	err = fizmo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fizmo_bin_url := "https://fizmo.spellbreaker.org/source/fizmo-0.8.5.bin"
	fizmo_cmd_bin := exec.Command("curl", "-L", fizmo_bin_url, "-o", "binary.bin")
	err = fizmo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fizmo_src_url := "https://fizmo.spellbreaker.org/source/fizmo-0.8.5.src.tar.gz"
	fizmo_cmd_src := exec.Command("curl", "-L", fizmo_src_url, "-o", "source.tar.gz")
	err = fizmo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fizmo_cmd_direct := exec.Command("./binary")
	err = fizmo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
