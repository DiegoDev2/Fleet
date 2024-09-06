package main

// Mednafen - Multi-system emulator
// Homepage: https://mednafen.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installMednafen() {
	// Método 1: Descargar y extraer .tar.gz
	mednafen_tar_url := "https://mednafen.github.io/releases/files/mednafen-1.32.1.tar.xz"
	mednafen_cmd_tar := exec.Command("curl", "-L", mednafen_tar_url, "-o", "package.tar.gz")
	err := mednafen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mednafen_zip_url := "https://mednafen.github.io/releases/files/mednafen-1.32.1.tar.xz"
	mednafen_cmd_zip := exec.Command("curl", "-L", mednafen_zip_url, "-o", "package.zip")
	err = mednafen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mednafen_bin_url := "https://mednafen.github.io/releases/files/mednafen-1.32.1.tar.xz"
	mednafen_cmd_bin := exec.Command("curl", "-L", mednafen_bin_url, "-o", "binary.bin")
	err = mednafen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mednafen_src_url := "https://mednafen.github.io/releases/files/mednafen-1.32.1.tar.xz"
	mednafen_cmd_src := exec.Command("curl", "-L", mednafen_src_url, "-o", "source.tar.gz")
	err = mednafen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mednafen_cmd_direct := exec.Command("./binary")
	err = mednafen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: musepack")
	exec.Command("latte", "install", "musepack").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
