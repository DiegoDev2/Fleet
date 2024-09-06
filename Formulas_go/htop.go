package main

// Htop - Improved top (interactive process viewer)
// Homepage: https://htop.dev/

import (
	"fmt"
	
	"os/exec"
)

func installHtop() {
	// Método 1: Descargar y extraer .tar.gz
	htop_tar_url := "https://github.com/htop-dev/htop/archive/refs/tags/3.3.0.tar.gz"
	htop_cmd_tar := exec.Command("curl", "-L", htop_tar_url, "-o", "package.tar.gz")
	err := htop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htop_zip_url := "https://github.com/htop-dev/htop/archive/refs/tags/3.3.0.zip"
	htop_cmd_zip := exec.Command("curl", "-L", htop_zip_url, "-o", "package.zip")
	err = htop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htop_bin_url := "https://github.com/htop-dev/htop/archive/refs/tags/3.3.0.bin"
	htop_cmd_bin := exec.Command("curl", "-L", htop_bin_url, "-o", "binary.bin")
	err = htop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htop_src_url := "https://github.com/htop-dev/htop/archive/refs/tags/3.3.0.src.tar.gz"
	htop_cmd_src := exec.Command("curl", "-L", htop_src_url, "-o", "source.tar.gz")
	err = htop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htop_cmd_direct := exec.Command("./binary")
	err = htop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: lm-sensors")
exec.Command("latte", "install", "lm-sensors")
}
