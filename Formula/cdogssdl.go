package main

// CdogsSdl - Classic overhead run-and-gun game
// Homepage: https://cxong.github.io/cdogs-sdl/

import (
	"fmt"
	
	"os/exec"
)

func installCdogsSdl() {
	// Método 1: Descargar y extraer .tar.gz
	cdogssdl_tar_url := "https://github.com/cxong/cdogs-sdl/archive/refs/tags/2.1.0.tar.gz"
	cdogssdl_cmd_tar := exec.Command("curl", "-L", cdogssdl_tar_url, "-o", "package.tar.gz")
	err := cdogssdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdogssdl_zip_url := "https://github.com/cxong/cdogs-sdl/archive/refs/tags/2.1.0.zip"
	cdogssdl_cmd_zip := exec.Command("curl", "-L", cdogssdl_zip_url, "-o", "package.zip")
	err = cdogssdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdogssdl_bin_url := "https://github.com/cxong/cdogs-sdl/archive/refs/tags/2.1.0.bin"
	cdogssdl_cmd_bin := exec.Command("curl", "-L", cdogssdl_bin_url, "-o", "binary.bin")
	err = cdogssdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdogssdl_src_url := "https://github.com/cxong/cdogs-sdl/archive/refs/tags/2.1.0.src.tar.gz"
	cdogssdl_cmd_src := exec.Command("curl", "-L", cdogssdl_src_url, "-o", "source.tar.gz")
	err = cdogssdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdogssdl_cmd_direct := exec.Command("./binary")
	err = cdogssdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
