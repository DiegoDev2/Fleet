package main

// Xinit - Start the X Window System server
// Homepage: https://gitlab.freedesktop.org/xorg/app/xinit

import (
	"fmt"
	
	"os/exec"
)

func installXinit() {
	// Método 1: Descargar y extraer .tar.gz
	xinit_tar_url := "https://www.x.org/releases/individual/app/xinit-1.4.2.tar.xz"
	xinit_cmd_tar := exec.Command("curl", "-L", xinit_tar_url, "-o", "package.tar.gz")
	err := xinit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xinit_zip_url := "https://www.x.org/releases/individual/app/xinit-1.4.2.tar.xz"
	xinit_cmd_zip := exec.Command("curl", "-L", xinit_zip_url, "-o", "package.zip")
	err = xinit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xinit_bin_url := "https://www.x.org/releases/individual/app/xinit-1.4.2.tar.xz"
	xinit_cmd_bin := exec.Command("curl", "-L", xinit_bin_url, "-o", "binary.bin")
	err = xinit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xinit_src_url := "https://www.x.org/releases/individual/app/xinit-1.4.2.tar.xz"
	xinit_cmd_src := exec.Command("curl", "-L", xinit_src_url, "-o", "source.tar.gz")
	err = xinit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xinit_cmd_direct := exec.Command("./binary")
	err = xinit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: tradcpp")
exec.Command("latte", "install", "tradcpp")
	fmt.Println("Instalando dependencia: xorg-server")
exec.Command("latte", "install", "xorg-server")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: xauth")
exec.Command("latte", "install", "xauth")
	fmt.Println("Instalando dependencia: xmodmap")
exec.Command("latte", "install", "xmodmap")
	fmt.Println("Instalando dependencia: xrdb")
exec.Command("latte", "install", "xrdb")
	fmt.Println("Instalando dependencia: xterm")
exec.Command("latte", "install", "xterm")
	fmt.Println("Instalando dependencia: lndir")
exec.Command("latte", "install", "lndir")
	fmt.Println("Instalando dependencia: mkfontscale")
exec.Command("latte", "install", "mkfontscale")
	fmt.Println("Instalando dependencia: quartz-wm")
exec.Command("latte", "install", "quartz-wm")
	fmt.Println("Instalando dependencia: twm")
exec.Command("latte", "install", "twm")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
