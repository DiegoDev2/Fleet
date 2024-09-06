package main

// A2ps - Any-to-PostScript filter
// Homepage: https://www.gnu.org/software/a2ps/

import (
	"fmt"

	"os/exec"
)

func installA2ps() {
	// Método 1: Descargar y extraer .tar.gz
	a2ps_tar_url := "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz"
	a2ps_cmd_tar := exec.Command("curl", "-L", a2ps_tar_url, "-o", "package.tar.gz")
	err := a2ps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	a2ps_zip_url := "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.zip"
	a2ps_cmd_zip := exec.Command("curl", "-L", a2ps_zip_url, "-o", "package.zip")
	err = a2ps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	a2ps_bin_url := "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.bin"
	a2ps_cmd_bin := exec.Command("curl", "-L", a2ps_bin_url, "-o", "binary.bin")
	err = a2ps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	a2ps_src_url := "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.src.tar.gz"
	a2ps_cmd_src := exec.Command("curl", "-L", a2ps_src_url, "-o", "source.tar.gz")
	err = a2ps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	a2ps_cmd_direct := exec.Command("./binary")
	err = a2ps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: libpaper")
	exec.Command("latte", "install", "libpaper")
}
