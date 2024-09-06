package main

// Xinput - Utility to configure and test X input devices
// Homepage: https://gitlab.freedesktop.org/xorg/app/xinput

import (
	"fmt"
	
	"os/exec"
)

func installXinput() {
	// Método 1: Descargar y extraer .tar.gz
	xinput_tar_url := "https://www.x.org/pub/individual/app/xinput-1.6.4.tar.xz"
	xinput_cmd_tar := exec.Command("curl", "-L", xinput_tar_url, "-o", "package.tar.gz")
	err := xinput_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xinput_zip_url := "https://www.x.org/pub/individual/app/xinput-1.6.4.tar.xz"
	xinput_cmd_zip := exec.Command("curl", "-L", xinput_zip_url, "-o", "package.zip")
	err = xinput_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xinput_bin_url := "https://www.x.org/pub/individual/app/xinput-1.6.4.tar.xz"
	xinput_cmd_bin := exec.Command("curl", "-L", xinput_bin_url, "-o", "binary.bin")
	err = xinput_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xinput_src_url := "https://www.x.org/pub/individual/app/xinput-1.6.4.tar.xz"
	xinput_cmd_src := exec.Command("curl", "-L", xinput_src_url, "-o", "source.tar.gz")
	err = xinput_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xinput_cmd_direct := exec.Command("./binary")
	err = xinput_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
}
