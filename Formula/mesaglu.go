package main

// MesaGlu - Mesa OpenGL Utility library
// Homepage: https://cgit.freedesktop.org/mesa/glu

import (
	"fmt"
	
	"os/exec"
)

func installMesaGlu() {
	// Método 1: Descargar y extraer .tar.gz
	mesaglu_tar_url := "ftp://ftp.freedesktop.org/pub/mesa/glu/glu-9.0.3.tar.xz"
	mesaglu_cmd_tar := exec.Command("curl", "-L", mesaglu_tar_url, "-o", "package.tar.gz")
	err := mesaglu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mesaglu_zip_url := "ftp://ftp.freedesktop.org/pub/mesa/glu/glu-9.0.3.tar.xz"
	mesaglu_cmd_zip := exec.Command("curl", "-L", mesaglu_zip_url, "-o", "package.zip")
	err = mesaglu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mesaglu_bin_url := "ftp://ftp.freedesktop.org/pub/mesa/glu/glu-9.0.3.tar.xz"
	mesaglu_cmd_bin := exec.Command("curl", "-L", mesaglu_bin_url, "-o", "binary.bin")
	err = mesaglu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mesaglu_src_url := "ftp://ftp.freedesktop.org/pub/mesa/glu/glu-9.0.3.tar.xz"
	mesaglu_cmd_src := exec.Command("curl", "-L", mesaglu_src_url, "-o", "source.tar.gz")
	err = mesaglu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mesaglu_cmd_direct := exec.Command("./binary")
	err = mesaglu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
