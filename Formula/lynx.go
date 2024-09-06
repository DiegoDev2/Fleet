package main

// Lynx - Text-based web browser
// Homepage: https://invisible-island.net/lynx/

import (
	"fmt"
	
	"os/exec"
)

func installLynx() {
	// Método 1: Descargar y extraer .tar.gz
	lynx_tar_url := "https://invisible-mirror.net/archives/lynx/tarballs/lynx2.8.9rel.1.tar.bz2"
	lynx_cmd_tar := exec.Command("curl", "-L", lynx_tar_url, "-o", "package.tar.gz")
	err := lynx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lynx_zip_url := "https://invisible-mirror.net/archives/lynx/tarballs/lynx2.8.9rel.1.tar.bz2"
	lynx_cmd_zip := exec.Command("curl", "-L", lynx_zip_url, "-o", "package.zip")
	err = lynx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lynx_bin_url := "https://invisible-mirror.net/archives/lynx/tarballs/lynx2.8.9rel.1.tar.bz2"
	lynx_cmd_bin := exec.Command("curl", "-L", lynx_bin_url, "-o", "binary.bin")
	err = lynx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lynx_src_url := "https://invisible-mirror.net/archives/lynx/tarballs/lynx2.8.9rel.1.tar.bz2"
	lynx_cmd_src := exec.Command("curl", "-L", lynx_src_url, "-o", "source.tar.gz")
	err = lynx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lynx_cmd_direct := exec.Command("./binary")
	err = lynx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
