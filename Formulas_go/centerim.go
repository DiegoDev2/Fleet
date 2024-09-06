package main

// CenterIm - Text-mode multi-protocol instant messaging client
// Homepage: https://github.com/petrpavlu/centerim5

import (
	"fmt"
	
	"os/exec"
)

func installCenterIm() {
	// Método 1: Descargar y extraer .tar.gz
	centerim_tar_url := "https://github.com/petrpavlu/centerim5/releases/download/v5.0.1/centerim5-5.0.1.tar.gz"
	centerim_cmd_tar := exec.Command("curl", "-L", centerim_tar_url, "-o", "package.tar.gz")
	err := centerim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	centerim_zip_url := "https://github.com/petrpavlu/centerim5/releases/download/v5.0.1/centerim5-5.0.1.zip"
	centerim_cmd_zip := exec.Command("curl", "-L", centerim_zip_url, "-o", "package.zip")
	err = centerim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	centerim_bin_url := "https://github.com/petrpavlu/centerim5/releases/download/v5.0.1/centerim5-5.0.1.bin"
	centerim_cmd_bin := exec.Command("curl", "-L", centerim_bin_url, "-o", "binary.bin")
	err = centerim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	centerim_src_url := "https://github.com/petrpavlu/centerim5/releases/download/v5.0.1/centerim5-5.0.1.src.tar.gz"
	centerim_cmd_src := exec.Command("curl", "-L", centerim_src_url, "-o", "source.tar.gz")
	err = centerim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	centerim_cmd_direct := exec.Command("./binary")
	err = centerim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libsigc++@2")
exec.Command("latte", "install", "libsigc++@2")
	fmt.Println("Instalando dependencia: pidgin")
exec.Command("latte", "install", "pidgin")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
