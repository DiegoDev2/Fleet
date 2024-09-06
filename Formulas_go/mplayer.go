package main

// Mplayer - UNIX movie player
// Homepage: https://mplayerhq.hu/

import (
	"fmt"
	
	"os/exec"
)

func installMplayer() {
	// Método 1: Descargar y extraer .tar.gz
	mplayer_tar_url := "https://mplayerhq.hu/MPlayer/releases/MPlayer-1.5.tar.xz"
	mplayer_cmd_tar := exec.Command("curl", "-L", mplayer_tar_url, "-o", "package.tar.gz")
	err := mplayer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mplayer_zip_url := "https://mplayerhq.hu/MPlayer/releases/MPlayer-1.5.tar.xz"
	mplayer_cmd_zip := exec.Command("curl", "-L", mplayer_zip_url, "-o", "package.zip")
	err = mplayer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mplayer_bin_url := "https://mplayerhq.hu/MPlayer/releases/MPlayer-1.5.tar.xz"
	mplayer_cmd_bin := exec.Command("curl", "-L", mplayer_bin_url, "-o", "binary.bin")
	err = mplayer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mplayer_src_url := "https://mplayerhq.hu/MPlayer/releases/MPlayer-1.5.tar.xz"
	mplayer_cmd_src := exec.Command("curl", "-L", mplayer_src_url, "-o", "source.tar.gz")
	err = mplayer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mplayer_cmd_direct := exec.Command("./binary")
	err = mplayer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: yasm")
exec.Command("latte", "install", "yasm")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libcaca")
exec.Command("latte", "install", "libcaca")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
