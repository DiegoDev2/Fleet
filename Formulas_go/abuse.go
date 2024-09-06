package main

// Abuse - Dark 2D side-scrolling platform game
// Homepage: http://abuse.zoy.org/

import (
	"fmt"
	
	"os/exec"
)

func installAbuse() {
	// Método 1: Descargar y extraer .tar.gz
	abuse_tar_url := "http://abuse.zoy.org/raw-attachment/wiki/download/abuse-0.8.tar.gz"
	abuse_cmd_tar := exec.Command("curl", "-L", abuse_tar_url, "-o", "package.tar.gz")
	err := abuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abuse_zip_url := "http://abuse.zoy.org/raw-attachment/wiki/download/abuse-0.8.zip"
	abuse_cmd_zip := exec.Command("curl", "-L", abuse_zip_url, "-o", "package.zip")
	err = abuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abuse_bin_url := "http://abuse.zoy.org/raw-attachment/wiki/download/abuse-0.8.bin"
	abuse_cmd_bin := exec.Command("curl", "-L", abuse_bin_url, "-o", "binary.bin")
	err = abuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abuse_src_url := "http://abuse.zoy.org/raw-attachment/wiki/download/abuse-0.8.src.tar.gz"
	abuse_cmd_src := exec.Command("curl", "-L", abuse_src_url, "-o", "source.tar.gz")
	err = abuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abuse_cmd_direct := exec.Command("./binary")
	err = abuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: sdl_mixer")
exec.Command("latte", "install", "sdl_mixer")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
