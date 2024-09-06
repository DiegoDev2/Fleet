package main

// Xorgproto - X.Org: Protocol Headers
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installXorgproto() {
	// Método 1: Descargar y extraer .tar.gz
	xorgproto_tar_url := "https://xorg.freedesktop.org/archive/individual/proto/xorgproto-2024.1.tar.gz"
	xorgproto_cmd_tar := exec.Command("curl", "-L", xorgproto_tar_url, "-o", "package.tar.gz")
	err := xorgproto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xorgproto_zip_url := "https://xorg.freedesktop.org/archive/individual/proto/xorgproto-2024.1.zip"
	xorgproto_cmd_zip := exec.Command("curl", "-L", xorgproto_zip_url, "-o", "package.zip")
	err = xorgproto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xorgproto_bin_url := "https://xorg.freedesktop.org/archive/individual/proto/xorgproto-2024.1.bin"
	xorgproto_cmd_bin := exec.Command("curl", "-L", xorgproto_bin_url, "-o", "binary.bin")
	err = xorgproto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xorgproto_src_url := "https://xorg.freedesktop.org/archive/individual/proto/xorgproto-2024.1.src.tar.gz"
	xorgproto_cmd_src := exec.Command("curl", "-L", xorgproto_src_url, "-o", "source.tar.gz")
	err = xorgproto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xorgproto_cmd_direct := exec.Command("./binary")
	err = xorgproto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
}
