package main

// Surfraw - Shell Users' Revolutionary Front Rage Against the Web
// Homepage: https://gitlab.com/surfraw/Surfraw

import (
	"fmt"
	
	"os/exec"
)

func installSurfraw() {
	// Método 1: Descargar y extraer .tar.gz
	surfraw_tar_url := "https://ftp.openbsd.org/pub/OpenBSD/distfiles/surfraw-2.3.0.tar.gz"
	surfraw_cmd_tar := exec.Command("curl", "-L", surfraw_tar_url, "-o", "package.tar.gz")
	err := surfraw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	surfraw_zip_url := "https://ftp.openbsd.org/pub/OpenBSD/distfiles/surfraw-2.3.0.zip"
	surfraw_cmd_zip := exec.Command("curl", "-L", surfraw_zip_url, "-o", "package.zip")
	err = surfraw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	surfraw_bin_url := "https://ftp.openbsd.org/pub/OpenBSD/distfiles/surfraw-2.3.0.bin"
	surfraw_cmd_bin := exec.Command("curl", "-L", surfraw_bin_url, "-o", "binary.bin")
	err = surfraw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	surfraw_src_url := "https://ftp.openbsd.org/pub/OpenBSD/distfiles/surfraw-2.3.0.src.tar.gz"
	surfraw_cmd_src := exec.Command("curl", "-L", surfraw_src_url, "-o", "source.tar.gz")
	err = surfraw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	surfraw_cmd_direct := exec.Command("./binary")
	err = surfraw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
