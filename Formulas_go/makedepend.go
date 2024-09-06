package main

// Makedepend - Creates dependencies in makefiles
// Homepage: https://x.org/

import (
	"fmt"
	
	"os/exec"
)

func installMakedepend() {
	// Método 1: Descargar y extraer .tar.gz
	makedepend_tar_url := "https://xorg.freedesktop.org/releases/individual/util/makedepend-1.0.9.tar.xz"
	makedepend_cmd_tar := exec.Command("curl", "-L", makedepend_tar_url, "-o", "package.tar.gz")
	err := makedepend_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	makedepend_zip_url := "https://xorg.freedesktop.org/releases/individual/util/makedepend-1.0.9.tar.xz"
	makedepend_cmd_zip := exec.Command("curl", "-L", makedepend_zip_url, "-o", "package.zip")
	err = makedepend_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	makedepend_bin_url := "https://xorg.freedesktop.org/releases/individual/util/makedepend-1.0.9.tar.xz"
	makedepend_cmd_bin := exec.Command("curl", "-L", makedepend_bin_url, "-o", "binary.bin")
	err = makedepend_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	makedepend_src_url := "https://xorg.freedesktop.org/releases/individual/util/makedepend-1.0.9.tar.xz"
	makedepend_cmd_src := exec.Command("curl", "-L", makedepend_src_url, "-o", "source.tar.gz")
	err = makedepend_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	makedepend_cmd_direct := exec.Command("./binary")
	err = makedepend_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
