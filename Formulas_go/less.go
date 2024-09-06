package main

// Less - Pager program similar to more
// Homepage: https://www.greenwoodsoftware.com/less/index.html

import (
	"fmt"
	
	"os/exec"
)

func installLess() {
	// Método 1: Descargar y extraer .tar.gz
	less_tar_url := "https://www.greenwoodsoftware.com/less/less-661.tar.gz"
	less_cmd_tar := exec.Command("curl", "-L", less_tar_url, "-o", "package.tar.gz")
	err := less_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	less_zip_url := "https://www.greenwoodsoftware.com/less/less-661.zip"
	less_cmd_zip := exec.Command("curl", "-L", less_zip_url, "-o", "package.zip")
	err = less_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	less_bin_url := "https://www.greenwoodsoftware.com/less/less-661.bin"
	less_cmd_bin := exec.Command("curl", "-L", less_bin_url, "-o", "binary.bin")
	err = less_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	less_src_url := "https://www.greenwoodsoftware.com/less/less-661.src.tar.gz"
	less_cmd_src := exec.Command("curl", "-L", less_src_url, "-o", "source.tar.gz")
	err = less_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	less_cmd_direct := exec.Command("./binary")
	err = less_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
