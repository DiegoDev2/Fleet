package main

// Libxcvt - VESA CVT standard timing modelines generator
// Homepage: https://www.x.org

import (
	"fmt"
	
	"os/exec"
)

func installLibxcvt() {
	// Método 1: Descargar y extraer .tar.gz
	libxcvt_tar_url := "https://www.x.org/releases/individual/lib/libxcvt-0.1.2.tar.xz"
	libxcvt_cmd_tar := exec.Command("curl", "-L", libxcvt_tar_url, "-o", "package.tar.gz")
	err := libxcvt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxcvt_zip_url := "https://www.x.org/releases/individual/lib/libxcvt-0.1.2.tar.xz"
	libxcvt_cmd_zip := exec.Command("curl", "-L", libxcvt_zip_url, "-o", "package.zip")
	err = libxcvt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxcvt_bin_url := "https://www.x.org/releases/individual/lib/libxcvt-0.1.2.tar.xz"
	libxcvt_cmd_bin := exec.Command("curl", "-L", libxcvt_bin_url, "-o", "binary.bin")
	err = libxcvt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxcvt_src_url := "https://www.x.org/releases/individual/lib/libxcvt-0.1.2.tar.xz"
	libxcvt_cmd_src := exec.Command("curl", "-L", libxcvt_src_url, "-o", "source.tar.gz")
	err = libxcvt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxcvt_cmd_direct := exec.Command("./binary")
	err = libxcvt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
