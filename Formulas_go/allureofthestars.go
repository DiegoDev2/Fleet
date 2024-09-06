package main

// Allureofthestars - Near-future Sci-Fi roguelike and tactical squad combat game
// Homepage: https://www.allureofthestars.com/

import (
	"fmt"
	
	"os/exec"
)

func installAllureofthestars() {
	// Método 1: Descargar y extraer .tar.gz
	allureofthestars_tar_url := "https://hackage.haskell.org/package/Allure-0.11.0.0/Allure-0.11.0.0.tar.gz"
	allureofthestars_cmd_tar := exec.Command("curl", "-L", allureofthestars_tar_url, "-o", "package.tar.gz")
	err := allureofthestars_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	allureofthestars_zip_url := "https://hackage.haskell.org/package/Allure-0.11.0.0/Allure-0.11.0.0.zip"
	allureofthestars_cmd_zip := exec.Command("curl", "-L", allureofthestars_zip_url, "-o", "package.zip")
	err = allureofthestars_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	allureofthestars_bin_url := "https://hackage.haskell.org/package/Allure-0.11.0.0/Allure-0.11.0.0.bin"
	allureofthestars_cmd_bin := exec.Command("curl", "-L", allureofthestars_bin_url, "-o", "binary.bin")
	err = allureofthestars_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	allureofthestars_src_url := "https://hackage.haskell.org/package/Allure-0.11.0.0/Allure-0.11.0.0.src.tar.gz"
	allureofthestars_cmd_src := exec.Command("curl", "-L", allureofthestars_src_url, "-o", "source.tar.gz")
	err = allureofthestars_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	allureofthestars_cmd_direct := exec.Command("./binary")
	err = allureofthestars_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.6")
exec.Command("latte", "install", "ghc@9.6")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
}
