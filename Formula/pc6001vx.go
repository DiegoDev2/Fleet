package main

// Pc6001vx - PC-6001 emulator
// Homepage: https://github.com/eighttails/PC6001VX

import (
	"fmt"
	
	"os/exec"
)

func installPc6001vx() {
	// Método 1: Descargar y extraer .tar.gz
	pc6001vx_tar_url := "https://eighttails.up.seesaa.net/bin/PC6001VX_4.2.9_src.tar.gz"
	pc6001vx_cmd_tar := exec.Command("curl", "-L", pc6001vx_tar_url, "-o", "package.tar.gz")
	err := pc6001vx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pc6001vx_zip_url := "https://eighttails.up.seesaa.net/bin/PC6001VX_4.2.9_src.zip"
	pc6001vx_cmd_zip := exec.Command("curl", "-L", pc6001vx_zip_url, "-o", "package.zip")
	err = pc6001vx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pc6001vx_bin_url := "https://eighttails.up.seesaa.net/bin/PC6001VX_4.2.9_src.bin"
	pc6001vx_cmd_bin := exec.Command("curl", "-L", pc6001vx_bin_url, "-o", "binary.bin")
	err = pc6001vx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pc6001vx_src_url := "https://eighttails.up.seesaa.net/bin/PC6001VX_4.2.9_src.src.tar.gz"
	pc6001vx_cmd_src := exec.Command("curl", "-L", pc6001vx_src_url, "-o", "source.tar.gz")
	err = pc6001vx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pc6001vx_cmd_direct := exec.Command("./binary")
	err = pc6001vx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
