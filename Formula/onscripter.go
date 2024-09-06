package main

// Onscripter - NScripter-compatible visual novel engine
// Homepage: https://onscripter.osdn.jp/onscripter.html

import (
	"fmt"
	
	"os/exec"
)

func installOnscripter() {
	// Método 1: Descargar y extraer .tar.gz
	onscripter_tar_url := "https://onscripter.osdn.jp/onscripter-20230825.tar.gz"
	onscripter_cmd_tar := exec.Command("curl", "-L", onscripter_tar_url, "-o", "package.tar.gz")
	err := onscripter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onscripter_zip_url := "https://onscripter.osdn.jp/onscripter-20230825.zip"
	onscripter_cmd_zip := exec.Command("curl", "-L", onscripter_zip_url, "-o", "package.zip")
	err = onscripter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onscripter_bin_url := "https://onscripter.osdn.jp/onscripter-20230825.bin"
	onscripter_cmd_bin := exec.Command("curl", "-L", onscripter_bin_url, "-o", "binary.bin")
	err = onscripter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onscripter_src_url := "https://onscripter.osdn.jp/onscripter-20230825.src.tar.gz"
	onscripter_cmd_src := exec.Command("curl", "-L", onscripter_src_url, "-o", "source.tar.gz")
	err = onscripter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onscripter_cmd_direct := exec.Command("./binary")
	err = onscripter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_image")
	exec.Command("latte", "install", "sdl_image").Run()
	fmt.Println("Instalando dependencia: sdl_mixer")
	exec.Command("latte", "install", "sdl_mixer").Run()
	fmt.Println("Instalando dependencia: sdl_ttf")
	exec.Command("latte", "install", "sdl_ttf").Run()
	fmt.Println("Instalando dependencia: smpeg")
	exec.Command("latte", "install", "smpeg").Run()
}
