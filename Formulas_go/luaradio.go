package main

// Luaradio - Lightweight, embeddable flow graph signal processing framework for SDR
// Homepage: https://luaradio.io/

import (
	"fmt"
	
	"os/exec"
)

func installLuaradio() {
	// Método 1: Descargar y extraer .tar.gz
	luaradio_tar_url := "https://github.com/vsergeev/luaradio/archive/refs/tags/v0.11.0.tar.gz"
	luaradio_cmd_tar := exec.Command("curl", "-L", luaradio_tar_url, "-o", "package.tar.gz")
	err := luaradio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luaradio_zip_url := "https://github.com/vsergeev/luaradio/archive/refs/tags/v0.11.0.zip"
	luaradio_cmd_zip := exec.Command("curl", "-L", luaradio_zip_url, "-o", "package.zip")
	err = luaradio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luaradio_bin_url := "https://github.com/vsergeev/luaradio/archive/refs/tags/v0.11.0.bin"
	luaradio_cmd_bin := exec.Command("curl", "-L", luaradio_bin_url, "-o", "binary.bin")
	err = luaradio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luaradio_src_url := "https://github.com/vsergeev/luaradio/archive/refs/tags/v0.11.0.src.tar.gz"
	luaradio_cmd_src := exec.Command("curl", "-L", luaradio_src_url, "-o", "source.tar.gz")
	err = luaradio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luaradio_cmd_direct := exec.Command("./binary")
	err = luaradio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: liquid-dsp")
exec.Command("latte", "install", "liquid-dsp")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
}
