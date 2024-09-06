package main

// Instead - Interpreter of simple text adventures
// Homepage: https://instead.hugeping.ru/

import (
	"fmt"
	
	"os/exec"
)

func installInstead() {
	// Método 1: Descargar y extraer .tar.gz
	instead_tar_url := "https://github.com/instead-hub/instead/archive/refs/tags/3.5.2.tar.gz"
	instead_cmd_tar := exec.Command("curl", "-L", instead_tar_url, "-o", "package.tar.gz")
	err := instead_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	instead_zip_url := "https://github.com/instead-hub/instead/archive/refs/tags/3.5.2.zip"
	instead_cmd_zip := exec.Command("curl", "-L", instead_zip_url, "-o", "package.zip")
	err = instead_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	instead_bin_url := "https://github.com/instead-hub/instead/archive/refs/tags/3.5.2.bin"
	instead_cmd_bin := exec.Command("curl", "-L", instead_bin_url, "-o", "binary.bin")
	err = instead_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	instead_src_url := "https://github.com/instead-hub/instead/archive/refs/tags/3.5.2.src.tar.gz"
	instead_cmd_src := exec.Command("curl", "-L", instead_src_url, "-o", "source.tar.gz")
	err = instead_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	instead_cmd_direct := exec.Command("./binary")
	err = instead_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: sdl2_ttf")
	exec.Command("latte", "install", "sdl2_ttf").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
