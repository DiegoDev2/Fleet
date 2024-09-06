package main

// Hatari - Atari ST/STE/TT/Falcon emulator
// Homepage: https://hatari.tuxfamily.org

import (
	"fmt"
	
	"os/exec"
)

func installHatari() {
	// Método 1: Descargar y extraer .tar.gz
	hatari_tar_url := "https://download.tuxfamily.org/hatari/2.5.0/hatari-2.5.0.tar.bz2"
	hatari_cmd_tar := exec.Command("curl", "-L", hatari_tar_url, "-o", "package.tar.gz")
	err := hatari_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hatari_zip_url := "https://download.tuxfamily.org/hatari/2.5.0/hatari-2.5.0.tar.bz2"
	hatari_cmd_zip := exec.Command("curl", "-L", hatari_zip_url, "-o", "package.zip")
	err = hatari_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hatari_bin_url := "https://download.tuxfamily.org/hatari/2.5.0/hatari-2.5.0.tar.bz2"
	hatari_cmd_bin := exec.Command("curl", "-L", hatari_bin_url, "-o", "binary.bin")
	err = hatari_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hatari_src_url := "https://download.tuxfamily.org/hatari/2.5.0/hatari-2.5.0.tar.bz2"
	hatari_cmd_src := exec.Command("curl", "-L", hatari_src_url, "-o", "source.tar.gz")
	err = hatari_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hatari_cmd_direct := exec.Command("./binary")
	err = hatari_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
