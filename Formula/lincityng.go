package main

// LincityNg - City simulation game
// Homepage: https://github.com/lincity-ng/lincity-ng/

import (
	"fmt"
	
	"os/exec"
)

func installLincityNg() {
	// Método 1: Descargar y extraer .tar.gz
	lincityng_tar_url := "https://github.com/lincity-ng/lincity-ng/archive/refs/tags/lincity-ng-2.0.tar.gz"
	lincityng_cmd_tar := exec.Command("curl", "-L", lincityng_tar_url, "-o", "package.tar.gz")
	err := lincityng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lincityng_zip_url := "https://github.com/lincity-ng/lincity-ng/archive/refs/tags/lincity-ng-2.0.zip"
	lincityng_cmd_zip := exec.Command("curl", "-L", lincityng_zip_url, "-o", "package.zip")
	err = lincityng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lincityng_bin_url := "https://github.com/lincity-ng/lincity-ng/archive/refs/tags/lincity-ng-2.0.bin"
	lincityng_cmd_bin := exec.Command("curl", "-L", lincityng_bin_url, "-o", "binary.bin")
	err = lincityng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lincityng_src_url := "https://github.com/lincity-ng/lincity-ng/archive/refs/tags/lincity-ng-2.0.src.tar.gz"
	lincityng_cmd_src := exec.Command("curl", "-L", lincityng_src_url, "-o", "source.tar.gz")
	err = lincityng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lincityng_cmd_direct := exec.Command("./binary")
	err = lincityng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: jam")
	exec.Command("latte", "install", "jam").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: physfs")
	exec.Command("latte", "install", "physfs").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_gfx")
	exec.Command("latte", "install", "sdl_gfx").Run()
	fmt.Println("Instalando dependencia: sdl_image")
	exec.Command("latte", "install", "sdl_image").Run()
	fmt.Println("Instalando dependencia: sdl_mixer")
	exec.Command("latte", "install", "sdl_mixer").Run()
	fmt.Println("Instalando dependencia: sdl_ttf")
	exec.Command("latte", "install", "sdl_ttf").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
