package main

// Mgba - Game Boy Advance emulator
// Homepage: https://mgba.io/

import (
	"fmt"
	
	"os/exec"
)

func installMgba() {
	// Método 1: Descargar y extraer .tar.gz
	mgba_tar_url := "https://github.com/mgba-emu/mgba/archive/refs/tags/0.10.3.tar.gz"
	mgba_cmd_tar := exec.Command("curl", "-L", mgba_tar_url, "-o", "package.tar.gz")
	err := mgba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mgba_zip_url := "https://github.com/mgba-emu/mgba/archive/refs/tags/0.10.3.zip"
	mgba_cmd_zip := exec.Command("curl", "-L", mgba_zip_url, "-o", "package.zip")
	err = mgba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mgba_bin_url := "https://github.com/mgba-emu/mgba/archive/refs/tags/0.10.3.bin"
	mgba_cmd_bin := exec.Command("curl", "-L", mgba_bin_url, "-o", "binary.bin")
	err = mgba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mgba_src_url := "https://github.com/mgba-emu/mgba/archive/refs/tags/0.10.3.src.tar.gz"
	mgba_cmd_src := exec.Command("curl", "-L", mgba_src_url, "-o", "source.tar.gz")
	err = mgba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mgba_cmd_direct := exec.Command("./binary")
	err = mgba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: libelf")
exec.Command("latte", "install", "libelf")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
